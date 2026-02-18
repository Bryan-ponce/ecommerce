package reportes

import (
	"ecommerce/pedidos"
	"encoding/json"
	"fmt"
	"net/http"
)

// 1. Interfaz Reportable
type Reportable interface {
	GenerarReporte(pedidos []pedidos.Pedido)
}

// 2. Implementación en consola
type ReporteConsola struct{}

func (r ReporteConsola) GenerarReporte(lista []pedidos.Pedido) {
	fmt.Println("====== Reporte General de Ventas ======")
	for _, pedido := range lista {
		fmt.Printf("Pedido #%d - Cliente: %s - Total: $%.2f\n",
			pedido.ID, pedido.Usuario.Nombre(), pedido.Total)
	}

	fmt.Println("=======================================")
	fmt.Println("=======Reporte en formato JSON=========")
}

// 3. Implementación en JSON
type ReporteJSON struct{}

func (r ReporteJSON) GenerarReporte(lista []pedidos.Pedido) {
	data, _ := json.Marshal(lista)
	fmt.Println(string(data))
}

// 4. Genera un reporte en consola sin necesidad de interfaz.
func ReporteVentas(lista []pedidos.Pedido) {
	fmt.Println("====== Reporte General de Ventas ======")
	for _, pedido := range lista {
		fmt.Printf("Pedido #%d - Cliente: %s - Total: %.2f\n",
			pedido.ID, pedido.Usuario.Nombre(), pedido.Total)
	}
}

// 5. ResponderJSON para endpoints HTTP
func ResponderJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
