package router

import (
	accountshandler "blogs/api/handler/accounts"
	accountauthhandler "blogs/api/handler/auth/account"
	accountsmiddlewares "blogs/api/middlewares/accounts"
	accountauthmiddlewares "blogs/api/middlewares/auth/account"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AccountRouter(accounts *gin.RouterGroup, db *gorm.DB) {
	accounts.GET("/:id", accountauthmiddlewares.AuthMiddleware(accountauthmiddlewares.AccessToken), accountshandler.FindAccount(db))
	accounts.POST("/login", accountauthhandler.Login(db))
	accounts.POST("/", accountauthmiddlewares.AuthMiddleware(accountauthmiddlewares.AccessToken), accountsmiddlewares.FileUploadMiddileware(), accountshandler.CreateAccount(db))
	accounts.PATCH("/:id", accountauthmiddlewares.AuthMiddleware(accountauthmiddlewares.AccessToken), accountsmiddlewares.FileUploadMiddileware(), accountshandler.UpdateAccount(db))
	accounts.DELETE("/:id", accountauthmiddlewares.AuthMiddleware(accountauthmiddlewares.AccessToken),accountshandler.DeleteAccount(db))
}
