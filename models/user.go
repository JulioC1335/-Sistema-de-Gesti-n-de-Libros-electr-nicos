// Este archivo define la estructura de un usuario en la Biblioteca Digital.
// Cada campo está anotado para serializar/deserializar desde JSON.

package models

// User representa un usuario registrado en el sistema.
// ID: identificador único del usuario.
// Name: nombre que utiliza para iniciar sesión o identificarlo.
type User struct {
	ID   int    `json:"id"`   // Identificador único del usuario
	Name string `json:"name"` // Nombre completo o alias del usuario
}
