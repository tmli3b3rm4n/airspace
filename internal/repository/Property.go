package repository

import (
	"context"
	"errors"
	"github.com/tmli3b3rm4n/airspace/internal/database"
	"gorm.io/gorm"
)

type propertyRepository struct {
	db *gorm.DB
}

// NewPropertyRepository creates a new PropertyRepository instance
func NewPropertyRepository(db *gorm.DB) PropertyRepository {
	return &propertyRepository{db: db}
}

func (r *propertyRepository) GetByID(ctx context.Context, id uint) (*database.Property, error) {
	var property database.Property
	if err := r.db.WithContext(ctx).First(&property, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil if not found
		}
		return nil, err
	}
	return &property, nil
}

func (r *propertyRepository) GetByObjectID(ctx context.Context, objectID int) (*database.Property, error) {
	var property database.Property
	if err := r.db.WithContext(ctx).Where("object_id = ?", objectID).First(&property).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil if not found
		}
		return nil, err
	}
	return &property, nil
}

func (r *propertyRepository) GetAll(ctx context.Context) ([]database.Property, error) {
	var properties []database.Property
	if err := r.db.WithContext(ctx).Find(&properties).Error; err != nil {
		return nil, err
	}
	return properties, nil
}

func (r *propertyRepository) Create(ctx context.Context, property *database.Property) error {
	if err := r.db.WithContext(ctx).Create(property).Error; err != nil {
		return err
	}
	return nil
}

func (r *propertyRepository) Update(ctx context.Context, property *database.Property) error {
	if err := r.db.WithContext(ctx).Save(property).Error; err != nil {
		return err
	}
	return nil
}

func (r *propertyRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&database.Property{}, id).Error; err != nil {
		return err
	}
	return nil
}
