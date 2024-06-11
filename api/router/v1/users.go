package router

import (
	usershandler "blogs/api/handler/users"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UsersRouter(users *gin.RouterGroup, db *gorm.DB) {
	// Đăng ký các định tuyến và xử lý
	users.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	users.GET("/:id", func(ctx *gin.Context) {})
	users.POST("/", usershandler.CreateUser(db))
	users.PUT("/:id", func(ctx *gin.Context) {})
	users.DELETE("/:id", func(ctx *gin.Context) {})

}
