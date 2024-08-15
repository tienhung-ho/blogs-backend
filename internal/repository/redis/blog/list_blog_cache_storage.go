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

func (rdb *redisStorage) ListItem(ctx context.Context, cond map[string]interface{}, paging *common.Paging, filter *filtermodel.Filter, morekyes ...string) ([]blogmodel.Blog, error) {
	key := cachehelper.GenerateCacheKey(blogmodel.EntityName, cond, *paging, *filter)
	result, err := rdb.rdb.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil // Cache miss
	} else if err != nil {
		return nil, err
	}

	var blogs []blogmodel.Blog

	if err := json.Unmarshal([]byte(result), &blogs); err != nil {
		return nil, err
	}

	return blogs, nil
}
