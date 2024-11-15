package repository

import (
	"context"
)

type PropertyRepository interface {
	GetByID(ctx context.Context, id uint) (*models.Property, error)
	GetByObjectID(ctx context.Context, objectID int) (*models.Property, error)
	GetAll(ctx context.Context) ([]models.Property, error)
	Create(ctx context.Context, property *models.Property) error
	Update(ctx context.Context, property *models.Property) error
	Delete(ctx context.Context, id uint) error
}
