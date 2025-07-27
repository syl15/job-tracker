package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/syl15/job-tracker/backend/database"
	"github.com/syl15/job-tracker/backend/router"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db := database.GetDB() // Get shared db connection 
	defer db.Close() // Close connection when main exits 
		// TODO: Better approach: graceful shutdown with os.Interrupt and syscall.SIGTERM

	// Check if db connection is alive
	if err := db.Ping(); err !=  nil {
		log.Fatalf("Database ping failed: %v", err)
	}

	fmt.Println("Database connection successful!")

	r := router.SetupRouter()
	r.Run(":8080")
}