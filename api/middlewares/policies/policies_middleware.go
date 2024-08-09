package policiesmiddleware

import (
	"net/http"
	"strings"

	policiescasbin "blogs/internal/policies"

	"github.com/gin-gonic/gin"
)

// Middleware returns a Casbin middleware function.
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		enforcer := policiescasbin.GetEnforcer()

		if enforcer == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Casbin enforcer not initialized"})
			c.Abort()
			return
		}

		permissionsI, exists := c.Get("permissions")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		permissions, ok := permissionsI.([]string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid permissions format"})
			c.Abort()
			return
		}

		resource := c.Request.URL.Path
		action := c.Request.Method

		// Map HTTP method to permission prefix
		var requiredPermissionPrefix string
		switch action {
		case "GET":
			requiredPermissionPrefix = "View"
		case "POST":
			requiredPermissionPrefix = "Create"
		case "PATCH":
			requiredPermissionPrefix = "Update"
		case "DELETE":
			requiredPermissionPrefix = "Delete"
		default:
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
			c.Abort()
			return
		}

		for _, permission := range permissions {
			// log.Printf("Checking permission: %v for resource: %v, action: %v, requiredPermission: %v", permission, resource, action, requiredPermissionPrefix)
			if strings.HasPrefix(permission, requiredPermissionPrefix) {
				ok, err := enforcer.Enforce(permission, resource, action)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					c.Abort()
					return
				}

				if ok {
					// Nếu tìm thấy quyền hợp lệ, tiếp tục xử lý yêu cầu
					c.Next()
					return
				}
			}
		}

		// Nếu không có quyền hợp lệ, trả về lỗi
		c.JSON(http.StatusForbidden, gin.H{"message": "Access denied"})
		c.Abort()
	}
}
