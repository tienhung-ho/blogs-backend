package authhandler

import (
	authbiz "blogs/internal/bussiness/auth"
	"blogs/internal/common"
	jwtcus "blogs/internal/helpers/token/jwt"
	authmodel "blogs/internal/model/auth"
	userstorage "blogs/internal/repository/mysql/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data authmodel.UserLogin

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		jwtService := jwtcus.NewJwtServices("conca", "user")

		store := userstorage.NewSqlStorage(db)
		biz := authbiz.NewLoginUserBiz(store, jwtService)

		tokens, err := biz.Login(c.Request.Context(), &data)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewReponseUserToken(tokens.AccessToken, tokens.RefreshToken))

	}
}
