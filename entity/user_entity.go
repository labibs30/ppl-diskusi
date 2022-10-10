package entity

import "time"

type User struct {
	ID        string        `gorm:"primaryKey;not null" json:"id"`
	Name      string        `json:"name" gorm:"not null"`
	Email     string        `json:"email" gorm:"not null"`
	Phone     string        `json:"telepon" gorm:"not null"`
	Password  string        `json:"-" gorm:"not null"`
	Role      SliceOfString `json:"role" gorm:"type:text;not null"`
	Token     string        `json:"token" gorm:"not null"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	DeletedAt time.Time    `json:"deletedAt"`
}

func (u User) IsZero() bool {
	return u.ID == "" && u.Name == "" && u.Email == "" && u.Phone == "" && u.Password == "" && len(u.Role) == 0 && u.Token == ""
}
