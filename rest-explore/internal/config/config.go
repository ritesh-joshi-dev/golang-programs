package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewDBConnection() (*sql.DB, error) {

	dsn := "root:rootroot@tcp(localhost:3306)/rest-explore"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %v", err)
	}

	fmt.Println("Successfully connected to MySQL database!")
	return db, nil
}
