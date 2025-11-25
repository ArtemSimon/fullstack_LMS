package repository

import (
	"fmt"
	"fullstack_LMS/backend/internal/config"
	"fullstack_LMS/backend/pkg/logger_module"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnectPostgresDB(logger *logger_module.Logger, config *config.Config_PG) (*gorm.DB, error) {

	connection_db := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
		config.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(connection_db), &gorm.Config{})
	if err != nil {
		logger.Fatal("Failed to open DB", "error", err)
	}

	return db, nil
}
