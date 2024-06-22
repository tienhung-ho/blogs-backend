package router

import (
	blogcategoryhandler "blogs/api/handler/blogcategory"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BlogCategoryRouter(blogcategory *gin.RouterGroup, db *gorm.DB) {

	// Đăng ký các định tuyến và xử lý
	blogcategory.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	blogcategory.GET("/:id", blogcategoryhandler.FindBlogCategory(db))
	blogcategory.POST("/", blogcategoryhandler.CreateBlogCategory(db))
	blogcategory.GET("/list", blogcategoryhandler.ListBlogCategory(db))

}
