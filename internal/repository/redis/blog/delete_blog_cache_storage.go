package blogcachestorage

import (
	"blogs/internal/common"
	cachehelper "blogs/internal/helpers/cache"
	blogmodel "blogs/internal/model/blog"
	filtermodel "blogs/internal/model/filter"
	"context"
)

func (rdb *redisStorage) DeleteBlog(ctx context.Context, cond map[string]interface{}, morekeys ...string) error {

	var paging common.Paging
	paging.Process()

	key := cachehelper.GenerateCacheKey(blogmodel.EntityName, cond, paging, filtermodel.Filter{})
	if err := rdb.rdb.Del(ctx, key).Err(); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
