package repository

import (
	"context"
	"dataekspor-be/entity"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SupplierRepository interface {
	InsertSupplier(ctx context.Context, supplier entity.Supplier) (entity.Supplier, error)
	GetAllSupplier(ctx context.Context) ([]entity.Supplier, error)
	GetSupplierByID(ctx context.Context, id string) (entity.Supplier, error)
	UpdateSupplierByID(ctx context.Context, id string, supplier entity.Supplier) (entity.Supplier, error)
	DeleteSupplierByID(ctx context.Context, id string) error
}

type supplierConnection struct {
	connection *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) SupplierRepository {
	return &supplierConnection{
		connection: db,
	}
}

func (db *supplierConnection) InsertSupplier(ctx context.Context, supplier entity.Supplier) (entity.Supplier, error) {
	tx := db.connection.WithContext(ctx).Create(&supplier)

	if supplier.IsZero() {
		return entity.Supplier{}, errors.New("cannot create supplier, supplier inserted is zero value")
	}

	return supplier, tx.Error
}

func (db *supplierConnection) GetAllSupplier(ctx context.Context) ([]entity.Supplier, error) {
	var cities []entity.Supplier

	tx := db.connection.WithContext(ctx).Find(&cities)

	if tx.Error != nil {
		return cities, tx.Error
	}

	if len(cities) <= 0 {
		return cities, errors.New("no supplier found")
	}

	return cities, nil
}

func (db *supplierConnection) GetSupplierByID(ctx context.Context, id string) (entity.Supplier, error) {
	var supplier entity.Supplier

	tx := db.connection.
		WithContext(ctx).
		Where("id = ?", id).
		Find(&supplier)

	if tx.Error != nil {
		return supplier, tx.Error
	}

	if tx.RowsAffected == 0 {
		return supplier, errors.New("no supplier found")
	}

	return supplier, nil
}

func (db *supplierConnection) UpdateSupplierByID(ctx context.Context, id string, supplier entity.Supplier) (entity.Supplier, error) {
	tx := db.
		connection.
		WithContext(ctx).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(&supplier)

	if tx.Error != nil {
		return supplier, tx.Error
	}

	if tx.RowsAffected == 0 {
		return supplier, errors.New("no supplier found on update")
	}

	return supplier, nil
}

func (db *supplierConnection) DeleteSupplierByID(ctx context.Context, id string) error {
	tx := db.connection.WithContext(ctx).Where("id = ?", id).Delete(&entity.Supplier{})

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("no supplier found on delete")
	}

	return nil
}
