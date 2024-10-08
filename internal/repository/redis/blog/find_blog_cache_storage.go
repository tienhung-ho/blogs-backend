package blogcachestorage

import (
	"blogs/internal/common"
	cachehelper "blogs/internal/helpers/cache"
	blogmodel "blogs/internal/model/blog"
	filtermodel "blogs/internal/model/filter"
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
)

func (rdb *redisStorage) GetBlog(ctx context.Context, cond map[string]interface{}, morekeys ...string) (*blogmodel.Blog, error) {
	var paging common.Paging
	paging.Process()

	key := cachehelper.GenerateCacheKey(blogmodel.EntityName, cond, paging, filtermodel.Filter{})
	result, err := rdb.rdb.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil // Cache miss
	} else if err != nil {
		return nil, err
	}

	var blog blogmodel.Blog

	if err := json.Unmarshal([]byte(result), &blog); err != nil {
		return nil, err
	}

	return &blog, nil
}
