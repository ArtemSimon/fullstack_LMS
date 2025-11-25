package service

import (
	"context"
	"fmt"
	"fullstack_LMS/backend/internal/model"
	"fullstack_LMS/backend/internal/repository"
	"fullstack_LMS/backend/pkg/logger_module"
	"fullstack_LMS/backend/pkg/validator"

	"github.com/google/uuid"
)

type ServiceCourseI interface {
	CreateCourse(ctx context.Context, model *model.Course) (*model.Course, error)
	GetAllCourses(ctx context.Context) ([]*model.Course, error)
	UpdateCourse(ctx context.Context, id uuid.UUID, model *model.Course) error
	DeleteCourse(ctx context.Context, id uuid.UUID) error
}
type CourseService struct {
	repo   repository.CourseRepository
	logger *logger_module.Logger
}

func NewCourseService(repo repository.CourseRepository, logger *logger_module.Logger) *CourseService {
	return &CourseService{repo: repo, logger: logger}
}

func (s *CourseService) CreateCourse(ctx context.Context, model_course *model.Course) (*model.Course, error) {
	if err := validator.ValidateStruct(model_course); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// 2. Создаём сущность с бизнес-правилами
	course := &model.Course{
		ID:          uuid.New(), // ← Генерация ID здесь — в service
		Title:       model_course.Title,
		Description: model_course.Description,
		Author:      model_course.Author,
		// CreatedAt/UpdatedAt — GORM проставит сам
	}

	// 3. Сохраняем через репозиторий
	if err := s.repo.Create(ctx, course); err != nil {
		return nil, fmt.Errorf("repo create failed: %w", err)
	}

	return course, nil
}

func (s *CourseService) GetAllCourses(ctx context.Context) ([]*model.Course, error) {
	return s.repo.GetAll(ctx)
}

func (s *CourseService) UpdateCourse(ctx context.Context, id uuid.UUID, model *model.Course) error {
	if err := validator.ValidateStruct(model); err != nil {
		return err
	}

	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err // пробрасываем gorm.ErrRecordNotFound → 404 в handler'е
	}

	existing.Title = model.Title
	existing.Description = model.Description
	existing.Author = model.Author

	return s.repo.Update(ctx, existing)
}

func (s *CourseService) DeleteCourse(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
