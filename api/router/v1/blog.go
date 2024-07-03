package router

import (
	bloghandler "blogs/api/handler/blog"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BlogRouter(blog *gin.RouterGroup, db *gorm.DB) {
	blog.GET("/:id", bloghandler.GetBlog(db))
}
