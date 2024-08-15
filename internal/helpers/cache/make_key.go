package cachehelper

import (
	"blogs/internal/common"
	filtermodel "blogs/internal/model/filter"
	"fmt"
)

// GenerateCacheKey Chuyển đổi điều kiện thành chuỗi để làm key
func GenerateCacheKey(entityName string, cond map[string]interface{}, paging common.Paging, filter filtermodel.Filter) string {
	key := fmt.Sprintf("%s:", entityName)
	for k, v := range cond {
		key += fmt.Sprintf("%s=%v:", k, v)
	}

	key += fmt.Sprintf("%s=%v:", "page", paging.Page)
	key += fmt.Sprintf("%s=%v:", "limit", paging.Limit)

	key += fmt.Sprintf("%s=%v:", "filter", filter.Status)
	
	return key
}
