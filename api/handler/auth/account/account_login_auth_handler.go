package accountauthhandler

import (
	accountauthbiz "blogs/internal/business/auth/account"
	rolebusiness "blogs/internal/business/role"
	"blogs/internal/common"
	jwtcus "blogs/internal/helpers/token/jwt"
	accountautmodel "blogs/internal/model/auth/account"
	accountsstogare "blogs/internal/repository/mysql/accounts"
	rolestorage "blogs/internal/repository/mysql/role"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		var data accountautmodel.AccountLogin

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		secretKey := os.Getenv("SECRET_KEY_ACCOUNT")

		jwtService := jwtcus.NewJwtServices(secretKey, "account")

		store := accountsstogare.NewMysqlStorage(db)
		biz := accountauthbiz.NewLoginOfAccountBiz(store, jwtService)

		simpleAccount, err := biz.Login(c.Request.Context(), &data)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		roleStore := rolestorage.NewMysqlStorage(db)
		roleBiz := rolebusiness.NewRoleFiditionBiz(roleStore)

		role, err := roleBiz.FindRole(c.Request.Context(), map[string]interface{}{"id": simpleAccount.RoleID})

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		tokeBiz := accountauthbiz.NewTokenBiz(store, jwtService)
		accountTokens, err := tokeBiz.GenerateToken(c.Request.Context(), simpleAccount.ID, role.Name)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.SetCookie("account_access_token", accountTokens.AccessToken, 3600, "/", "localhost", false, true)
		c.SetCookie("account_refresh_token", accountTokens.RefreshToken, 3600*24*250, "/", "localhost", false, true)

		c.JSON(http.StatusOK, common.NewReponseUserToken(accountTokens.AccessToken, accountTokens.RefreshToken))

	}

}
