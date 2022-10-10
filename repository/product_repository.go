package repository

import (
	"context"
	"dataekspor-be/entity"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository interface {
	InsertProduct(ctx context.Context, product entity.Product) (entity.Product, error)
	GetAllProducts(ctx context.Context) ([]entity.Product, error)
	GetProductByID(ctx context.Context, id string) (entity.Product, error)
	GetProductByNameOrDesc(ctx context.Context, param string) ([]entity.Product, error)
	UpdateProductByID(ctx context.Context, id string, product entity.Product) (entity.Product, error)
	DeleteProductByID(ctx context.Context, id string) error
}

type productConnection struct {
	connection *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productConnection{
		connection: db,
	}
}

func (db *productConnection) InsertProduct(ctx context.Context, product entity.Product) (entity.Product, error) {
	tx := db.connection.WithContext(ctx).Create(&product)

	if product.IsZero() {
		return entity.Product{}, errors.New("cannot create product, product inserted is zero value")
	}

	return product, tx.Error
}

func (db *productConnection) GetAllProducts(ctx context.Context) ([]entity.Product, error) {
	var products []entity.Product

	tx := db.connection.WithContext(ctx).Find(&products)

	if tx.Error != nil {
		return products, tx.Error
	}

	if len(products) <= 0 {
		return products, errors.New("no product found")
	}

	return products, nil
}

func (db *productConnection) GetProductByID(ctx context.Context, id string) (entity.Product, error) {
	var product entity.Product

	tx := db.connection.WithContext(ctx).Where("id = ?", id).Find(&product)

	if tx.Error != nil {
		return product, tx.Error
	}

	return product, nil
}

func (db *productConnection) GetProductByNameOrDesc(ctx context.Context, param string) ([]entity.Product, error) {
	var products []entity.Product

	tx := db.connection.
		WithContext(ctx).
		Where("name ILIKE ?", param+"%").
		Or(db.connection.Where("description ILIKE ?", param+"%")).
		Find(&products)

	if tx.Error != nil {
		return products, tx.Error
	}

	if (len(products) == 0) || (tx.RowsAffected == 0) {
		return products, errors.New("no product found")
	}

	return products, nil
}

func (db *productConnection) UpdateProductByID(ctx context.Context, id string, product entity.Product) (entity.Product, error) {
	tx := db.
		connection.
		WithContext(ctx).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(&product)

	if tx.Error != nil {
		return product, tx.Error
	}

	if tx.RowsAffected == 0 {
		return product, errors.New("no product found on update")
	}

	return product, nil
}

func (db *productConnection) DeleteProductByID(ctx context.Context, id string) error {
	tx := db.connection.WithContext(ctx).Where("id = ?", id).Delete(&entity.Category{})

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("no category found on delete")
	}

	return nil
}
