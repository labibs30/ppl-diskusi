package dto

type AddCityDTO struct {
	Name     string `json:"name" binding:"required"`
	Province string `json:"province" binding:"required"`
}

type UpdateCityDTO struct {
	Name     string `json:"name" binding:"required"`
	Province string `json:"province" binding:"required"`
}
