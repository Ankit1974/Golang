package handlers

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/imp012-deep/buildingApi/internal/models"
	"github.com/imp012-deep/buildingApi/internal/repository"
	"github.com/imp012-deep/buildingApi/pkg/response"
)

/*
   HTTP request handlers

   Responsibilities:

   1. Read request (JSON, params)
   2. Validate input
   3. Call business logic / repository
   4. Return HTTP response
*/

/* 🚫 What should NOT be here:

   SQL queries
   Business rules
   Complex logic
*/

/*
    In a small or medium API (like this one), yes, the logic often sits in the handler. However,
	as an app grows, senior engineers follow the "Thin Handler, Thick Service" rule:

    Handler (Controller): Should only handle "HTTP logic" (status codes, JSON parsing).
    Service Layer (Business Logic): A separate folder (internal/service) where the actual business rules
	live (e.g., "A student cannot enroll in more than 5 courses").

    Repository (Data Logic): Handles only database queries (SQL, NoSQL
*/

// CourseHandler handles HTTP requests for course resources
type CourseHandler struct {
	repo   repository.CourseRepository
	logger *slog.Logger
}

// NewCourseHandler creates a new course handler with dependency injection
func NewCourseHandler(repo repository.CourseRepository, logger *slog.Logger) *CourseHandler {
	return &CourseHandler{
		repo:   repo,
		logger: logger,
	}
}

// ServeHome handles the home endpoint
func (h *CourseHandler) ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1>Course API - Production Ready</h1><p>API endpoints available at /api/v1/courses</p>"))
}

// GetAllCourses retrieves all courses
// GET /api/v1/courses
func (h *CourseHandler) GetAllCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := h.repo.GetAll(r.Context())
	if err != nil {
		h.logger.Error("Failed to get all courses", slog.Any("error", err))
		response.InternalServerError(w, "Failed to retrieve courses")
		return
	}

	response.OK(w, courses)
}

// GetCourse retrieves a single course by ID
// GET /api/v1/courses/{id}
func (h *CourseHandler) GetCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		response.BadRequest(w, "Course ID is required")
		return
	}

	course, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		if err == repository.ErrCourseNotFound {
			response.NotFound(w, fmt.Sprintf("Course with ID %s not found", id))
			return
		}
		h.logger.Error("Failed to get course", slog.String("id", id), slog.Any("error", err))
		response.InternalServerError(w, "Failed to retrieve course")
		return
	}

	response.OK(w, course)
}

// CreateCourse creates a new course
// POST /api/v1/courses
func (h *CourseHandler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	// Check if request body is empty
	if r.Body == nil {
		response.BadRequest(w, "Request body is required")
		return
	}
	defer r.Body.Close()

	// Decode the request body
	var course models.Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		response.BadRequest(w, "Invalid JSON format")
		return
	}

	// Validate the course data
	if valid, errMsg := course.Validate(); !valid {
		response.BadRequest(w, errMsg)
		return
	}

	// Generate a unique ID (cryptographically secure random number)
	id, err := generateSecureID()
	if err != nil {
		h.logger.Error("Failed to generate ID", slog.Any("error", err))
		response.InternalServerError(w, "Failed to generate course ID")
		return
	}
	course.ID = id

	// Create the course
	if err := h.repo.Create(r.Context(), &course); err != nil {
		if err == repository.ErrCourseExists {
			response.BadRequest(w, "Course with this ID already exists")
			return
		}
		h.logger.Error("Failed to create course", slog.Any("error", err))
		response.InternalServerError(w, "Failed to create course")
		return
	}

	h.logger.Info("Course created", slog.String("id", course.ID), slog.String("name", course.Name))
	response.Created(w, course)
}

// UpdateCourse updates an existing course
// PUT /api/v1/courses/{id}
func (h *CourseHandler) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		response.BadRequest(w, "Course ID is required")
		return
	}

	// Check if request body is empty
	if r.Body == nil {
		response.BadRequest(w, "Request body is required")
		return
	}
	defer r.Body.Close()

	// Decode the request body
	var course models.Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		response.BadRequest(w, "Invalid JSON format")
		return
	}

	// Validate the course data
	if valid, errMsg := course.Validate(); !valid {
		response.BadRequest(w, errMsg)
		return
	}

	// Update the course
	if err := h.repo.Update(r.Context(), id, &course); err != nil {
		if err == repository.ErrCourseNotFound {
			response.NotFound(w, fmt.Sprintf("Course with ID %s not found", id))
			return
		}
		h.logger.Error("Failed to update course", slog.String("id", id), slog.Any("error", err))
		response.InternalServerError(w, "Failed to update course")
		return
	}

	h.logger.Info("Course updated", slog.String("id", id), slog.String("name", course.Name))
	response.OK(w, course)
}

// DeleteCourse deletes a course by ID
// DELETE /api/v1/courses/{id}
func (h *CourseHandler) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		response.BadRequest(w, "Course ID is required")
		return
	}

	// Delete the course
	if err := h.repo.Delete(r.Context(), id); err != nil {
		if err == repository.ErrCourseNotFound {
			response.NotFound(w, fmt.Sprintf("Course with ID %s not found", id))
			return
		}
		h.logger.Error("Failed to delete course", slog.String("id", id), slog.Any("error", err))
		response.InternalServerError(w, "Failed to delete course")
		return
	}

	h.logger.Info("Course deleted", slog.String("id", id))
	response.SuccessWithMessage(w, http.StatusOK, fmt.Sprintf("Course with ID %s deleted successfully", id), nil)
}

// generateSecureID generates a cryptographically secure random ID
func generateSecureID() (string, error) {
	// Generate a random number between 1 and 999999
	n, err := rand.Int(rand.Reader, big.NewInt(999999))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", n.Int64()+1), nil
}
