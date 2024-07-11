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

func FindRole(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		store := rolestorage.NewMysqlStorage(db)
		biz := rolebusiness.NewRoleFiditionBiz(store)

		cond := map[string]interface{}{
			"deleted": false,
			"id":      id,
		}

		role, err := biz.FindRole(c.Request.Context(), cond)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccesResponse(role))
	}
}
