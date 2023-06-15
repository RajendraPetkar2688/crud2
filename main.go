package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RajendraPetkar2688/crud2/controller"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/getEmployee", controller.AllEmployee).Methods("GET")
	router.HandleFunc("/insertEmployee", controller.InsertEmployee).Methods("POST")
	router.HandleFunc("/updateEmployee", controller.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/deleteEmployee", controller.DeleteEmployee).Methods("DELETE")
	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))
}
