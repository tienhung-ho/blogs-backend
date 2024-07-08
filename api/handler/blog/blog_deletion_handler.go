package bloghandler

import (
	blogbusiness "blogs/internal/business/blog"
	"blogs/internal/common"
	blogstorage "blogs/internal/repository/mysql/blog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteBlog(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		store := blogstorage.NewMysqlStorage(db)
		biz := blogbusiness.NewBlogDeletionBiz(store)

		if err := biz.DeleteBlog(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewDataResponse(true))
	}
}
