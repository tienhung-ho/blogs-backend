package rolehandler

import (
	rolebusiness "blogs/internal/business/role"
	"blogs/internal/common"
	rolemodel "blogs/internal/model/role"
	rolestorage "blogs/internal/repository/mysql/role"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateRole(db *gorm.DB) func(c *gin.Context) {

	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		var data rolemodel.RoleUpdation

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := rolestorage.NewMysqlStorage(db)
		biz := rolebusiness.NewRoleUpdationBiz(store)

		if err := biz.UpdateRole(c.Request.Context(), map[string]interface{}{"id": id}, data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewDataResponse(true))
	}
}
