package repository

import (
	"context"
	"dataekspor-be/entity"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CityRepository interface {
	InsertCity(ctx context.Context, city entity.City) (entity.City, error)
	GetAllCity(ctx context.Context) ([]entity.City, error)
	GetCityByID(ctx context.Context, id string) (entity.City, error)
	UpdateCityByID(ctx context.Context, id string, city entity.City) (entity.City, error)
	DeleteCityByID(ctx context.Context, id string) error
}

type cityConnection struct {
	connection *gorm.DB
}

func NewCityRepository(db *gorm.DB) CityRepository {
	return &cityConnection{
		connection: db,
	}
}

func (db *cityConnection) InsertCity(ctx context.Context, city entity.City) (entity.City, error) {
	tx := db.connection.WithContext(ctx).Create(&city)

	if city == (entity.City{}) {
		return entity.City{}, errors.New("cannot create city, city inserted is zero value")
	}

	return city, tx.Error
}

func (db *cityConnection) GetAllCity(ctx context.Context) ([]entity.City, error) {
	var cities []entity.City

	tx := db.connection.WithContext(ctx).Find(&cities)

	if tx.Error != nil {
		return cities, tx.Error
	}

	if len(cities) <= 0 {
		return cities, errors.New("no city found")
	}

	return cities, nil
}

func (db *cityConnection) GetCityByID(ctx context.Context, id string) (entity.City, error) {
	var city entity.City

	tx := db.connection.
		WithContext(ctx).
		Where("id = ?", id).
		Find(&city)

	if tx.Error != nil {
		return city, tx.Error
	}

	if tx.RowsAffected == 0 {
		return city, errors.New("no city found")
	}

	return city, nil
}

func (db *cityConnection) UpdateCityByID(ctx context.Context, id string, city entity.City) (entity.City, error) {
	tx := db.
		connection.
		WithContext(ctx).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(&city)

	if tx.Error != nil {
		return city, tx.Error
	}

	if tx.RowsAffected == 0 {
		return city, errors.New("no city found on update")
	}

	return city, nil
}

func (db *cityConnection) DeleteCityByID(ctx context.Context, id string) error {
	tx := db.connection.WithContext(ctx).Where("id = ?", id).Delete(&entity.City{})

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("no city found on delete")
	}

	return nil
}
