package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Stock  int    `json:"stock"`
	UserID int    `json:"user_id"`
}

func loadBooks() []Book {
	data, _ := os.ReadFile("books.json")
	var books []Book
	json.Unmarshal(data, &books)
	return books
}

func saveBooks(books []Book) {
	data, _ := json.MarshalIndent(books, "", "  ")
	os.WriteFile("books.json", data, 0644)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var b Book
	json.NewDecoder(r.Body).Decode(&b)
	b.ID = len(loadBooks()) + 1
	if loggedUser != nil {
		b.UserID = loggedUser.ID
	}
	books := loadBooks()
	books = append(books, b)
	saveBooks(books)
	json.NewEncoder(w).Encode(b)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(loadBooks())
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	books := loadBooks()
	newList := []Book{}
	for _, b := range books {
		if b.ID != id || (loggedUser != nil && b.UserID != loggedUser.ID) {
			newList = append(newList, b)
		}
	}
	saveBooks(newList)
	json.NewEncoder(w).Encode(map[string]string{"message": "Libro eliminado"})
}
