// Este archivo es el punto de entrada del servidor web de la Biblioteca Digital
// Aquí se configuran las rutas de la API y la página principal

package main

import (
	// Importamos el paquete donde están los handlers de todas las funciones
	"biblioteca/handlers"
	"fmt"
	"net/http"
)

func main() {
	// Esta ruta sirve la página principal HTML desde la carpeta 'static'
	http.HandleFunc("/", serveHTML)

	// Esta línea sirve para mostrar archivos estáticos como imágenes o estilos
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Rutas de la API para gestionar libros y usuarios

	// Ruta para obtener y agregar libros (GET y POST)
	http.HandleFunc("/api/books", route(handlers.GetBooks, handlers.AddBook))

	// Ruta para obtener y agregar usuarios (GET y POST)
	http.HandleFunc("/api/users", route(handlers.GetUsers, handlers.AddUser))

	// Ruta para ver y crear préstamos de libros (GET y POST)
	http.HandleFunc("/api/loans", route(handlers.GetLoans, handlers.AddLoan))

	// Ruta para finalizar un préstamo (devolver un libro)
	http.HandleFunc("/api/returnLoan", handlers.FinalizeLoan)

	// Ruta para eliminar libros (por ID o nombre, depende del handler)
	http.HandleFunc("/api/deleteBook", handlers.DeleteBook)

	// Ruta para eliminar usuarios del sistema
	http.HandleFunc("/api/deleteUser", handlers.DeleteUser)

	// Mostrar mensaje cuando el servidor está corriendo
	fmt.Println("Servidor corriendo en http://localhost:8080")

	// Esta línea enciende el servidor en el puerto 8080
	http.ListenAndServe(":8080", nil)
}

// Esta función recibe dos funciones: una para GET y otra para POST
// Se usa para evitar repetir código y definir ambas acciones en la misma ruta
func route(get, post http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			get(w, r)
		case http.MethodPost:
			post(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	}
}

// Esta función muestra el archivo HTML principal en la ruta raíz
func serveHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}
