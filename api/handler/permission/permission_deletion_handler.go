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

func DeletePermission(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		store := permissionstorage.NewMysqlStorage(db)
		biz := permissionbusiness.NewPermissionDeletionBiz(store)

		if err := biz.DeletePermission(c.Request.Context(), map[string]interface{}{
			"id":      id,
			"deleted": false,
		}); err != nil {

			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewDataResponse(true))

	}
}
