package repository

import (
	"context"
	"dataekspor-be/entity"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PartnershipRepository interface {
	InsertPartnership(ctx context.Context, partnership entity.Partnership) (entity.Partnership, error)
	GetAllPartnerships(ctx context.Context) ([]entity.Partnership, error)
	GetPartnershipByID(ctx context.Context, id string) (entity.Partnership, error)
	UpdatePartnershipByID(ctx context.Context, id string, partnership entity.Partnership) (entity.Partnership, error)
	DeletePartnershipByID(ctx context.Context, id string) error
}

type partnershipConnection struct {
	connection *gorm.DB
}

func NewPartnershipRepository(db *gorm.DB) PartnershipRepository {
	return &partnershipConnection{
		connection: db,
	}
}

func (db *partnershipConnection) InsertPartnership(ctx context.Context, partnership entity.Partnership) (entity.Partnership, error) {
	tx := db.connection.WithContext(ctx).Create(&partnership)

	return partnership, tx.Error
}

func (db *partnershipConnection) GetAllPartnerships(ctx context.Context) ([]entity.Partnership, error) {
	var partnerships []entity.Partnership

	tx := db.connection.WithContext(ctx).Find(&partnerships)

	if tx.Error != nil {
		return partnerships, tx.Error
	}

	if len(partnerships) <= 0 {
		return partnerships, errors.New("no partnership found")
	}

	return partnerships, nil
}

func (db *partnershipConnection) GetPartnershipByID(ctx context.Context, id string) (entity.Partnership, error) {
	var partnership entity.Partnership

	tx := db.connection.
		WithContext(ctx).
		Where("id = ?", id).
		Find(&partnership)

	if tx.Error != nil {
		return partnership, tx.Error
	}

	if tx.RowsAffected == 0 {
		return partnership, errors.New("no partnership found")
	}

	return partnership, nil
}

func (db *partnershipConnection) UpdatePartnershipByID(ctx context.Context, id string, partnership entity.Partnership) (entity.Partnership, error) {
	tx := db.
		connection.
		WithContext(ctx).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(&partnership)

	if tx.Error != nil {
		return partnership, tx.Error
	}

	if tx.RowsAffected == 0 {
		return partnership, errors.New("no partnership found on update")
	}

	return partnership, nil
}

func (db *partnershipConnection) DeletePartnershipByID(ctx context.Context, id string) error {
	tx := db.connection.WithContext(ctx).Where("id = ?", id).Delete(&entity.Partnership{})

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("no partnership found on delete")
	}

	return nil
}
