package internal

import (
	"fmt"
	"net/http"
	"rest-explore/internal/handler"

	"github.com/gorilla/mux"
)

func Initroutes() {
	router := mux.NewRouter()
	router.HandleFunc("/api/google/test", handler.HelloTest)
	router.HandleFunc("/api/CreateData", handler.CreateDataRegistration).Methods("POST")
	router.HandleFunc("/api/insertData", handler.InsertDataRegistration).Methods("POST")
	router.HandleFunc("/api/GetData", handler.GetByNameDataRegistration).Methods("GET")
	router.HandleFunc("/api/GetAllData", handler.GetALLDataRegistration).Methods("GET")
	router.HandleFunc("/api/UpdateData", handler.UpdateDataRegistration).Methods("POST")
	router.HandleFunc("/api/DeleteData", handler.DeleteDataRegistration).Methods("DELETE")

	fmt.Println("Server started running on 8090.....")

	http.ListenAndServe(":8090", router)
}
