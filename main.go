package main

import (
	"resumeorganizer/handler"
	"resumeorganizer/migrations"
	"resumeorganizer/service"
	"resumeorganizer/store"

	"gofr.dev/pkg/gofr"
)

func main() {
	// Create a new application
	app := gofr.New()

	// Add migrations to run
	app.Migrate(migrations.All())

	// Create store and service
	resumeStore := store.NewResumeStore()
	resumeService := service.NewResumeService(resumeStore)
	resumeHandler := handler.NewResumeHandler(resumeService)

	// Register routes
	app.POST("/resumes", resumeHandler.Create)
	app.GET("/resumes/{id}", resumeHandler.Get)
	app.GET("/resumes", resumeHandler.GetAll)
	app.PUT("/resumes/{id}/status", resumeHandler.UpdateStatus)
	app.DELETE("/resumes/{id}", resumeHandler.Delete)

	// Start the server
	app.Run()
}
