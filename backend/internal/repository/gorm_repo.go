package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"fullstack_LMS/backend/internal/model"
	"fullstack_LMS/backend/pkg/logger_module"
)

type GormRepo struct {
	db     *gorm.DB
	logger *logger_module.Logger
}

func NewGormRepo(db *gorm.DB, logger *logger_module.Logger) *GormRepo {
	return &GormRepo{db: db, logger: logger}
}

func (gr *GormRepo) Create(ctx context.Context, course *model.Course) error {
	return gr.db.WithContext(ctx).Create(course).Error
}

func (gr *GormRepo) GetAll(ctx context.Context) ([]*model.Course, error) {
	var courses []*model.Course
	err := gr.db.WithContext(ctx).Order("created_at DESC").Find(&courses).Error
	return courses, err
}
func (gr *GormRepo) GetByID(ctx context.Context, id uuid.UUID) (*model.Course, error) {
	var course model.Course
	err := gr.db.WithContext(ctx).Where("id = ?", id).First(&course).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &course, nil
}

func (gr *GormRepo) Update(ctx context.Context, course *model.Course) error {
	return gr.db.WithContext(ctx).Save(course).Error
}

func (gr *GormRepo) Delete(ctx context.Context, id uuid.UUID) error {
	delete_course := gr.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Course{})
	if delete_course != nil {
		return delete_course.Error
	}
	if delete_course.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
