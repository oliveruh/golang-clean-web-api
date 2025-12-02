package repository

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/domain/filter"
	"github.com/naeemaei/golang-clean-web-api/domain/model"
)

type BaseRepository[TEntity any] interface {
	Create(ctx context.Context, entity TEntity) (TEntity, error)
	Update(ctx context.Context, id int, entity map[string]interface{}) (TEntity, error)
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (TEntity, error)
	GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (int64, *[]TEntity, error)
}
type CountryRepository interface {
	BaseRepository[model.Country]
}

type CityRepository interface {
	BaseRepository[model.City]
	// Create(ctx context.Context, City model.City) (model.City, error)
	// Update(ctx context.Context, id int, City model.City) (model.City, error)
	// Delete(ctx context.Context, id int) error
	// GetById(ctx context.Context, id int) (model.City, error)
	// GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (int64, *[]model.City, error)
}

type ColorRepository interface {
	BaseRepository[model.Color]
}

type CompanyRepository interface {
	BaseRepository[model.Company]
}
