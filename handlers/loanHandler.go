package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

type Loan struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	BookID int `json:"book_id"`
}

func loadLoans() []Loan {
	data, _ := os.ReadFile("loans.json")
	var loans []Loan
	json.Unmarshal(data, &loans)
	return loans
}

func saveLoans(loans []Loan) {
	data, _ := json.MarshalIndent(loans, "", "  ")
	os.WriteFile("loans.json", data, 0644)
}

func AddLoan(w http.ResponseWriter, r *http.Request) {
	var l Loan
	json.NewDecoder(r.Body).Decode(&l)
	l.ID = len(loadLoans()) + 1
	loans := loadLoans()
	loans = append(loans, l)
	saveLoans(loans)
	json.NewEncoder(w).Encode(l)
}

func GetLoans(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(loadLoans())
}

func DeleteLoan(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	loans := loadLoans()
	newList := []Loan{}
	for _, l := range loans {
		if l.ID != id {
			newList = append(newList, l)
		}
	}
	saveLoans(newList)
	json.NewEncoder(w).Encode(map[string]string{"message": "Préstamo eliminado"})
}
