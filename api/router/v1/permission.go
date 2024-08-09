package router

import (
	permissionhandler "blogs/api/handler/permission"
	accountauthmiddlewares "blogs/api/middlewares/auth/account"
	policiesmiddleware "blogs/api/middlewares/policies"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PermissionRouter(permissions *gin.RouterGroup, db *gorm.DB) {

	permissions.Use(accountauthmiddlewares.AuthMiddleware(accountauthmiddlewares.AccessToken, db), policiesmiddleware.Middleware())

	// Đăng ký các định tuyến và xử lý
	permissions.GET("/:id", permissionhandler.FindPermissions(db))
	permissions.GET("/list", permissionhandler.ListPermissions(db))
	permissions.POST("/", permissionhandler.CreatePermission(db))
	permissions.DELETE("/:id", permissionhandler.DeletePermission(db))
	permissions.PATCH("/:id", permissionhandler.UpdatePermission(db))
}
