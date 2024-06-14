package helpers

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePass(pass interface{}) (string, error) {
	var password string
	switch v := pass.(type) {
	case string:
		password = v
	case []byte:
		password = string(v)
	default:
		return "", fmt.Errorf("unsupported data type: %T", v)
	}
	cost := os.Getenv("COST")
	costInt, err := strconv.Atoi(cost)

	if err != nil {
		return "", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), costInt)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func ComparePass(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
