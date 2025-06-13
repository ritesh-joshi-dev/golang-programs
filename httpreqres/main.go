package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RegistrationData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `josn:"password"`
	MobileNo string `json:"mobile_no"`
}

func HelloTest(wr http.ResponseWriter, r *http.Request) {

	fmt.Fprint(wr, "Hello this is our first API.....")
}

func GetDataRegistration(wr http.ResponseWriter, r *http.Request) {

	var registrationData RegistrationData

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error while reading the data from request.")
	}

	err = json.Unmarshal(data, &registrationData)
	if err != nil {
		log.Println("Error while unmarshaling the data.")
	}

	fmt.Println("Registration Data :", registrationData)

	wr.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(wr).Encode(map[string]RegistrationData{
		"Data": registrationData,
	})

	if err != nil {
		log.Println("Error while unmarshaling the data.")
	}
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/google/test", HelloTest)
	router.HandleFunc("/api/getData", GetDataRegistration).Methods("POST")

	fmt.Println("Server started running on 8090.....")

	http.ListenAndServe(":8090", router)

}
