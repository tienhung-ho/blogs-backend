package router

import (
	"blogs/internal/common"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Đăng ký các định tuyến và xử lý
	// Tự thêm middleware nếu cần
	router.Use(common.RecoverMiddleware)
	v1 := router.Group("/v1")
	{
		users := v1.Group("/users")
		{
			UsersRouter(users, db)
		}
		blogcategory := v1.Group("/blog-category")
		{
			BlogCategoryRouter(blogcategory, db)
		}

		blogs := v1.Group("/blogs")
		{
			BlogRouter(blogs, db)
		}

		roles := v1.Group("/roles")
		{
			RoleRouter(roles, db)
		}

		permissions := v1.Group("/permissions")
		{
			PermissionRouter(permissions, db)
		}
	}

	return router
}
