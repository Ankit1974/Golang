package routes

import (
	"log/slog"

	"github.com/gorilla/mux"
	"github.com/imp012-deep/buildingApi/internal/handlers"
	"github.com/imp012-deep/buildingApi/internal/middleware"
	"github.com/imp012-deep/buildingApi/internal/repository"
)

/*
   Route definitions only

   Responsibilities:

   1. Map URL → handler
   2. Attach middleware
*/

// SetupRouter configures all application routes with middleware
func SetupRouter(repo repository.CourseRepository, logger *slog.Logger) *mux.Router {
	router := mux.NewRouter()

	// Initialize handlers with dependency injection
	courseHandler := handlers.NewCourseHandler(repo, logger)

	// Apply global middleware
	router.Use(middleware.CORSMiddleware)
	router.Use(middleware.RecoveryMiddleware(logger))
	router.Use(middleware.LoggingMiddleware(logger))

	// Home route
	router.HandleFunc("/", courseHandler.ServeHome).Methods("GET")

	// API v1 routes
	api := router.PathPrefix("/api/v1").Subrouter()

	// Course endpoints
	api.HandleFunc("/courses", courseHandler.GetAllCourses).Methods("GET")
	api.HandleFunc("/courses/{id}", courseHandler.GetCourse).Methods("GET")
	api.HandleFunc("/courses", courseHandler.CreateCourse).Methods("POST")
	api.HandleFunc("/courses/{id}", courseHandler.UpdateCourse).Methods("PUT")
	api.HandleFunc("/courses/{id}", courseHandler.DeleteCourse).Methods("DELETE")

	return router
}
