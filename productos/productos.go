package productos

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Producto struct {
	ID     int     `json:"id"`
	nombre string  //privado
	precio float64 //privado
	stock  int     //privado
}

func CrearProducto(id int, nombre string, precio float64, stock int) Producto {
	return Producto{ID: id, nombre: nombre, precio: precio, stock: stock}
}

func (p Producto) Nombre() string  { return p.nombre }
func (p Producto) Precio() float64 { return p.precio }
func (p Producto) Stock() int      { return p.stock }

func (p *Producto) SetStock(cantidad int) error {
	if cantidad < 0 {
		return fmt.Errorf("El stock no puede ser negativo")
	}
	p.stock = cantidad
	return nil
}

func ValidarStock(p Producto, cantidad int) bool {
	return p.stock >= cantidad
}

func ResponderJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
