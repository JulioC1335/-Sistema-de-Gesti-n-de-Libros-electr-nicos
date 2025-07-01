// Aquí defino la estructura de un préstamo en la Biblioteca Digital.
// Cada campo está anotado para JSON y acompañado de un comentario que explica su propósito.

package models

// Loan contiene la información sobre cada préstamo de libro.
// ID: identificador único del préstamo.
// UserID: referencia al ID del usuario que lo solicitó.
// BookID: referencia al ID del libro prestado.
// Date: fecha en que se hizo el préstamo (formato YYYY-MM-DD).
// ReturnDate: fecha en que se devolvió el libro (formato YYYY-MM-DD), vacío si aún no se devuelve.
type Loan struct {
	ID         int    `json:"id"`          // Identificador único del préstamo
	UserID     int    `json:"user_id"`     // ID del usuario que pide el libro
	BookID     int    `json:"book_id"`     // ID del libro que se presta
	Date       string `json:"date"`        // Fecha de inicio del préstamo
	ReturnDate string `json:"return_date"` // Fecha de devolución (vacío si pendiente)
}
