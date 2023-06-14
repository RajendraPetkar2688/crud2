package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RajendraPetkar2688/crud2/config"
	"github.com/RajendraPetkar2688/crud2/model"
)

// AllEmployee = Select Employee API
func AllEmployee(w http.ResponseWriter, r *http.Request) {
	var employee model.Employee
	var response model.Response
	var arrEmployee []model.Employee

	db := config.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, city, mobile FROM employee")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&employee.Id, &employee.Name, &employee.City, &employee.Mobile)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrEmployee = append(arrEmployee, employee)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrEmployee

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

// InsertEmployee = Insert Employee API
func InsertEmployee(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	//w.Header().Set("Content-Type", "application/json")
	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}
	id := r.FormValue("id")
	name := r.FormValue("name")
	city := r.FormValue("city")
	mobile := r.FormValue("mobile")

	_, err = db.Exec("INSERT INTO employee(id, name, city, mobile) VALUES(?, ?, ?,?)", id, name, city, mobile)

	if err != nil {
		log.Print(err)
		return
	}
	response.Status = 200
	response.Message = "Insert data successfully"
	fmt.Print("Insert data to database")

	w.Header().Set("Content-Type", "multipart/form-data")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}
