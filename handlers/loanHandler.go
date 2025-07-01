// Aquí tengo las funciones para gestionar los préstamos de la biblioteca.
// Cada función recibe peticiones HTTP y responde con JSON.

package handlers

import (
	// data: carga y guarda datos en archivos JSON
	"biblioteca/data"
	// models: define la estructura de Loan
	"biblioteca/models"
	// encoding/json: codifica y decodifica JSON
	"encoding/json"
	// fmt: para convertir strings a otros tipos
	"fmt"
	// net/http: maneja solicitudes y respuestas HTTP
	"net/http"
	// time: trabaja con fechas y horas
	"time"
)

// ResponseLoan es la estructura que devuelvo al cliente para cada préstamo,
// incluyendo nombres legibles y fechas en formato string.
type ResponseLoan struct {
	ID          int    `json:"id"`           // Identificador del préstamo
	Usuario     string `json:"usuario"`      // Nombre del usuario que pide el libro
	Libro       string `json:"libro"`        // Título del libro prestado
	FechaInicio string `json:"fecha_inicio"` // Fecha en que se realizó el préstamo
	FechaFin    string `json:"fecha_fin"`    // Fecha en que se devolvió el libro
}

// GetLoans responde con la lista de préstamos actuales.
// Lee los datos, busca nombres de usuario y títulos de libro, y arma el JSON.
func GetLoans(w http.ResponseWriter, r *http.Request) {
	data.LoadData()           // Cargo usuarios, libros y préstamos de archivos
	var result []ResponseLoan // Preparo un slice para la respuesta

	for _, l := range data.Loans { // Recorro cada préstamo
		var userName, bookTitle string

		// Busco el nombre del usuario que hizo el préstamo
		for _, u := range data.Users {
			if u.ID == l.UserID {
				userName = u.Name
			}
		}
		// Busco el título del libro que se prestó
		for _, b := range data.Books {
			if b.ID == l.BookID {
				bookTitle = b.Title
			}
		}
		// Agrego el préstamo formateado al resultado
		result = append(result, ResponseLoan{
			ID:          l.ID,
			Usuario:     userName,
			Libro:       bookTitle,
			FechaInicio: l.Date,
			FechaFin:    l.ReturnDate,
		})
	}
	// Envío la lista de préstamos en formato JSON
	json.NewEncoder(w).Encode(result)
}

// AddLoan recibe datos de un nuevo préstamo, lo guarda y devuelve el registro creado.
func AddLoan(w http.ResponseWriter, r *http.Request) {
	var loan models.Loan                  // Estructura donde decodifico el JSON entrante
	json.NewDecoder(r.Body).Decode(&loan) // Leo los campos enviados por el cliente

	data.LoadData()       // Cargo datos actuales
	data.LoanID++         // Aumento el contador de ID
	loan.ID = data.LoanID // Asigno el nuevo ID al préstamo
	// Guardo la fecha de inicio con formato YYYY-MM-DD
	loan.Date = time.Now().Format("2006-01-02")
	loan.ReturnDate = "" // Dejo la fecha de devolución vacía por ahora

	data.Loans = append(data.Loans, loan) // Agrego el préstamo al slice
	data.SaveData()                       // Guardo cambios en el archivo JSON

	// Respondo con el préstamo recién creado
	json.NewEncoder(w).Encode(loan)
}

// FinalizeLoan marca la devolución de un préstamo existente.
// Busca por ID y actualiza la fecha de retorno con la fecha actual.
func FinalizeLoan(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id") // Obtengo el ID desde la URL
	var id int
	fmt.Sscanf(idStr, "%d", &id) // Convierto el ID a entero

	data.LoadData() // Cargo datos actuales

	for i, l := range data.Loans { // Busco el préstamo con ese ID
		if l.ID == id {
			// Marco la fecha de devolución con la fecha actual
			data.Loans[i].ReturnDate = time.Now().Format("2006-01-02")
			data.SaveData() // Guardo los cambios
			// Devuelvo el préstamo actualizado
			json.NewEncoder(w).Encode(data.Loans[i])
			return
		}
	}
	// Si no lo encuentra, respondo con error 404
	http.Error(w, "Préstamo no encontrado", http.StatusNotFound)
}
