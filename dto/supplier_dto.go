package dto

import "gorm.io/datatypes"

type AddSupplierDTO struct {
	Name        string            `json:"name" binding:"required"`
	Description string            `json:"description" binding:"required"`
	CityID      uint64            `json:"cityId" binding:"required"`
	Location    datatypes.JSONMap `gorm:"type:JSONB" json:"location" binding:"required"`
}

type UpdateSupplierDTO struct {
	Name        string            `json:"name" binding:"required"`
	Description string            `json:"description" binding:"required"`
	CityID      uint64            `json:"cityId" binding:"required"`
	Location    datatypes.JSONMap `gorm:"type:JSONB" json:"location" binding:"required"`
}
