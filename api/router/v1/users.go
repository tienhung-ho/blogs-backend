package router

import (
	authhandler "blogs/api/handler/auth"
	usershandler "blogs/api/handler/users"
	authmiddleware "blogs/api/middlewares/auth"
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
	users.GET("/:id", authmiddleware.AuthMiddleware(authmiddleware.AccessToken), usershandler.FindUser(db))
	users.POST("/", authhandler.CreateUser(db))
	users.POST("/login", authhandler.Login(db))
	users.POST("/refreshtoken", authmiddleware.AuthMiddleware(authmiddleware.RefreshToken), authhandler.RefreshToken(db))

}
