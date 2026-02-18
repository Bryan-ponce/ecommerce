package pedidos

import (
	"ecommerce/carrito"
	"ecommerce/usuarios"
)

type Pedido struct {
	ID      int
	Usuario usuarios.Usuario
	Items   []carrito.ItemCarrito
	Total   float64
}

func CrearPedido(id int, u usuarios.Usuario, items []carrito.ItemCarrito, total float64) Pedido {
	return Pedido{ID: id, Usuario: u, Items: items, Total: total}
}
