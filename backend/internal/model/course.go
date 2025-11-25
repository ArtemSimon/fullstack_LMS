package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var validate = validator.New()

// Модель курса
type Course struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Title       string    `gorm:"type:text;not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Author      string    `gorm:"type:text;not null" json:"author"`
	CreatedAt   time.Time `gorm:"<-:create" json:"created_at"`
	UpdatedAt   time.Time `gorm:"" json:"updated_at"`
}

// Модель создания курса
type CreateCourse struct {
	Title       string `json:"title" validate:"required,min=3"`
	Description string `json:"description"`
	Author      string `json:"author" validate:"required"`
}

// Модель для редактирования
type UpdateCourse struct {
	Title       string `json:"title" validate:"required,min=3"`
	Description string `json:"description"`
	Author      string `json:"author" validate:"required"`
}
