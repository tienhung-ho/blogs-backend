package authhandler

import (
	authbiz "blogs/internal/business/auth"
	"blogs/internal/common"
	authmodel "blogs/internal/model/auth"
	userstorage "blogs/internal/repository/mysql/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data authmodel.UserRegister
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return

		}

		store := userstorage.NewSqlStorage(db)
		biz := authbiz.NewCreateUserBiz(store)
		dataId, err := biz.CreateUser(c.Request.Context(), &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccesResponse(dataId))

	}
}
