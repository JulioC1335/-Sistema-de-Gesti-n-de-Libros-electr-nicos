// data/data.go
// Este paquete gestiona la carga y el guardado de datos en archivos JSON.
// Aquí se almacenan libros, usuarios y préstamos junto con los últimos IDs.

package data

import (
	// models: define las estructuras Book, User y Loan
	"biblioteca/models"
	// encoding/json: para codificar y decodificar JSON
	"encoding/json"
	// os: para abrir y crear archivos
	"os"
)

// Slices globales donde guardo temporalmente los datos en memoria
var Books []models.Book
var Users []models.User
var Loans []models.Loan

// Variables que llevo como contadores de IDs para nuevos registros
var BookID, UserID, LoanID int

// SaveData guarda en disco todos los slices: libros, usuarios y préstamos
func SaveData() {
	// Llamo a la función genérica para cada tipo de dato
	saveToFile("data/books.json", Books)
	saveToFile("data/users.json", Users)
	saveToFile("data/loans.json", Loans)
}

// LoadData lee los archivos JSON y actualiza los slices y los contadores de ID
func LoadData() {
	// Cargo cada archivo y decodifico su JSON en el slice correspondiente
	loadFromFile("data/books.json", &Books)
	loadFromFile("data/users.json", &Users)
	loadFromFile("data/loans.json", &Loans)

	// Reviso los IDs en los datos cargados para inicializar los contadores
	for _, b := range Books {
		if b.ID > BookID {
			BookID = b.ID
		}
	}
	for _, u := range Users {
		if u.ID > UserID {
			UserID = u.ID
		}
	}
	for _, l := range Loans {
		if l.ID > LoanID {
			LoanID = l.ID
		}
	}
}

// saveToFile crea o sobrescribe un archivo y escribe el JSON de “data” en él
func saveToFile(filename string, data interface{}) {
	file, _ := os.Create(filename)     // Creo el archivo (lo borra si existe)
	defer file.Close()                 // Aseguro cerrar el archivo al final
	json.NewEncoder(file).Encode(data) // Escribo el slice codificado en JSON
}

// loadFromFile abre un archivo JSON y decodifica su contenido en “out”
func loadFromFile(filename string, out interface{}) {
	file, err := os.Open(filename) // Abro el archivo para lectura
	if err != nil {
		return // Si no existe o hay error, salgo sin hacer nada
	}
	defer file.Close()                // Cierro el archivo al terminar
	json.NewDecoder(file).Decode(out) // Decodifico el JSON en la variable apuntada
}
