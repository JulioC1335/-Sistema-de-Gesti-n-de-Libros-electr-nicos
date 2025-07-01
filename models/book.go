// Este archivo define las estructuras de datos que uso en mi Biblioteca Digital.

package models

// Book representa la información de un libro en el sistema.
// Cada campo está anotado para serializarse/deserializarse desde JSON.
type Book struct {
	ID     int    `json:"id"`     // Identificador único del libro
	Title  string `json:"title"`  // Título del libro
	Author string `json:"author"` // Nombre del autor del libro
	Stock  int    `json:"stock"`  // Número de ejemplares disponibles
}
