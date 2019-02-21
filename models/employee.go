package models

import (
	_"database/sql"
	"github.com/ilhamarrouf/echo-example/db"
	"fmt"
	"database/sql"
)

type Employee struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Salary string `json:"salary"`
	Age    string `json:"age"`
}

type Employees struct {
	Employees []Employee `json:"data"`
}

var connection *sql.DB

func GetEmployee() Employees {
	connection := db.CreateConnection()

	sqlStatement := "SELECT * FROM employee ORDER BY id"

	rows, err := connection.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	result := Employees{}

	for rows.Next() {
		employee := Employee{}
		err2 := rows.Scan(&employee.Id, &employee.Name, &employee.Salary, &employee.Age)

		if err2 != nil {
			fmt.Println(err2)
		}

		result.Employees = append(result.Employees, employee)
	}

	return result
}