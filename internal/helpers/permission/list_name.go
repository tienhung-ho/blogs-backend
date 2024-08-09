package permissionhelper

import (
	permissionmodel "blogs/internal/model/permission"
)

func ListName(permissions []permissionmodel.Permission) []string {
	var names []string
	for _, perm := range permissions {
		names = append(names, perm.Name)
	}
	return names
}
