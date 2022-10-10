package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type Product struct {
	UUIDModel
	Name               string               `json:"name" gorm:"not null"`
	Hscode             SliceOfString        `gorm:"type:text;not null" json:"hscode"`
	CategoryID         uint64               `json:"categoryId" gorm:"not null"`
	Category           Category             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Description        string               `json:"description" gorm:"not null"`
	ProductionCapacity ProductionCapacities `gorm:"type:JSONB" json:"productionCapacity"`
	Specification      Specifications       `gorm:"type:JSONB" json:"specification"`
	SupplierID         uuid.UUID            `json:"supplierId" gorm:"not null"`
	Supplier           Supplier             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Certification      Certifications       `gorm:"type:JSONB" json:"certification"`
	Images             SliceOfString        `gorm:"type:text;not null" json:"images"`
}

func (p Product) IsZero() bool {
	return (p.Name == "") &&
		(len(p.Hscode) <= 0) &&
		(p.CategoryID == 0) &&
		(p.Description == "") &&
		(p.ProductionCapacity == nil) &&
		(p.Specification == nil) &&
		(p.SupplierID == uuid.Nil) &&
		(p.Certification == nil) &&
		(p.ID == uuid.Nil) &&
		(p.CreatedAt.IsZero()) &&
		(p.UpdatedAt.IsZero())
}

type ProductionCapacity struct {
	Size              string `json:"size"`
	TimeManufacturing string `json:"timeManufacturing"`
}

type ProductionCapacities []ProductionCapacity

func (p *ProductionCapacities) Scan(src any) error {
	switch v := src.(type) {
	case []byte:
		return json.Unmarshal(v, p)
	case string:
		return json.Unmarshal([]byte(v), p)
	}
	return errors.New("type assertion production capacities failed")
}

func (p ProductionCapacities) Value() (driver.Value, error) {
	return json.Marshal(p)
}

type Specification struct {
	Key    string `json:"key" binding:"required"`
	Values string `json:"value" binding:"required"`
}

type Specifications []Specification

func (s *Specifications) Scan(src any) error {
	switch v := src.(type) {
	case []byte:
		return json.Unmarshal(v, s)
	case string:
		return json.Unmarshal([]byte(v), s)
	}
	return errors.New("type assertion specifications failed")
}

func (s Specifications) Value() (driver.Value, error) {
	return json.Marshal(s)
}

type Certification struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	File        string `json:"file" binding:"required"`
}

type Certifications []Certification

func (c *Certifications) Scan(src any) error {
	switch v := src.(type) {
	case []byte:
		return json.Unmarshal(v, c)
	case string:
		return json.Unmarshal([]byte(v), c)
	}
	return errors.New("type assertion certifications failed")
}

func (c Certifications) Value() (driver.Value, error) {
	return json.Marshal(c)
}
