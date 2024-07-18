package router

import (
	rolehandler "blogs/api/handler/role"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoleRouter(role *gin.RouterGroup, db *gorm.DB) {
	role.GET("/:id", rolehandler.FindRole(db))
	role.GET("/", rolehandler.RoleListItem(db))
	role.POST("/", rolehandler.CreateRole(db))
	role.DELETE("/:id", rolehandler.DeleteRole(db))
}
