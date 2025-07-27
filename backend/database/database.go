package database 

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db 	*sql.DB 
	once sync.Once // Concurrency-safe way to ensure code runs only once (don't initialize multiple DB connections)
) 

// Returns singleton DB connection
func GetDB() *sql.DB {
	once.Do(func() {

		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
		var errOpen error 
		if db, errOpen = sql.Open("mysql", dsn); errOpen != nil {
			log.Fatalf("Failed to connect to database: %v", errOpen)
		}

		// Verify connection is alive 
		if errPing := db.Ping(); errPing != nil {
			log.Fatalf("Failed to ping database: %v", errPing)
		}

		})

	return db

}