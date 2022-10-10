package entity

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Supplier struct {
	UUIDModel
	Name        string            `json:"name" gorm:"not null"`
	Description string            `json:"description" gorm:"not null"`
	CityID      uint64            `json:"idCity" gorm:"not null"`
	City        City              `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Location    datatypes.JSONMap `gorm:"type:JSONB" json:"location" `
}

func (s Supplier) IsZero() bool {
	return (s.Description == "") &&
		(s.CityID == 0) &&
		(s.Location == nil) &&
		(s.ID == uuid.Nil) &&
		(s.CreatedAt.IsZero()) &&
		(s.UpdatedAt.IsZero())
}
