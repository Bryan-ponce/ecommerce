package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"ecommerce/carrito"
	"ecommerce/pedidos"
	"ecommerce/productos"
	"ecommerce/reportes"
	"ecommerce/usuarios"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando archivo .env")
	}
	port := os.Getenv("PORT")

	// Crear productos para la venta
	zapato1 := productos.CrearProducto(1, "Zapato Deportivo", 59.99, 10)
	zapato2 := productos.CrearProducto(2, "Zapato Formal", 89.99, 5)
	zapato3 := productos.CrearProducto(3, "Zapato Casual", 45.50, 20)
	zapato4 := productos.CrearProducto(4, "Zapato Sandalia", 35.00, 15)
	zapato5 := productos.CrearProducto(5, "Zapato Botín", 70.00, 8)
	zapato6 := productos.CrearProducto(6, "Zapato Running", 65.00, 12)
	zapato7 := productos.CrearProducto(7, "Zapato Elegante", 120.00, 4)

	// Crear usuarios o clientes
	cliente1 := usuarios.CrearUsuario(1, "Bryan", "cliente")
	cliente2 := usuarios.CrearUsuario(2, "Ana", "cliente")
	cliente3 := usuarios.CrearUsuario(3, "Carlos", "cliente")
	cliente4 := usuarios.CrearUsuario(4, "Diana", "cliente")
	cliente5 := usuarios.CrearUsuario(5, "Luis", "cliente")

	// Carrito Bryan
	var carritoBryan carrito.Carrito
	carritoBryan.SetUsuarioID(cliente1.ID)
	carritoBryan.AgregarProducto(zapato1, 2)
	carritoBryan.AgregarProducto(zapato2, 1)

	// Carrito Ana
	var carritoAna carrito.Carrito
	carritoAna.SetUsuarioID(cliente2.ID)
	carritoAna.AgregarProducto(zapato3, 3)
	carritoAna.AgregarProducto(zapato4, 2)

	// Carrito Carlos
	var carritoCarlos carrito.Carrito
	carritoCarlos.SetUsuarioID(cliente3.ID)
	carritoCarlos.AgregarProducto(zapato5, 1)
	carritoCarlos.AgregarProducto(zapato6, 2)

	// Carrito Diana
	var carritoDiana carrito.Carrito
	carritoDiana.SetUsuarioID(cliente4.ID)
	carritoDiana.AgregarProducto(zapato7, 1)

	// Carrito Luis
	var carritoLuis carrito.Carrito
	carritoLuis.SetUsuarioID(cliente5.ID)
	carritoLuis.AgregarProducto(zapato1, 1)
	carritoLuis.AgregarProducto(zapato3, 2)

	// Crear pedidos
	pedido1 := pedidos.CrearPedido(1, cliente1, carritoBryan.Items(), carritoBryan.CalcularTotal())
	pedido2 := pedidos.CrearPedido(2, cliente2, carritoAna.Items(), carritoAna.CalcularTotal())
	pedido3 := pedidos.CrearPedido(3, cliente3, carritoCarlos.Items(), carritoCarlos.CalcularTotal())
	pedido4 := pedidos.CrearPedido(4, cliente4, carritoDiana.Items(), carritoDiana.CalcularTotal())
	pedido5 := pedidos.CrearPedido(5, cliente5, carritoLuis.Items(), carritoLuis.CalcularTotal())

	pedidosRealizados := []pedidos.Pedido{pedido1, pedido2, pedido3, pedido4, pedido5}

	// ===================================================================================================
	// Uso de polimorfismo en reportes

	// Reporte en consola
	var reporte reportes.Reportable
	reporte = reportes.ReporteConsola{}
	reporte.GenerarReporte(pedidosRealizados)

	// Reporte en JSON (solo imprime en consola formato JSON)
	reporte = reportes.ReporteJSON{}
	reporte.GenerarReporte(pedidosRealizados)

	// Servidor HTTP con Gorilla Mux
	r := mux.NewRouter()

	r.HandleFunc("/productos", func(w http.ResponseWriter, r *http.Request) {
		lista := []productos.Producto{zapato1, zapato2, zapato3, zapato4, zapato5, zapato6, zapato7}
		productos.ResponderJSON(w, lista)
	}).Methods("GET")

	r.HandleFunc("/reportes", func(w http.ResponseWriter, r *http.Request) {
		reportes.ResponderJSON(w, pedidosRealizados)
	}).Methods("GET")

	r.HandleFunc("/carrito/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		switch id {
		case "1":
			productos.ResponderJSON(w, carritoBryan.Items())
		case "2":
			productos.ResponderJSON(w, carritoAna.Items())
		case "3":
			productos.ResponderJSON(w, carritoCarlos.Items())
		case "4":
			productos.ResponderJSON(w, carritoDiana.Items())
		case "5":
			productos.ResponderJSON(w, carritoLuis.Items())
		default:
			http.Error(w, "Carrito no encontrado", http.StatusNotFound)
		}
	}).Methods("GET")

	fmt.Println("Servidor corriendo en puerto", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
