package repository

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/imp012-deep/buildingApi/internal/models"
)

/*
   Data access layer

   Responsibilities:

     1. Database queries
     2. External API calls
     3. Cache access
*/

/*
   Why this layer matters:

    . DB can change (Postgres → Mongo)
    . Easy to mock in tests
    . Keeps SQL out of handlers

     💡 This is where real backend engineers stand out
*/

var (
	// ErrCourseNotFound is returned when a course is not found
	ErrCourseNotFound = errors.New("course not found")

	// ErrCourseExists is returned when trying to create a course with existing ID
	ErrCourseExists = errors.New("course already exists")
)

// CourseRepository defines the interface for course data access
type CourseRepository interface {
	GetAll(ctx context.Context) ([]models.Course, error)
	GetByID(ctx context.Context, id string) (*models.Course, error)
	Create(ctx context.Context, course *models.Course) error
	Update(ctx context.Context, id string, course *models.Course) error
	Delete(ctx context.Context, id string) error
	Exists(ctx context.Context, id string) bool
}

// InMemoryCourseRepository implements CourseRepository using in-memory storage
// Thread-safe implementation using RWMutex
type InMemoryCourseRepository struct {
	mu      sync.RWMutex
	courses map[string]models.Course
}

// NewInMemoryCourseRepository creates a new in-memory course repository
func NewInMemoryCourseRepository() *InMemoryCourseRepository {
	return &InMemoryCourseRepository{
		courses: make(map[string]models.Course),
	}
}

// GetAll retrieves all courses
func (r *InMemoryCourseRepository) GetAll(ctx context.Context) ([]models.Course, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Check context cancellation
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	courses := make([]models.Course, 0, len(r.courses))
	for _, course := range r.courses {
		courses = append(courses, course)
	}

	return courses, nil
}

// GetByID retrieves a course by ID
func (r *InMemoryCourseRepository) GetByID(ctx context.Context, id string) (*models.Course, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Check context cancellation
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	course, exists := r.courses[id]
	if !exists {
		return nil, ErrCourseNotFound
	}

	return &course, nil
}

// Create adds a new course
func (r *InMemoryCourseRepository) Create(ctx context.Context, course *models.Course) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check context cancellation
	if ctx.Err() != nil {
		return ctx.Err()
	}

	if _, exists := r.courses[course.ID]; exists {
		return ErrCourseExists
	}

	r.courses[course.ID] = *course
	return nil
}

// Update modifies an existing course
func (r *InMemoryCourseRepository) Update(ctx context.Context, id string, course *models.Course) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check context cancellation
	if ctx.Err() != nil {
		return ctx.Err()
	}

	if _, exists := r.courses[id]; !exists {
		return ErrCourseNotFound
	}

	// Preserve the original ID
	course.ID = id
	r.courses[id] = *course

	return nil
}

// Delete removes a course
func (r *InMemoryCourseRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check context cancellation
	if ctx.Err() != nil {
		return ctx.Err()
	}

	if _, exists := r.courses[id]; !exists {
		return ErrCourseNotFound
	}

	delete(r.courses, id)
	return nil
}

// Exists checks if a course exists
func (r *InMemoryCourseRepository) Exists(ctx context.Context, id string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, exists := r.courses[id]
	return exists
}

// Seed initializes the repository with sample data
func (r *InMemoryCourseRepository) Seed() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	sampleCourses := []models.Course{
		{
			ID:    "2",
			Name:  "react js",
			Price: 299,
			Author: &models.Author{
				FullName: "Ankit Raj",
				Website:  "https://learnonline.in",
			},
		},
		{
			ID:    "3",
			Name:  "react js3",
			Price: 2993,
			Author: &models.Author{
				FullName: "Ankit Raj3",
				Website:  "https://learnonline.in",
			},
		},
	}

	for _, course := range sampleCourses {
		r.courses[course.ID] = course
	}

	return nil
}

// Count returns the total number of courses
func (r *InMemoryCourseRepository) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.courses)
}

// String returns a string representation for debugging
func (r *InMemoryCourseRepository) String() string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return fmt.Sprintf("InMemoryCourseRepository{count: %d}", len(r.courses))
}
