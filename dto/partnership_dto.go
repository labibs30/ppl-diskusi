package dto

import "github.com/google/uuid"

type AddPartnershipDTO struct {
	ProductID   uuid.UUID `json:"productId" binding:"required"`
	UserID      uuid.UUID `json:"userId"`
	TextOffer   string    `json:"textOffer" binding:"required"`
	IsDisetujui int       `json:"isDisetujui" binding:"required"`
}

type UpdatePartnershipDTO struct {
	ProductID   uuid.UUID `json:"productId" binding:"required"`
	UserID      uuid.UUID `json:"userId"`
	TextOffer   string    `json:"textOffer" binding:"required"`
	IsDisetujui int       `json:"isDisetujui" binding:"required"`
}
