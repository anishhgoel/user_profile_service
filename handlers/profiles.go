package handlers

import (
	"encoding/json"
	"log"
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
		log.Printf("Error creating profile: %v", err)
		http.Error(w, "Failed to creage profile", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdProfile)
}

//get profile handler

func GetProfilesHandler(w http.ResponseWriter, r *http.Request) {
	profiles, err := db.GetProfiles()
	if err != nil {
		http.Error(w, "Failed to retrieve profiles", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/profiles/"):])
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	var updatedProfile models.Profile
	if err := json.NewDecoder(r.Body).Decode(&updatedProfile); err != nil {
		http.Error(w, "Invlaid JSON", http.StatusBadRequest)
		return
	}

	profile, err := db.UpdateProfile(id, updatedProfile)
	if err != nil {
		http.Error(w, "Failed to update profile", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)

}

func DeleteProfileHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/profiles/"):])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := db.DeleteProfile(id); err != nil {
		http.Error(w, "Failed to delete profile", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
