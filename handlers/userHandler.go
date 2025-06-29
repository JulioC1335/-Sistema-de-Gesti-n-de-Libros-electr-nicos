package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

var loggedUser *User

func loadUsers() []User {
	data, _ := os.ReadFile("users.json")
	var users []User
	json.Unmarshal(data, &users)
	return users
}

func saveUsers(users []User) {
	data, _ := json.MarshalIndent(users, "", "  ")
	os.WriteFile("users.json", data, 0644)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var u User
	json.NewDecoder(r.Body).Decode(&u)
	users := loadUsers()
	u.ID = len(users) + 1
	users = append(users, u)
	saveUsers(users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(loadUsers())
}

func Login(w http.ResponseWriter, r *http.Request) {
	var cred User
	json.NewDecoder(r.Body).Decode(&cred)
	for _, u := range loadUsers() {
		if u.Name == cred.Name && u.Password == cred.Password {
			loggedUser = &u
			json.NewEncoder(w).Encode(map[string]string{"status": "ok", "message": "Login correcto"})
			return
		}
	}
	http.Error(w, "Usuario o contraseña incorrectos", http.StatusUnauthorized)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	users := loadUsers()
	result := []User{}
	for _, u := range users {
		if u.ID != id {
			result = append(result, u)
		}
	}
	saveUsers(result)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuario eliminado"})
}
