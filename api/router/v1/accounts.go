package router

import (
	accountshandler "blogs/api/handler/accounts"
	accountsmiddlewares "blogs/api/middlewares/accounts"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AccountRouter(accounts *gin.RouterGroup, db *gorm.DB) {
	accounts.GET("/:id", accountshandler.FindAccount(db))
	accounts.POST("/", accountsmiddlewares.FileUploadMiddileware(), accountshandler.CreateAccount(db))

}
