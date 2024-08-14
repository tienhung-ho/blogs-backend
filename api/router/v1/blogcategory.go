package router

import (
	blogcategoryhandler "blogs/api/handler/blogcategory"
	accountauthmiddlewares "blogs/api/middlewares/auth/account"
	policiesmiddleware "blogs/api/middlewares/policies"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BlogCategoryRouter(blogcategory *gin.RouterGroup, db *gorm.DB) {

	blogcategory.Use(accountauthmiddlewares.AuthMiddleware(accountauthmiddlewares.AccessToken, db), policiesmiddleware.Middleware())

	blogcategory.GET("/:id", blogcategoryhandler.FindBlogCategory(db))
	blogcategory.POST("/", blogcategoryhandler.CreateBlogCategory(db))
	blogcategory.PATCH("/:id", blogcategoryhandler.UpdateBlogCategory(db))
	blogcategory.GET("/list", blogcategoryhandler.ListBlogCategory(db))
	blogcategory.DELETE("/:id", blogcategoryhandler.DeletionBlogCategory(db))
}
