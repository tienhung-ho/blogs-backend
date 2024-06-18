package authhandler

import (
	authbiz "blogs/internal/business/auth"
	"blogs/internal/common"
	jwtcus "blogs/internal/helpers/token/jwt"
	authmodel "blogs/internal/model/auth"
	usersmodel "blogs/internal/model/users"
	userstorage "blogs/internal/repository/mysql/user"
	"net/http"
	"os"

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

		secretKey := os.Getenv("SECRET_KEY")

		jwtService := jwtcus.NewJwtServices(secretKey, "user")

		store := userstorage.NewSqlStorage(db)
		biz := authbiz.NewLoginUserBiz(store, jwtService)

		simpleUser, err := biz.Login(c.Request.Context(), &data)

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrEmailOrPasswordInvalid(usersmodel.EntityName, err))
			return
		}

		tokeBiz := authbiz.NewTokenBiz(store, jwtService)
		userTokens, err := tokeBiz.GenerateToken(c.Request.Context(), simpleUser.ID)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.SetCookie("access_token", userTokens.AccessToken, 3600, "/", "localhost", false, true)
		c.SetCookie("refresh_token", userTokens.RefreshToken, 3600*24*250, "/", "localhost", false, true)

		c.JSON(http.StatusOK, common.NewReponseUserToken(userTokens.AccessToken, userTokens.RefreshToken))

	}
}
