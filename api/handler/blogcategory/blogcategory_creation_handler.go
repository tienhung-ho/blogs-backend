package blogcategoryhandler

import (
	blogcategorybusiness "blogs/internal/business/blogcategory"
	"blogs/internal/common"
	blogcategorymodel "blogs/internal/model/blogcategory"
	blogcategorystorage "blogs/internal/repository/mysql/blogcategory"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateBlogCategory(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data blogcategorymodel.CreationBlogCategory

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := blogcategorystorage.NewSqlStorage(db)
		biz := blogcategorybusiness.NewCreationBlogCategoryBiz(store)

		recordId, err := biz.CreateBlogCategory(c.Request.Context(), &data)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccesResponse(recordId))

	}
}
