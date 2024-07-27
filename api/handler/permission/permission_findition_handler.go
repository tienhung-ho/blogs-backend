package permissionhandler

import (
	permissionbusiness "blogs/internal/business/permission"
	"blogs/internal/common"
	permissionstorage "blogs/internal/repository/mysql/permission"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindPermissions(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		store := permissionstorage.NewMysqlStorage(db)
		biz := permissionbusiness.NewPermissionFinditionBiz(store)

		record, err := biz.FindPermissions(c.Request.Context(), map[string]interface{}{
			"id": id,
		})

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewDataResponse(record))

		return
	}
}
