package carrito

import (
	"ecommerce/productos"
	"fmt"
)

type ItemCarrito struct {
	producto productos.Producto
	cantidad int
}

type Carrito struct {
	usuarioID int
	items     []ItemCarrito
}

// Getter para obtener los ítems del carrito
func (c Carrito) Items() []ItemCarrito { return c.items }

// Getter para obtener el ID del usuario
func (c Carrito) UsuarioID() int { return c.usuarioID }

// Setter para asignar el usuario al carrito
func (c *Carrito) SetUsuarioID(id int) {
	c.usuarioID = id
}

// AgregarProducto añade un producto al carrito con validación de stock
func (c *Carrito) AgregarProducto(p productos.Producto, cantidad int) error {
	if !productos.ValidarStock(p, cantidad) {
		return fmt.Errorf("Stock insuficiente para %s", p.Nombre())
	}
	c.items = append(c.items, ItemCarrito{producto: p, cantidad: cantidad})
	return nil
}

// CalcularTotal recorre todos los ítems del carrito,
// multiplicando el precio por la cantidad y acumulando el resultado.
// Este método encapsula la lógica de negocio para obtener el total
// de la compra de un usuario.
func (c Carrito) CalcularTotal() float64 {
	total := 0.0
	for _, item := range c.items {
		total += item.producto.Precio() * float64(item.cantidad)
	}
	return total
}
