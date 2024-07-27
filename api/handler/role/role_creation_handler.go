package rolehandler

import (
	rolebusiness "blogs/internal/business/role"
	"blogs/internal/common"
	rolemodel "blogs/internal/model/role"
	permissionstorage "blogs/internal/repository/mysql/permission"
	rolestorage "blogs/internal/repository/mysql/role"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRole(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		var data rolemodel.RoleCreation

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		roleStore := rolestorage.NewMysqlStorage(db)
		permissionStore := permissionstorage.NewMysqlStorage(db)
		biz := rolebusiness.NewRoleCreationBiz(roleStore, permissionStore)

		dataId, err := biz.CreateRole(c.Request.Context(), data)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewDataResponse(dataId))
	}
}
