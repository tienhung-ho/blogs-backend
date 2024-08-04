package accountshandler

import (
	accountsbusiness "blogs/internal/business/accounts"
	"blogs/internal/common"
	accountmodel "blogs/internal/model/accounts"
	accountsstogare "blogs/internal/repository/mysql/accounts"
	rolestorage "blogs/internal/repository/mysql/role"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateAccount(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		fileURLInterface, ok := c.Get("fileURL")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "fileURL not found in context"})
			return
		}

		fileURL, ok := fileURLInterface.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert fileURL to string"})
			return
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInternal(err))
			return
		}

		var data accountmodel.AccountUpdation

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInternal(err))
			return
		}

		data.Image = fileURL

		accountStore := accountsstogare.NewMysqlStorage(db)
		roleStore := rolestorage.NewMysqlStorage(db)

		biz := accountsbusiness.NewAccountUpdationBiz(accountStore, roleStore)

		if err := biz.UpdateAccount(c.Request.Context(), map[string]interface{}{"id": id}, &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewDataResponse(true))

	}
}
