package accountauthmiddlewares

import (
	rolebusiness "blogs/internal/business/role"
	"blogs/internal/common"
	permissionhelper "blogs/internal/helpers/permission"
	jwtcus "blogs/internal/helpers/token/jwt"
	rolestorage "blogs/internal/repository/mysql/role"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	AccessToken  = "account_access_token"
	RefreshToken = "account_refresh_token"
)

func AuthMiddleware(tokenType string, db *gorm.DB) gin.HandlerFunc {
	secretKey := os.Getenv("SECRET_KEY_ACCOUNT")
	jwtService := jwtcus.NewJwtServices(secretKey, "account")
	return func(c *gin.Context) {
		tokenString, err := c.Cookie(tokenType)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "auth token required"})
			c.Abort()
			return
		}

		claims, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, common.NewReponseErrToken(fmt.Sprintf("%v", err), "accesstoken"))
			c.Abort()
			return
		}

		store := rolestorage.NewMysqlStorage(db)
		biz := rolebusiness.NewRoleFiditionBiz(store)

		role, err := biz.FindRole(c.Request.Context(), map[string]interface{}{"name": claims.Role})

		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			c.Abort()
			return
		}

		permissionName := permissionhelper.ListName(role.Permissions)

		c.Set("accountID", claims.Id)
		c.Set("role", claims.Role)
		c.Set("permissions", permissionName)
		c.Next()
	}
}
