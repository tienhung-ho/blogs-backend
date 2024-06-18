package helpers

import (
	"log"
	"os"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type hashBcrypt struct {
	password string
}

type compareResult struct {
	match bool
	err   error
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

// ComparePass compares the hashed password with the provided password using bcrypt.
func (h hashBcrypt) ComparePass(hashedPassword string) bool {
	resultChan := make(chan compareResult)

	// Start a goroutine to compare passwords
	go func() {
		start := time.Now()
		err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(h.password))
		duration := time.Since(start)
		log.Printf("Compare pass took %v", duration)

		if duration > time.Second {
			log.Printf("Warning: bcrypt.CompareHashAndPassword took longer than expected: %v", duration)
		}

		resultChan <- compareResult{
			match: err == nil,
			err:   err,
		}
	}()

	// Wait for the result
	result := <-resultChan
	return result.match
}
