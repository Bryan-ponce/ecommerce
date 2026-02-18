package usuarios

import "fmt"

type Usuario struct {
	ID     int
	nombre string
	rol    string
}

func CrearUsuario(id int, nombre string, rol string) Usuario {
	return Usuario{ID: id, nombre: nombre, rol: rol}
}

func (u Usuario) Nombre() string { return u.nombre }
func (u Usuario) Rol() string    { return u.rol }

func (u *Usuario) SetRol(rol string) error {
	if rol != "cliente" && rol != "admin" {
		return fmt.Errorf("Rol inválido: %s", rol)
	}
	u.rol = rol
	return nil
}
