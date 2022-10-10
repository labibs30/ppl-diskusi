package id

import "github.com/google/uuid"

func GenerateUUID() string {
	return uuid.NewString()
}

func ValidateUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}
