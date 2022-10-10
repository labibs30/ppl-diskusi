package service

import (
	"dataekspor-be/entity"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	GenerateToken(userID string, Roles ...string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID string               `json:"user_id"`
	Roles  entity.SliceOfString `json:"role"`
	jwt.RegisteredClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "dataEkspor",
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "d4t4_3k5p0r"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(UserID string, Roles ...string) string {
	claims := &jwtCustomClaim{
		UserID,
		Roles,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
			Issuer:    j.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		fmt.Println("hi")
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(token, func(t_ *jwt.Token) (any, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
	return parsedToken, err
}
