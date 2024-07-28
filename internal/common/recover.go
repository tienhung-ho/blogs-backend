package common

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RecoverTransaction is a utility function to handle recover in transactions
func RecoverTransaction(db *gorm.DB) {
	if r := recover(); r != nil {
		fmt.Printf("Recovered from panic: %v\n", r)
		db.Rollback()
	}
}

func RecoverMiddleware(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Recovered from panic: %v\n", err)
			debug.PrintStack()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			c.Abort()
		}
	}()
	c.Next()
}
