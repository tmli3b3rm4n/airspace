package repository

import (
	"context"
	"github.com/tmli3b3rm4n/airspace/internal/database"
)

type IRestrictedAreaRepository interface {
	Create(ctx context.Context, area *database.RestrictedArea) error
	FindByID(ctx context.Context, id uint) (*database.RestrictedArea, error)
	FindAll(ctx context.Context) ([]database.RestrictedArea, error)
	Update(ctx context.Context, area *database.RestrictedArea) error
	Delete(ctx context.Context, id uint) error
}

type RestrictedAreaRepository struct {
	RestrictedArea *database.RestrictedArea
	Db             database.Database
}

func NewRestrictedAreaRepository(db database.Database) *RestrictedAreaRepository {
	property := &database.Property{}
	restrictedArea := &database.RestrictedArea{
		Property: property,
	}

	return &RestrictedAreaRepository{
		Db:             db,
		RestrictedArea: restrictedArea,
	}
}

func (r *RestrictedAreaRepository) Create(ctx context.Context, area *database.RestrictedArea) error {
	return r.Db.Create(area).Error
}

func (r *RestrictedAreaRepository) FindByID(ctx context.Context, id uint) (*database.RestrictedArea, error) {
	var area database.RestrictedArea
	if err := r.Db.First(&area, id).Error; err != nil {
		return nil, err
	}
	return &area, nil
}

func (r *RestrictedAreaRepository) FindAll(ctx context.Context) ([]database.RestrictedArea, error) {
	var areas []database.RestrictedArea
	if err := r.Db.Find(&areas).Error; err != nil {
		return nil, err
	}
	return areas, nil
}

func (r *RestrictedAreaRepository) Update(ctx context.Context, area *database.RestrictedArea) error {
	return r.Db.Save(area).Error
}

func (r *RestrictedAreaRepository) Delete(ctx context.Context, id uint) error {
	return r.Db.Where("id = ?", id).Delete(&database.RestrictedArea{}).Error
}
