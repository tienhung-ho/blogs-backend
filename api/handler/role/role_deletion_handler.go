package rolehandler

import (
	rolebusiness "blogs/internal/business/role"
	"blogs/internal/common"
	rolestorage "blogs/internal/repository/mysql/role"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteRole(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := rolestorage.NewMysqlStorage(db)
		biz := rolebusiness.NewRoleDeletionBusiness(store)

		if err := biz.DeleteRole(c.Request.Context(), map[string]interface{}{"id": id}); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewDataResponse(true))
	}
}
