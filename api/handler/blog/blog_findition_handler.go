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

func GetBlog(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		store := blogstorage.NewMysqlStorage(db)
		biz := blogbusiness.NewBlogBiz(store)

		data, err := biz.GetBlog(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewDataResponse(data))
	}
}
