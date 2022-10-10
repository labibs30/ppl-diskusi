package entity

type Category struct {
	CommonModel
	Name        string `json:"name" gorm:"unique;not null"`
	Description string `json:"description" gorm:"not null"`
}
