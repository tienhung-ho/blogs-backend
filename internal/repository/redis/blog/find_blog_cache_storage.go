package blogcachestorage

import (
	blogmodel "blogs/internal/model/blog"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func (rdb *redisStorage) GetBlog(ctx context.Context, cond map[string]interface{}) (*blogmodel.Blog, error) {
	blogID, ok := cond["id"].(int)
	if !ok {
		return nil, errors.New("invalid blog ID")
	}

	key := blogmodel.EntityName + fmt.Sprintf("%d", blogID)
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
