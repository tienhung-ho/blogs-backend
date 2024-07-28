package permissionhandler

import (
	permissionbusiness "blogs/internal/business/permission"
	"blogs/internal/common"
	permissionmodel "blogs/internal/model/permission"
	permissionstorage "blogs/internal/repository/mysql/permission"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePermission(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		var data permissionmodel.PermissionCreation

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInternal(err))
			return
		}

		store := permissionstorage.NewMysqlStorage(db)
		biz := permissionbusiness.NewPermissionCreationBiz(store)

		recordId, err := biz.CreatPermission(c.Request.Context(), &data)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewDataResponse(recordId))
	}
}
