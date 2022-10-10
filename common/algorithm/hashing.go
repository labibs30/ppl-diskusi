package algorithm

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		fmt.Println("hi")
		panic(err)
	}
	return string(hash)
}
