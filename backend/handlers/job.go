package handlers 

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syl15/job-tracker/backend/models"
	"github.com/syl15/job-tracker/backend/database"
)

// Handler: defines how to respond to each API request

func GetJobs(c *gin.Context) {

	db := database.GetDB() 

	rows, err := db.Query("SELECT id, company, title, status, date_applied FROM jobs")
	if err != nil {
		log.Println("Query error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch jobs"})
		return
	}
	defer rows.Close() 

	var jobs[]models.Job 

	for rows.Next() {
		var job models.Job 
		if err := rows.Scan(&job.ID, &job.Company, &job.Title, &job.Status, &job.DateApplied); err != nil {
			log.Println("Scan error:", err)
			continue
		}
		jobs = append(jobs, job)
	}

	c.IndentedJSON(http.StatusOK, jobs)
}


func CreateJob(c * gin.Context) {
	var newJob models.Job 

	// Improper formatting
	if err := c.BindJSON(&newJob); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return 
	}

	db := database.GetDB() 
	result, err := db.Exec("INSERT INTO jobs (company, title, status, date_applied) VALUES (?, ?, ?, ?)", 
		newJob.Company, newJob.Title, newJob.Status, newJob.DateApplied)
	if err != nil {
		log.Println("Insert error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert job"})
		return 
	}

	id, _ := result.LastInsertId() 
	newJob.ID = int(id) 

	c.JSON(http.StatusCreated, newJob)
}