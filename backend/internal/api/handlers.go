package api

import (
	"context"
	"encoding/json"
	"errors"
	"fullstack_LMS/backend/internal/model"
	"fullstack_LMS/backend/internal/service"
	"fullstack_LMS/backend/pkg/logger_module"
	"fullstack_LMS/backend/pkg/validator"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type CourseHandler struct {
	service service.ServiceCourseI
	logger  *logger_module.Logger
}

func NewCourseHandler(service service.ServiceCourseI, logger *logger_module.Logger) *CourseHandler {
	return &CourseHandler{service: service, logger: logger}
}

// Ответ с ошибкой в унифицированном формате
type ErrorResponse struct {
	Error string `json:"error"`
}

func sendError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func renderJSON(w http.ResponseWriter, code int, object interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(object)
}

func (h *CourseHandler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var req model.Course
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("invalid JSON", "error", err)
		renderJSON(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	// Валидация
	if err := validator.ValidateStruct(req); err != nil {
		h.logger.Error("validation failed", "error", err)
		renderJSON(w, http.StatusBadRequest, err)
		return
	}

	// course := &model.Course{
	// 	ID:          uuid.New(),
	// 	Title:       req.Title,
	// 	Description: req.Description,
	// 	Author:      req.Author,
	// }

	course, err := h.service.CreateCourse(ctx, &req)
	if err != nil {
		h.logger.Error("failed to create course", "error", err)
		renderJSON(w, http.StatusInternalServerError, "failed to create course")
		return
	}

	h.logger.Info("course created", "id", course.ID)
	renderJSON(w, http.StatusCreated, course) // ← теперь
}

func (handler *CourseHandler) GetAllCourses(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	handler.logger.Debug("Fetching all courses")

	courses, err := handler.service.GetAllCourses(ctx)
	if err != nil {
		handler.logger.Error("Failed to fetch courses", "error", err)
		sendError(w, http.StatusInternalServerError, "failed to fetch courses")
		return
	}

	handler.logger.Info("Courses fetched successfully", "count", len(courses))
	renderJSON(w, http.StatusOK, courses)
}

func (handler *CourseHandler) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	courseIDStr := mux.Vars(r)["id"]
	handler.logger.Debug("Starting to process update course request", "course_id", courseIDStr)

	// Парсинг UUID
	courseID, err := uuid.Parse(courseIDStr)
	if err != nil {
		handler.logger.Error("Invalid course ID format", "error", err, "id", courseIDStr)
		sendError(w, http.StatusBadRequest, "invalid course ID format")
		return
	}

	var req model.UpdateCourse
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler.logger.Error("Failed to decode request body", "error", err)
		sendError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	handler.logger.Debug("Request body decoded for update",
		"title", req.Title,
		"author", req.Author,
		"course_id", courseID)

	// Валидация
	if err := validator.ValidateStruct(req); err != nil {
		handler.logger.Error("Validation failed", "error", err)
		sendError(w, http.StatusBadRequest, "Validation failed")
		return
	}

	course := &model.Course{
		ID:          courseID,
		Title:       req.Title,
		Description: req.Description,
		Author:      req.Author,
	}

	handler.logger.Info("Updating course", "id", course.ID, "title", course.Title)

	if err := handler.service.UpdateCourse(ctx, course.ID, course); err != nil {
		handler.logger.Error("Course not found for update", "id", course.ID)
		sendError(w, http.StatusNotFound, "course not found")
		return

	}

	handler.logger.Info("Course updated successfully", "id", course.ID)
	renderJSON(w, http.StatusOK, course)
}

func (handler *CourseHandler) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	courseIDStr := mux.Vars(r)["id"]
	handler.logger.Debug("Starting to process delete course request", "course_id", courseIDStr)

	// Парсинг UUID
	courseID, err := uuid.Parse(courseIDStr)
	if err != nil {
		handler.logger.Error("Invalid course ID format", "error", err, "id", courseIDStr)
		sendError(w, http.StatusBadRequest, "invalid course ID format")
		return
	}

	handler.logger.Info("Deleting course", "id", courseID)

	if err := handler.service.DeleteCourse(ctx, courseID); err != nil {
		// Проверяем: ошибка "запись не найдена" от GORM?
		if errors.Is(err, gorm.ErrRecordNotFound) {
			handler.logger.Error("Course not found for deletion", "id", courseID)
			sendError(w, http.StatusNotFound, "course not found")
			return
		}

		// Любая другая ошибка — 500
		handler.logger.Error("Failed to delete course", "error", err, "id", courseID)
		sendError(w, http.StatusInternalServerError, "failed to delete course")
		return
	}

	handler.logger.Info("Course deleted successfully", "id", courseID)
	renderJSON(w, http.StatusOK, map[string]bool{"success": true})
}
