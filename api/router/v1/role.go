package router

import (
	rolehandler "blogs/api/handler/role"
	accountauthmiddlewares "blogs/api/middlewares/auth/account"
	policiesmiddleware "blogs/api/middlewares/policies"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoleRouter(role *gin.RouterGroup, db *gorm.DB) {

	role.Use(accountauthmiddlewares.AuthMiddleware(accountauthmiddlewares.AccessToken, db), policiesmiddleware.Middleware())

	role.GET("/:id", rolehandler.FindRole(db))
	role.GET("/", rolehandler.RoleListItem(db))
	role.POST("/", rolehandler.CreateRole(db))
	role.PATCH("/:id", rolehandler.UpdateRole(db))
	role.DELETE("/:id", rolehandler.DeleteRole(db))
}
