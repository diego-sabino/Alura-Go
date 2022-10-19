package controllers

import (
	"encoding/json"
	"fmt"
	"go-rest-api/db"
	"go-rest-api/models"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var p []models.Personality
	db.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func ReturnPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p []models.Personality
	db.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}

func CreateNewPersonality(w http.ResponseWriter, r *http.Request) {
	var newPersonality models.Personality
	json.NewDecoder(r.Body).Decode(&newPersonality)
	db.DB.Create(&newPersonality)
	json.NewEncoder(w).Encode(newPersonality)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personality
	db.DB.Delete(&p, id)
	json.NewEncoder(w).Encode("Personalidade deletada com sucesso")
}

func EditPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personality
	db.DB.First(&p, id)
	json.NewDecoder(r.Body).Decode(&p)
	db.DB.Save(&p)
	json.NewEncoder(w).Encode(p)
}
