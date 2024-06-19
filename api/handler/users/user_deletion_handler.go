package usershandler

import (
	userbiz "blogs/internal/business/user"
	"blogs/internal/common"
	userstorage "blogs/internal/repository/mysql/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteUser(db *gorm.DB) func(*gin.Context) {
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

		store := userstorage.NewSqlStorage(db)
		biz := userbiz.NewDeleteUserBiz(store)

		if err := biz.DeleteUser(c.Request.Context(), userIdInt); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusBadRequest, common.SimpleSuccesResponse(true))

	}
}
