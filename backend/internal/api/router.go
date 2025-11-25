package api

import (
	"github.com/gorilla/mux"
)

func (h *CourseHandler) Register(r *mux.Router) {
	r.HandleFunc("/api/courses", h.GetAllCourses).Methods("GET")
	r.HandleFunc("/api/courses", h.CreateCourse).Methods("POST")
	r.HandleFunc("/api/courses/{id}", h.UpdateCourse).Methods("PUT")
	r.HandleFunc("/api/courses/{id}", h.DeleteCourse).Methods("DELETE")
}
