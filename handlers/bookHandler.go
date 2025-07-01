// Este archivo contiene los handlers de los libros: obtener, agregar y eliminar
// Cada función responde a una petición del cliente por HTTP

package handlers

import (
	// Importamos el paquete donde están los datos y los modelos de los libros
	"biblioteca/data"
	"biblioteca/models"
	"encoding/json"
	"fmt"
	"net/http"
)

// Función para obtener todos los libros guardados
func GetBooks(w http.ResponseWriter, r *http.Request) {
	data.LoadData()                       // Cargamos los datos desde el archivo JSON
	json.NewEncoder(w).Encode(data.Books) // Enviamos los libros como respuesta JSON
}

// Función para agregar un nuevo libro al sistema
func AddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book                  // Creamos una variable tipo Book
	json.NewDecoder(r.Body).Decode(&book) // Leemos el JSON enviado desde el cliente
	data.LoadData()                       // Cargamos los datos actuales

	data.BookID++                         // Aumentamos el ID para el nuevo libro
	book.ID = data.BookID                 // Le asignamos ese ID al libro nuevo
	data.Books = append(data.Books, book) // Lo agregamos al slice de libros
	data.SaveData()                       // Guardamos los cambios en el archivo JSON

	json.NewEncoder(w).Encode(book) // Enviamos el libro nuevo como respuesta
}

// Función para eliminar un libro según su ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id") // Obtenemos el ID del libro desde la URL
	var id int
	fmt.Sscanf(idStr, "%d", &id) // Convertimos el ID a tipo entero

	data.LoadData() // Cargamos los libros actuales

	// Recorremos la lista de libros para encontrar el que tiene el ID
	for i, b := range data.Books {
		if b.ID == id {
			// Si lo encontramos, lo quitamos usando slicing
			data.Books = append(data.Books[:i], data.Books[i+1:]...)
			break
		}
	}

	data.SaveData()                       // Guardamos la lista actualizada
	json.NewEncoder(w).Encode(data.Books) // Enviamos la lista actual como respuesta
}
