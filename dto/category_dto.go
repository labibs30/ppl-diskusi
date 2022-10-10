package dto

type AddCategoryDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateCategoryDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
