package blogcategoryhandler

import (
	blogcategorybusiness "blogs/internal/business/blogcategory"
	"blogs/internal/common"
	blogcategorystorage "blogs/internal/repository/mysql/blogcategory"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindBlogCategory(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := blogcategorystorage.NewSqlStorage(db)
		biz := blogcategorybusiness.NewFinditionBlogCategory(store)

		record, err := biz.GetBlogCategoryByCondition(c.Request.Context(), map[string]interface{}{"id": id})

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccesResponse(record))

	}
}
