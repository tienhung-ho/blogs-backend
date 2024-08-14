package blogcachestorage

import (
	"blogs/internal/common"
	blogmodel "blogs/internal/model/blog"
	"context"
	"errors"
	"fmt"
)

func (rdb *redisStorage) DeleteBlog(ctx context.Context, cond map[string]interface{}) error {
	blogID, ok := cond["id"].(int)
	if !ok {
		return errors.New("invalid blog ID")
	}
	key := blogmodel.EntityName + fmt.Sprintf("%d", blogID)
	if err := rdb.rdb.Del(ctx, key).Err(); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
