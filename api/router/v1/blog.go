package router

import (
	bloghandler "blogs/api/handler/blog"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BlogRouter(blog *gin.RouterGroup, db *gorm.DB) {
	blog.GET("/:id", bloghandler.GetBlog(db))
	blog.GET("/", bloghandler.ListBlog(db))
	blog.POST("/", bloghandler.CreateBlog(db))
	blog.PATCH("/:id", bloghandler.UpdateBlog(db))
	blog.DELETE("/:id", bloghandler.DeleteBlog(db))
}
