package bloghandler

import (
	blogbusiness "blogs/internal/business/blog"
	"blogs/internal/common"
	blogstorage "blogs/internal/repository/mysql/blog"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListBlog(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		condition := map[string]interface{}{
			"deleted": false,
			"status":  []string{"pending", "active"},
		}

		store := blogstorage.NewMysqlStorage(db)
		biz := blogbusiness.NewListItemBlogBiz(store)
		data, err := biz.ListItem(c.Request.Context(), condition)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewDataResponse(data))
	}
}
