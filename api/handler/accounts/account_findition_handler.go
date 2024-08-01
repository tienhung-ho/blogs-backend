package accountshandler

import (
	accountsbusiness "blogs/internal/business/accounts"
	"blogs/internal/common"
	accountsstogare "blogs/internal/repository/mysql/accounts"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindAccount(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInternal(err))
			return
		}

		store := accountsstogare.NewMysqlStorage(db)
		biz := accountsbusiness.NewAccountFinditionBiz(store)

		record, err := biz.FindAccount(c.Request.Context(), map[string]interface{}{
			"id":      id,
			"deleted": false,
		})

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewDataResponse(record))

	}

}
