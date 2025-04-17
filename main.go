package main

import (
	"resumeorganizer/handler"
	"resumeorganizer/service"
	"resumeorganizer/store"

	"gofr.dev/pkg/gofr"
)

func main() {
	// Create a new GoFr application
	app := gofr.New()

	// Initialize store
	resumeStore := store.NewInMemoryResumeStore()

	// Initialize service
	resumeService := service.NewResumeService(resumeStore)

	// Initialize handler
	resumeHandler := handler.NewResumeHandler(resumeService)

	// Register routes
	app.POST("/resumes", resumeHandler.Create)
	app.GET("/resumes", resumeHandler.GetAll)
	app.GET("/resumes/{id}", resumeHandler.Get)
	app.PATCH("/resumes/{id}/status", resumeHandler.UpdateStatus)
	app.POST("/resumes/{id}/upload", resumeHandler.UploadFile)
	app.DELETE("/resumes/{id}", resumeHandler.Delete)

	// Start the server
	app.Run()
}
