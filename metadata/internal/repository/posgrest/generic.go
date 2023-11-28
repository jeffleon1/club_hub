package posgrest

import (
	"context"

	"github.com/jeffleon1/club_hub/metadata/pkg/entities"
	"gorm.io/gorm"
)

type repository[T interface{}] struct {
	db        *gorm.DB
	relations string
}

func New[T interface{}](db *gorm.DB, relations string) entities.Repository[T] {
	return &repository[T]{
		db,
		relations,
	}
}

func (r *repository[T]) Create(ctx context.Context, entity *T) error {
	return r.db.WithContext(ctx).Create(&entity).Error
}

func (r *repository[T]) GetAll(ctx context.Context) (*[]T, error) {
	var entities []T
	err := r.db.WithContext(ctx).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return &entities, nil
}

func (r *repository[T]) GetByID(ctx context.Context, id uint) (*T, error) {
	var entity T
	if err := r.db.WithContext(ctx).First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *repository[T]) GetBy(ctx context.Context, key string, value interface{}) (*[]T, error) {
	var entity []T
	if err := r.db.WithContext(ctx).Preload(r.relations).Where(key, value).Find(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *repository[T]) Update(ctx context.Context, entity *T, id uint) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Updates(entity).Error
}

func (r *repository[T]) Delete(ctx context.Context, id uint) error {
	var entity T
	return r.db.WithContext(ctx).Delete(&entity, id).Error
}
