package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/* *** I CAN'T SEEM TO FIX LOCAL FILE IMPORTS, SO I PUT THE ENTITY + THE REPOSITORY METHODS IN THE MAIN FILE *** */
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// Setting up a slice for all the users here, since im not goint to use a database for this project
var users []User

// http://localhost:8000/users
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// http://localhost:8000/users
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}

// http://localhost:8000/users/1
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// http://localhost:8000/users (body: { "email": "133@331", "phone": "98989898", "password": "hidden" })
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = strconv.Itoa(rand.Intn(100000000))

	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

// http://localhost:8000/users/2 (body: { "email": "update@new", "phone": "777777777", "password": "hidden" })
func updateUser(w http.ResponseWriter, r *http.Request) {
	// Set json content Yupe
	w.Header().Set("Content-Type", "application/json")
	// Params
	params := mux.Vars(r)
	// Loop over the users range
	// Delete the user with teh ID sent
	// Add a new user(the one we send in the body)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			var user User
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.ID = params["id"]
			users = append(users, user)
			json.NewEncoder(w).Encode(user)
			return
		}
	}
}

func main() {
	// Basic Print function
	fmt.Println("Hello, World!")

	r := mux.NewRouter()

	// Creating mock data
	users = append(users, User{ID: "1", Email: "1@1", Phone: "1234", Password: "pass"})
	users = append(users, User{ID: "2", Email: "2@2", Phone: "3232", Password: "word"})

	// The 5 methods
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
