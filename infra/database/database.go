package database

import (
	"fmt"
	"log"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB    *sql.DB
	DBErr error
)

// DBConnection create database connection
func DBConnection(masterDSN string) error {
	var db = DB

	var err error
	db, err = sql.Open("mysql", masterDSN)
	if err != nil {
		panic(err.Error())
	}
	// defer db.Close()
	fmt.Println("DB Connection Success!")

	if err != nil {
		DBErr = err
		log.Println("Db connection error")
		return err
	}

	// err = db.AutoMigrate(migrationModels...)

	if err != nil {
		return err
	}
	DB = db

	return nil
}

// DB Con Close
func CloseDB() error {
	return DB.Close()
}

// GetDB connection
func GetDB() *sql.DB {
	return DB
}

// GetDBError connection error
func GetDBError() error {
	return DBErr
}
