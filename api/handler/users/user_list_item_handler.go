package usershandler

import (
	userbiz "blogs/internal/business/user"
	"blogs/internal/common"
	filtermodel "blogs/internal/model/filter"
	userstorage "blogs/internal/repository/mysql/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListUser(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		paging.Process()

		var filter filtermodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := userstorage.NewSqlStorage(db)
		biz := userbiz.NewUserListItemBiz(store)

		records, err := biz.ListItem(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewSuccesResponse(records, paging, filter))
	}
}
