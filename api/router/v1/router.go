package router

import (
	"blogs/internal/common"
	"github.com/redis/go-redis/v9"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB, rdb *redis.Client) *gin.Engine {

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
		blogcategory := v1.Group("/blog-categories")
		{
			BlogCategoryRouter(blogcategory, db)
		}

		blogs := v1.Group("/blogs")
		{
			BlogRouter(blogs, db, rdb)
		}

		roles := v1.Group("/roles")
		{
			RoleRouter(roles, db)
		}

		permissions := v1.Group("/permissions")
		{
			PermissionRouter(permissions, db)
		}

		accounts := v1.Group("/accounts")
		{
			AccountRouter(accounts, db)
		}
	}

	return router
}
