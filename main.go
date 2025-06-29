package main

import (
	"biblioteca/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveHTML)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/api/books", route(handlers.GetBooks, handlers.AddBook))
	http.HandleFunc("/api/users", route(handlers.GetUsers, handlers.AddUser))
	http.HandleFunc("/api/loans", route(handlers.GetLoans, handlers.AddLoan))

	http.HandleFunc("/api/deleteBook", handlers.DeleteBook)
	http.HandleFunc("/api/deleteUser", handlers.DeleteUser)
	http.HandleFunc("/api/deleteLoan", handlers.DeleteLoan)

	http.HandleFunc("/api/login", handlers.Login)

	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

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

func serveHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}
