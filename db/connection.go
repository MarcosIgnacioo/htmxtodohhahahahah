package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)
var database *sql.DB

func GetConnection() *sql.DB {
    
    if database != nil {
        return database
    }    
    var err error     
    database, err = sql.Open("sqlite3", "usuarios.db")
    if err != nil {
        panic(err)
    }
    return database
}
