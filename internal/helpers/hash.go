package helpers

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type hashBcrypt struct {
	password string
}

func NewHashBcrypt(password string) *hashBcrypt {
	return &hashBcrypt{
		password: password,
	}
}

func (h hashBcrypt) GeneratePass() (string, error) {
	cost := os.Getenv("COST")
	costInt, err := strconv.Atoi(cost)

	if err != nil {
		return "", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(h.password), costInt)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (h hashBcrypt) ComparePass(hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(h.password))
	return err == nil
}
