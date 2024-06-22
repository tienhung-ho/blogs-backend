package blogcategoryhandler

import (
	blogcategorybusiness "blogs/internal/business/blogcategory"
	"blogs/internal/common"
	blogcategorystorage "blogs/internal/repository/mysql/blogcategory"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListBlogCategory(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		store := blogcategorystorage.NewSqlStorage(db)
		biz := blogcategorybusiness.NewListItemBlogCategoryStorage(store)

		categoryTree, err := biz.ListItem(c.Request.Context(), map[string]interface{}{"status": "Active"})

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccesResponse(categoryTree))

	}
}
