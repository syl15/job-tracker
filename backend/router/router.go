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
	r.PUT("/jobs/:id", handlers.UpdateJob)
	r.DELETE("/jobs/:id", handlers.DeleteJob)
	return r
}
