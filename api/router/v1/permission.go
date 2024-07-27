package router

import (
	permissionhandler "blogs/api/handler/permission"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PermissionRouter(permissions *gin.RouterGroup, db *gorm.DB) {

	// Đăng ký các định tuyến và xử lý
	permissions.GET("/:id", permissionhandler.FindPermissions(db))
}