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

func DeleteAccount(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInternal(err))
			return
		}

		store := accountsstogare.NewMysqlStorage(db)
		biz := accountsbusiness.AccountDeletionStorage(store)

		if err := biz.DeleteAccount(c.Request.Context(), map[string]interface{}{"id": id}); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInternal(err))
			return
		}

		c.JSON(http.StatusOK, common.NewDataResponse(true))
	}
}
