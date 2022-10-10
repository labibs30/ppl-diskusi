package service

import (
	"dataekspor-be/common/algorithm"
	"dataekspor-be/common/id"
	dto "dataekspor-be/dto"
	"dataekspor-be/entity"
	"dataekspor-be/repository"

	"github.com/mashingan/smapping"
)

type AuthService interface {
	CreateUser(user dto.UserRegisterDTO) entity.User
	IsDuplicateEmail(email string) bool
	VerifyCredential(email string, password string) any
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepository,
	}
}

func (service *authService) VerifyCredential(email string, password string) any {
	res := service.userRepository.VerifyCredential(email)
	if v, ok := res.(entity.User); ok {
		comparedPassword := algorithm.ComparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return true
}

func (service *authService) CreateUser(user dto.UserRegisterDTO) entity.User {
	CreatedUser := entity.User{}
	err := smapping.FillStruct(&CreatedUser, smapping.MapFields(&user))
	if err != nil {
		panic(err)
	}
	CreatedUser.ID = id.GenerateUUID()
	CreatedUser.Password = algorithm.HashAndSalt([]byte(CreatedUser.Password))
	res := service.userRepository.InsertUser(CreatedUser)
	return res
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}
