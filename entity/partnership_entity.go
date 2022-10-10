package entity

import "github.com/google/uuid"

type Partnership struct {
	UUIDModel
	ProductID   uuid.UUID `json:"productId" gorm:"not null"`
	Product     Product   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	UserID      uuid.UUID `json:"UserId" gorm:"not null"`
	User        User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	TextOffer   string    `json:"textOffer" gorm:"not null"`
	IsDisetujui int       `json:"isDisetujui" gorm:"not null"`
}
