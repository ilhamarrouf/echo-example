package db

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

func CreateConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("DB is connected.")
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("DB isn't connected.")
	}

	return db
}