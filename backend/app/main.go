package main

import (
	"context"
	"fullstack_LMS/backend/internal/api"
	"fullstack_LMS/backend/internal/config"
	"fullstack_LMS/backend/internal/repository"
	"fullstack_LMS/backend/internal/service"
	"fullstack_LMS/backend/pkg/logger_module"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	log_file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed open log file", err)
	}
	defer log_file.Close()

	miltu_writer := io.MultiWriter(log_file, os.Stdout)

	logger := logger_module.New(miltu_writer, "[APP]", log.LstdFlags|log.Lshortfile)

	logger.Info("Logger start")

	// 2. Загружаем конфиг
	conf, err := config.Load_Config_PG(logger)
	if err != nil {
		logger.Fatal("Failed to load config", "error", err)
	}

	db, err := repository.NewConnectPostgresDB(logger, conf)
	if err != nil {
		log.Fatal("Failed to connect to database", "error", err)
	}
	// 4. Инициализируем слои
	courseRepo := repository.NewGormRepo(db, logger)
	courseService := service.NewCourseService(courseRepo, logger)
	courseHandler := api.NewCourseHandler(courseService, logger)

	// 5. Настраиваем роутер
	r := mux.NewRouter()
	r.HandleFunc("/api/courses", courseHandler.CreateCourse).Methods("POST")
	r.HandleFunc("/api/courses", courseHandler.GetAllCourses).Methods("GET")
	r.HandleFunc("/api/courses/{id}", courseHandler.UpdateCourse).Methods("PUT")
	r.HandleFunc("/api/courses/{id}", courseHandler.DeleteCourse).Methods("DELETE")

	// 6. Запускаем сервер
	server := &http.Server{
		Addr:    ":" + conf.Http_Port,
		Handler: r,
	}

	go func() {
		logger.Info("🚀 Server starting", "addr", server.Addr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			logger.Error("Server failed", "error", err)
			os.Exit(1)
		}
	}()

	// 7. Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("⏳ Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown", "error", err)
		os.Exit(1)
	}
	logger.Info("✅ Server stopped")
}
