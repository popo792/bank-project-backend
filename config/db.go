package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/microsoft/go-mssqldb"
)

var DB *sql.DB

func ConnectDB() {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		"localhost", "sa", "sabfy1995", 1433, "bank_db")

	var err error
	DB, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err.Error())
	}

	fmt.Println("Connected to MSSQL successfully!")
}
