package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alexsantossilva/go-api-rest/database"
	"github.com/alexsantossilva/go-api-rest/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersonas(w http.ResponseWriter, r *http.Request) {
	var personas []models.Persona
	database.DB.Find(&personas)

	json.NewEncoder(w).Encode(personas)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var person models.Persona
	database.DB.First(&person, id)

	json.NewEncoder(w).Encode(person)
}

func NewPerson(w http.ResponseWriter, r *http.Request) {
	var newPerson models.Persona
	json.NewDecoder(r.Body).Decode(&newPerson)
	database.DB.Create(&newPerson)
	json.NewEncoder(w).Encode(newPerson)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var person models.Persona
	database.DB.Delete(&person, id)
	json.NewEncoder(w).Encode(person)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var person models.Persona
	database.DB.First(&person, id)

	json.NewDecoder(r.Body).Decode(&person)
	database.DB.Save(&person)
	json.NewEncoder(w).Encode(person)
}
