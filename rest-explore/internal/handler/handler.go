package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"rest-explore/internal/model"
	"rest-explore/internal/service"
)

func HelloTest(wr http.ResponseWriter, r *http.Request) {

	fmt.Fprint(wr, "Hello this is our first API.....")
}

func CreateDataRegistration(wr http.ResponseWriter, r *http.Request) {

	var registrationData model.RegistrationData

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error while reading the data from request.")
	}

	err = json.Unmarshal(data, &registrationData)
	if err != nil {
		log.Println("Error while unmarshaling the data.")
	}

	fmt.Println("Registration Data :", registrationData)

	res, err := service.CreateDataRegistration(registrationData)
	if err != nil {
		json.NewEncoder(wr).Encode(map[string]error{
			"Error": err,
		})
	}

	wr.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(wr).Encode(map[string]*model.RegistrationData{
		"Data": res,
	})

	if err != nil {
		log.Println("Error while unmarshaling the data.")
	}
}

func InsertDataRegistration(w http.ResponseWriter, r *http.Request) {
	var user model.RegistrationData

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	inserted, err := service.InsertData(user)
	if err != nil {
		http.Error(w, "Failed to insert user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]*model.RegistrationData{
		"Inserted": inserted,
	})
}

func GetByNameDataRegistration(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing 'name' parameter", http.StatusBadRequest)
		return
	}

	data, err := service.GetByName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Get All Records
func GetALLDataRegistration(w http.ResponseWriter, r *http.Request) {
	data := service.GetAll()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Update Record: expects full JSON body with ID
func UpdateDataRegistration(w http.ResponseWriter, r *http.Request) {
	var updateData model.RegistrationData
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	updated, err := service.Update(updateData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}

// Delete Record: expects ?email="example@example.com"
func DeleteDataRegistration(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Missing 'email' parameter", http.StatusBadRequest)
		return
	}

	err := service.Delete(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted successfully"))
}
