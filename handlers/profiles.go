package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"user-profile-service/models"
	"user-profile-service/store"
)

var db *store.Database

func SetDatabase(database *store.Database) {
	db = database
}

//Create profile handler -> handles POST /profiles

func CreateProfileHandler(w http.ResponseWriter, r *http.Request) {
	var profile models.Profile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	createdProfile, err := db.CreateProfile(profile)
	if err != nil {
		http.Error(w, "Failed to creage profile", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdProfile)
}

//get profile handler

func GetProfilesHandler(w http.ResponseWriter, r *http.Request) {
	profiles, err := db.GetProfiles()
	if err != nil {
		http.Error(w, "Failed to retrieve profiles", http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(profiles)
}

func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/profiles/"):])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	profile, err := db.GetProfile(id)
	if err != nil {
		http.Error(w, "Failed to receive profile", http.StatusInternalServerError)
		return
	}

	if profile.ID == 0 {
		http.Error(w, "Profile not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(profile)
}
