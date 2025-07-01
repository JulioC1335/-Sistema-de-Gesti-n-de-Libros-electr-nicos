// Aquí definimos las funciones para gestionar usuarios: listar, agregar y eliminar.
// Cada función atiende peticiones HTTP y responde con JSON.

package handlers

import (
	// data: carga y guarda datos en archivos JSON
	"biblioteca/data"
	// models: define la estructura de User
	"biblioteca/models"
	// encoding/json: codifica y decodifica JSON
	"encoding/json"
	// fmt: para convertir strings a otros tipos
	"fmt"
	// net/http: maneja solicitudes y respuestas HTTP
	"net/http"
)

// GetUsers devuelve la lista de todos los usuarios registrados.
// Carga los datos del archivo y responde con el slice data.Users en JSON.
func GetUsers(w http.ResponseWriter, r *http.Request) {
	data.LoadData()                       // Cargo los datos actualizados
	json.NewEncoder(w).Encode(data.Users) // Envío la lista de usuarios en formato JSON
}

// AddUser registra un nuevo usuario en el sistema.
// Decodifica el JSON recibido, asigna un ID y guarda el cambio.
func AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User                  // Variable para decodificar el cuerpo de la petición
	json.NewDecoder(r.Body).Decode(&user) // Leo el JSON enviado por el cliente

	data.LoadData()                       // Cargo usuarios actuales
	data.UserID++                         // Incremento el contador de IDs
	user.ID = data.UserID                 // Asigno el nuevo ID al usuario
	data.Users = append(data.Users, user) // Agrego el usuario al slice
	data.SaveData()                       // Guardo los cambios en el archivo JSON

	json.NewEncoder(w).Encode(user) // Devuelvo el usuario recién creado
}

// DeleteUser elimina un usuario según su ID.
// Obtiene el ID de la URL, busca el usuario y lo quita del slice.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id") // Obtengo el ID desde la ruta: ?id=X
	var id int
	fmt.Sscanf(idStr, "%d", &id) // Convierto el ID de string a entero

	data.LoadData() // Cargo usuarios actuales

	// Recorro el slice para encontrar el usuario con el ID dado
	for i, u := range data.Users {
		if u.ID == id {
			// Al encontrarlo, lo elimino usando slicing
			data.Users = append(data.Users[:i], data.Users[i+1:]...)
			break
		}
	}

	data.SaveData()                       // Guardo la lista actualizada
	json.NewEncoder(w).Encode(data.Users) // Envío la lista de usuarios tras la eliminación
}
