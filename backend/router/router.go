package router 

import (
	"github.com/gin-gonic/gin"

	"github.com/syl15/job-tracker/backend/handlers"
)

// Router: defines which URL maps to which handler 

func SetupRouter() *gin.Engine {
	r := gin.Default() 
	r.GET("/jobs", handlers.GetJobs)
	r.POST("/jobs", handlers.CreateJob)
	return r
}

// curl -X POST http://localhost:8080/jobs \
//   -H "Content-Type: application/json" \
//   -d '{
//     "company": "Microsoft",
//     "title": "PM",
//     "status": "applied",
//     "date_applied": "2025-07-27"
//   }'