package authmiddleware

import (
	"blogs/internal/common"
	jwtcus "blogs/internal/helpers/token/jwt"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// type authMiddlewares struct {
// 	tokenType string
// }

// func NewAuthMiddlewares(tken string) *authMiddlewares {
// 	return &authMiddlewares{
// 		tokenType: tken,
// 	}
// }

const (
	AccessToken  = "access_token"
	RefreshToken = "refresh_token"
)

func AuthMiddleware(tokenType string) gin.HandlerFunc {

	secretKey := os.Getenv("SECRET_KEY")
	jwtService := jwtcus.NewJwtServices(secretKey, "user")

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

		// Token is valid, store the claims in the context
		c.Set("userID", claims.UserId)
		c.Next()
	}
}
