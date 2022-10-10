package service

import (
	"dataekspor-be/dto"
	"dataekspor-be/entity"
	"dataekspor-be/repository"

	"github.com/mashingan/smapping"
)

type UserService interface {
	GetAllEksportir() []entity.User
	GetUserByID(userID string) entity.User
	GetUserProfile(userID string) entity.User
	UpdateUser(userDTO dto.UserUpdateDTO) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) GetAllEksportir() []entity.User {
	return s.userRepository.GetAllEksportir()
}

func (s *userService) GetUserByID(userID string) entity.User {
	return s.userRepository.FindUserByID(userID)
}

func (s *userService) GetUserProfile(userID string) entity.User {
	return s.userRepository.FindUserByID(userID)
}

func (s *userService) UpdateUser(userDTO dto.UserUpdateDTO) entity.User {
	user := entity.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&userDTO))
	if err != nil {
		panic(err)
	}
	res := s.userRepository.UpdateUser(user)
	return res
}
