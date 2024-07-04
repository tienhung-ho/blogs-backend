package bloghandler

import (
	blogbusiness "blogs/internal/business/blog"
	"blogs/internal/common"
	blogmodel "blogs/internal/model/blog"
	blogstorage "blogs/internal/repository/mysql/blog"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateBlog(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data blogmodel.BlogCreation

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := blogstorage.NewMysqlStorage(db)
		biz := blogbusiness.NewBlogCreationBusiness(store)

		dataId, err := biz.CreateBlog(c.Request.Context(), &data)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewDataResponse(dataId))

	}
}
