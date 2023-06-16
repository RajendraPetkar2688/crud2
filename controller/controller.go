package controller

import (
	"database/sql"
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

// singleEmployee = Single Employee API
func SingleEmployee(w http.ResponseWriter, r *http.Request) {
	var employee model.Employee
	var response model.Response

	db := config.Connect()
	defer db.Close()
	id := r.FormValue("id")
	// Prepare the SQL statement
	stmt, err := db.Prepare("select * from employee where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	// Execute the statement
	//var data model.Employee
	//var Employee model.Employee

	err = row.Scan(&employee.Id, &employee.Name, &employee.City, &employee.Mobile)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Record not found", http.StatusNotFound)
			return
		}
		log.Fatal(err)
	}
	//fmt.Print(row)
	jsonData, err := json.Marshal(employee)
	//fmt.Print(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Send the response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Record fetched successfully\n"))
	w.Write(jsonData)

	json.NewEncoder(w).Encode(response)
}

// UpdateEmployee = Update Employee API
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var response model.Response
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

	if name != "" && city == "" {
		_, err = db.Exec("UPDATE employee SET name=? WHERE id=?", name, id)
	} else if city != "" && name == "" {
		_, err = db.Exec("UPDATE employee SET city=? WHERE id=?", city, id)
	} else if name == "" && mobile == "" {
		_, err = db.Exec("UPDATE employee SET name=?, mobile=? WHERE id=? ", name, mobile, id)
	} else {
		_, err = db.Exec("UPDATE employee SET name=?, city=?, mobile=? WHERE id=? ", name, city, mobile, id)
	}

	if err != nil {
		log.Print(err)
	}

	response.Status = 200
	response.Message = "Record updated successfully"
	fmt.Print("Record updated successfully")

	w.Header().Set("Content-Type", "application/json")
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

// DeleteEmployee = Delete Employee API
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}
	id := r.FormValue("id")

	_, err = db.Exec("DELETE  from employee where id=? ", id)

	if err != nil {
		log.Print(err)
	}

	response.Status = 200
	response.Message = " Record deleted successfully"
	fmt.Print("Record deleted successfully")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// delete by ID
// singleEmployee = Single Employee API
func DeleteEmployeeByID(w http.ResponseWriter, r *http.Request) {
	var employee model.Employee
	var response model.Response

	db := config.Connect()
	defer db.Close()
	id := r.FormValue("id")
	// Prepare the SQL statement
	stmt, err := db.Prepare("select * from employee where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	// Execute the statement
	//var data model.Employee
	//var Employee model.Employee

	err = row.Scan(&employee.Id, &employee.Name, &employee.City, &employee.Mobile)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Record not found", http.StatusNotFound)
			return
		}
		log.Fatal(err)
	}

	_, err = db.Exec("DELETE  from employee where id=? ", id)

	if err != nil {
		log.Print(err)
	}

	response.Status = 200
	response.Message = " Record deleted successfully"
	fmt.Print("Record deleted successfully")

	//fmt.Print(row)
	jsonData, err := json.Marshal(employee)
	//fmt.Print(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Send the response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Record fetched successfully\n"))
	w.Write(jsonData)

	json.NewEncoder(w).Encode(response)
}
