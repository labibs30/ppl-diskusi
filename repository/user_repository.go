package repository

import (
	"dataekspor-be/common/algorithm"
	"dataekspor-be/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllEksportir() []entity.User
	InsertUser(user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	VerifyCredential(email string) any
	IsDuplicateEmail(email string) (tx *gorm.DB)
	// FindByEmail(email string) entity.User
	FindUserByID(UserID string) entity.User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user entity.User) entity.User {
	db.connection.Save(&user)
	return user
}

func (db *userConnection) VerifyCredential(email string) any {
	var user entity.User
	res := db.connection.Where(("email = ?"), email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return db.connection.Where(("email = ?"), email).Take(&user)
}

func (db *userConnection) UpdateUser(user entity.User) entity.User {
	if user.Password != "" {
		user.Password = algorithm.HashAndSalt([]byte(user.Password))
	}
	db.connection.Updates(&user)
	return user
}

func (db *userConnection) GetAllEksportir() []entity.User {
	var users []entity.User
	db.connection.Where(("role = ?"), "eksportir").Find(&users)
	return users
}

func (db *userConnection) FindUserByID(userID string) entity.User {
	var user entity.User
	db.connection.Where(("id = ?"), userID).Find(&user)
	return user
}
