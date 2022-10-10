package repository

import (
	"context"
	"dataekspor-be/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CategoryRepository interface {
	InsertCategory(ctx context.Context, category entity.Category) (entity.Category, error)
	GetAllCategory(ctx context.Context) ([]entity.Category, error)
	GetCategoryByID(ctx context.Context, id string) (entity.Category, error)
	GetCategoryByNameOrDesc(ctx context.Context, param string) ([]entity.Category, error)
	UpdateCategoryByID(ctx context.Context, id string, category entity.Category) (entity.Category, error)
	DeleteCategoryByID(ctx context.Context, id string) error
}

type categoryConnection struct {
	connection *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryConnection{
		connection: db,
	}
}

func (db *categoryConnection) GetAllCategory(ctx context.Context) ([]entity.Category, error) {
	var categories []entity.Category

	tx := db.connection.WithContext(ctx).Find(&categories)

	if tx.Error != nil {
		return categories, tx.Error
	}

	if len(categories) <= 0 {
		return categories, errors.New("no category found")
	}

	return categories, nil
}

func (db *categoryConnection) InsertCategory(ctx context.Context, category entity.Category) (entity.Category, error) {
	fmt.Println("sebelum masuk:", category)

	tx := db.connection.WithContext(ctx).Create(&category)

	if category == (entity.Category{}) {
		return entity.Category{}, errors.New("cannot create category, category inserted is zero value")
	}

	return category, tx.Error
}

func (db *categoryConnection) GetCategoryByID(ctx context.Context, id string) (entity.Category, error) {
	var category entity.Category

	tx := db.connection.
		WithContext(ctx).
		Where("id = ?", id).
		Find(&category)

	if tx.Error != nil {
		return category, tx.Error
	}

	if tx.RowsAffected == 0 {
		return category, errors.New("no category found")
	}

	return category, nil
}

func (db *categoryConnection) GetCategoryByNameOrDesc(ctx context.Context, param string) ([]entity.Category, error) {
	var category []entity.Category

	tx := db.connection.
		WithContext(ctx).
		Where("name like ?", param).
		Or(db.connection.Where("description like ?", param)).
		Find(&category)

	if tx.Error != nil {
		return category, tx.Error
	}

	if (len(category) == 0) || (tx.RowsAffected == 0) {
		return category, errors.New("no category found")
	}

	return category, nil
}

func (db *categoryConnection) UpdateCategoryByID(ctx context.Context, id string, category entity.Category) (entity.Category, error) {
	tx := db.
		connection.
		WithContext(ctx).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(&category)

	if tx.Error != nil {
		return category, tx.Error
	}

	if tx.RowsAffected == 0 {
		return category, errors.New("no category found on update")
	}

	return category, nil
}

func (db *categoryConnection) DeleteCategoryByID(ctx context.Context, id string) error {
	tx := db.connection.WithContext(ctx).Where("id = ?", id).Delete(&entity.Category{})

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("no category found on delete")
	}

	return nil
}
