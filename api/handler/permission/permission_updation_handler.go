package permissionhandler

import (
	permissionbusiness "blogs/internal/business/permission"
	"blogs/internal/common"
	permissionmodel "blogs/internal/model/permission"
	permissionstorage "blogs/internal/repository/mysql/permission"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdatePermission(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInternal(err))
			return
		}

		var data permissionmodel.PermissionUpdation

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInternal(err))
			return
		}

		store := permissionstorage.NewMysqlStorage(db)
		biz := permissionbusiness.NewPermissionUpdationBiz(store)

		if err := biz.UpdatePermission(c.Request.Context(), map[string]interface{}{"id": id}, &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewDataResponse(true))
	}
}
