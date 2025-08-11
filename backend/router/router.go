package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/syl15/job-tracker/backend/handlers"
)

// Router: defines which URL maps to which handler

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Set trusted proxies 
	r.SetTrustedProxies(nil)

	// CORS config
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	// Enable CORS
	r.Use(cors.New(corsConfig)) // Allows all origins, methods, and headers

	r.GET("/jobs", handlers.GetJobs)
	r.POST("/jobs", handlers.CreateJob)
	r.PUT("/jobs/:id", handlers.UpdateJob)
	r.DELETE("/jobs/:id", handlers.DeleteJob)
	return r
}
