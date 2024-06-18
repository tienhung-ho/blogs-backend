package authhandler

import (
	authbiz "blogs/internal/business/auth"
	"blogs/internal/common"
	jwtcus "blogs/internal/helpers/token/jwt"
	userstorage "blogs/internal/repository/mysql/user"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RefreshToken(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
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
		secretKey := os.Getenv("SECRET_KEY")

		jwtService := jwtcus.NewJwtServices(secretKey, "user")
		store := userstorage.NewSqlStorage(db)
		biz := authbiz.NewTokenBiz(store, jwtService)

		userTokens, err := biz.GenerateToken(c.Request.Context(), userIdInt)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.SetCookie("access_token", userTokens.AccessToken, 3600, "/", "localhost", false, true)
		c.SetCookie("refresh_token", userTokens.RefreshToken, 3600*24*250, "/", "localhost", false, true)

		c.JSON(http.StatusOK, common.NewReponseUserToken(userTokens.AccessToken, userTokens.RefreshToken))

	}
}
