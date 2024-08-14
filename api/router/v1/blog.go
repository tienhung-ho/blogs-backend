package router

import (
	bloghandler "blogs/api/handler/blog"
	"github.com/redis/go-redis/v9"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BlogRouter(blog *gin.RouterGroup, db *gorm.DB, rdb *redis.Client) {
	blog.GET("/:id", bloghandler.GetBlog(db, rdb))
	blog.GET("/", bloghandler.ListBlog(db))
	blog.POST("/", bloghandler.CreateBlog(db))
	blog.PATCH("/:id", bloghandler.UpdateBlog(db, rdb))
	blog.DELETE("/:id", bloghandler.DeleteBlog(db, rdb))
}
