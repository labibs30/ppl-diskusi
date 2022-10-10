package entity

type City struct {
	CommonModel
	Name     string `json:"name" gorm:"not null"`
	Province string `json:"province" gorm:"not null"`
}
