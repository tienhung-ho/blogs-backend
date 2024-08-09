package policiescasbin

import (
	permissionmodel "blogs/internal/model/permission"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

var (
	instance *casbin.Enforcer
	once     sync.Once
)

// InitEnforcer initializes the Casbin enforcer as a singleton.
func InitEnforcer(db *gorm.DB, modelPath string) (*casbin.Enforcer, error) {
	var err error
	once.Do(func() {
		var adapter *gormadapter.Adapter
		adapter, err = gormadapter.NewAdapterByDB(db)
		if err != nil {
			return
		}
		instance, err = casbin.NewEnforcer(modelPath, adapter)
		if err != nil {
			return
		}
		err = instance.LoadPolicy()
	})
	return instance, err
}

// GetEnforcer returns the singleton instance of the Casbin enforcer.
func GetEnforcer() *casbin.Enforcer {
	return instance
}

func SyncPermissions(permissions []permissionmodel.Permission, version string) {
	// Xóa tất cả các chính sách hiện tại
	enforcer := GetEnforcer()
	enforcer.RemoveFilteredPolicy(0, "*", "*", "*")

	for _, permission := range permissions {
		// Tách tên quyền thành các phần: Hành động và Đối tượng
		// log.Printf("permission: %v", permission)
		parts := strings.Split(permission.Name, "_")
		if len(parts) != 2 {
			// Nếu tên quyền không theo định dạng "Action_Object", bỏ qua
			continue
		}

		action := parts[0]
		object := parts[1]

		var method string
		switch action {
		case "View":
			method = "GET"
		case "Create":
			method = "POST"
		case "Update":
			method = "PATCH"
		case "Delete":
			method = "DELETE"
		default:
			// Nếu hành động không được hỗ trợ, bỏ qua
			continue
		}

		// Tạo endpoint tương ứng từ đối tượng
		endpoint := fmt.Sprintf("/%s/%s/*", version, strings.ToLower(object))

		// Thêm chính sách vào enforcer
		if ok, err := enforcer.AddPolicy(permission.Name, endpoint, method); !ok || err != nil {
			log.Printf("Failed to add policy: %v", err)
		}
	}
	if err := enforcer.SavePolicy(); err != nil {
		log.Printf("Failed to save policies: %v", err)
	}
}
