package router

import (
	accountshandler "blogs/api/handler/accounts"
	accountauthhandler "blogs/api/handler/auth/account"
	accountsmiddlewares "blogs/api/middlewares/accounts"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AccountRouter(accounts *gin.RouterGroup, db *gorm.DB) {
	accounts.GET("/:id", accountshandler.FindAccount(db))
	accounts.POST("/login", accountauthhandler.Login(db))
	accounts.POST("/", accountsmiddlewares.FileUploadMiddileware(), accountshandler.CreateAccount(db))
	accounts.PATCH("/:id", accountsmiddlewares.FileUploadMiddileware(), accountshandler.UpdateAccount(db))
	accounts.DELETE("/:id", accountshandler.DeleteAccount(db))
}
