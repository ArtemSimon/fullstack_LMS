package repository

import (
	"context"
	"fullstack_LMS/backend/internal/model"

	"github.com/google/uuid"
)

// Реализуем интерфейс для работы с БД
type CourseRepository interface {
	Create(ctx context.Context, course *model.Course) error
	GetAll(ctx context.Context) ([]*model.Course, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Course, error)
	Update(ctx context.Context, course *model.Course) error
	Delete(ctx context.Context, id uuid.UUID) error
}
