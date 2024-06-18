package usershandler

import (
	userbiz "blogs/internal/business/user"
	"blogs/internal/common"
	usersmodel "blogs/internal/model/users"
	userstorage "blogs/internal/repository/mysql/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var dataUpdate usersmodel.UserEdition

		userId, ok := c.Get("userID")

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "userId not found in context"})
			return
		}

		userIdInt, ok := userId.(int)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "userId is not of type int"})
			return
		}

		if err := c.ShouldBind(&dataUpdate); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSqlStorage(db)
		biz := userbiz.NewEditUserBiz(store)

		if err := biz.UpdateUser(c.Request.Context(), userIdInt, &dataUpdate); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccesResponse(true))
	}
}
