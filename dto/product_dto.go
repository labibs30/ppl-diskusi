package dto

import (
	"dataekspor-be/entity"

	"github.com/google/uuid"
)

type AddProductDTO struct {
	Name               string                      `json:"name" binding:"required"`
	Hscode             entity.SliceOfString        `gorm:"type:text" json:"hscode" binding:"required"`
	CategoryID         uint64                      `json:"categoryId" binding:"required"`
	Description        string                      `json:"description" binding:"required"`
	ProductionCapacity entity.ProductionCapacities `gorm:"type:JSONB" json:"productionCapacity" binding:"required"`
	Specification      entity.Specifications       `gorm:"type:JSONB" json:"specification" binding:"required"`
	SupplierID         uuid.UUID                   `json:"supplierId" binding:"required"`
	Certification      entity.Certifications       `gorm:"type:JSONB" json:"certification" binding:"required"`
	Images             entity.SliceOfString        `json:"images" binding:"required"`
}

type UpdateProductDTO struct {
	Name               string                      `json:"name" binding:"required"`
	Hscode             entity.SliceOfString        `gorm:"type:text" json:"hscode" binding:"required"`
	CategoryID         uint64                      `json:"categoryId" binding:"required"`
	Description        string                      `json:"description" binding:"required"`
	ProductionCapacity entity.ProductionCapacities `gorm:"type:JSONB" json:"productionCapacity" binding:"required"`
	Specification      entity.Specifications       `gorm:"type:JSONB" json:"specification" binding:"required"`
	SupplierID         uuid.UUID                   `json:"supplierId" binding:"required"`
	Certification      entity.Certifications       `gorm:"type:JSONB" json:"certification" binding:"required"`
	Images             entity.SliceOfString        `json:"images" binding:"required"`
}
