package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syl15/job-tracker/backend/database"
	"github.com/syl15/job-tracker/backend/models"
)

// Handler: defines how to respond to each API request

// GET/jobs
func GetJobs(c *gin.Context) {

	db := database.GetDB()

	rows, err := db.Query("SELECT id, company, title, status, date_applied FROM jobs")
	if err != nil {
		log.Println("Query error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch jobs"})
		return
	}
	defer rows.Close()

	var jobs []models.Job

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

// POST/jobs
func CreateJob(c *gin.Context) {
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert job"})
		return
	}

	id, _ := result.LastInsertId()
	newJob.ID = int(id)

	c.JSON(http.StatusCreated, newJob)
}

// PUT/jobs/:id
func UpdateJob(c *gin.Context) {
	id := c.Param("id")

	var updatedJob models.Job
	if err := c.BindJSON(&updatedJob); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	db := database.GetDB()
	_, err := db.Exec(
		"UPDATE jobs SET company = ?, title = ?, status = ?, date_applied = ? WHERE id = ?",
		updatedJob.Company, updatedJob.Title, updatedJob.Status, updatedJob.DateApplied, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update job"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job updated successfully"})

}

// DELETE/jobs/:id
func DeleteJob(c *gin.Context) {
	id := c.Param("id")

	db := database.GetDB()
	_, err := db.Exec("DELETE FROM jobs WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete job"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job deleted successfully"})
}
