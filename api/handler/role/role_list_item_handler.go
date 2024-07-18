package rolehandler

import (
	rolebusiness "blogs/internal/business/role"
	"blogs/internal/common"
	rolestorage "blogs/internal/repository/mysql/role"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoleListItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		store := rolestorage.NewMysqlStorage(db)
		biz := rolebusiness.NewRoleListItemBiz(store)

		roles, err := biz.ListItem(c.Request.Context(), map[string]interface{}{"deleted": false})

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewDataResponse(roles))
	}

}
