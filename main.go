package main

import (
	"fmt"
	"log"
	"net/http"
	"user-profile-service/handlers"
	"user-profile-service/store"
)

func main() {
	database := store.NewDataBase("./profiles.db")
	defer database.Close()

	handlers.SetDatabase(database)

	http.HandleFunc("/profiles", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetProfilesHandler(w, r)
		case http.MethodPost:
			handlers.CreateProfileHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/profiles/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetProfileHandler(w, r)
		case http.MethodPut:
			handlers.UpdateProfileHandler(w, r)
		case http.MethodDelete:
			handlers.DeleteProfileHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server runnong on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
