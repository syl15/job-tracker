package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	defer db.Close() 

	createTable := `
	CREATE TABLE IF NOT EXISTS jobs (
		id INT AUTO_INCREMENT PRIMARY KEY, 
		company VARCHAR(255), 
		title VARCHAR(255), 
		status VARCHAR(50), 
		date_applied DATE
	);`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	// Clear existing data
	_, err = db.Exec("DELETE FROM jobs")
	if err != nil {
		log.Fatal("Failed to clear jobs table:", err)
	}

	insert := `INSERT INTO jobs (company, title, status, date_applied) VALUES (?, ?, ?, ?)`
	_, err = db.Exec(insert, "Meta", "SWE", "interviewing", "2025-07-25")
	if err != nil {
		log.Fatal("Failed to insert Meta job:", err)
	}
	_, err = db.Exec(insert, "Google", "SWE", "applied", "2025-07-27")
	if err != nil {
		log.Fatal("Failed to insert Google job:", err)
	}


	log.Println("Database seeded succesfully")


}
