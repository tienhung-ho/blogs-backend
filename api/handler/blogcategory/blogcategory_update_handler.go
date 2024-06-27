package blogcategoryhandler

import (
	blogcategorybusiness "blogs/internal/business/blogcategory"
	"blogs/internal/common"
	blogcategorymodel "blogs/internal/model/blogcategory"
	blogcategorystorage "blogs/internal/repository/mysql/blogcategory"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateBlogCategory(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data blogcategorymodel.UpdateBlogCategory
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := blogcategorystorage.NewSqlStorage(db)
		biz := blogcategorybusiness.NewUpdateBlogCategoryBiz(store)

		if err := biz.UpdateBlogCategory(c.Request.Context(), id, data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccesResponse(true))

	}

}
