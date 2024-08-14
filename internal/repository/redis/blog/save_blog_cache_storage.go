package blogcachestorage

import (
	"blogs/internal/common"
	blogmodel "blogs/internal/model/blog"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

func (rdb *redisStorage) CreateBlog(ctx context.Context, blog *blogmodel.BlogCreation) (int, error) {

	key := blogmodel.EntityName + fmt.Sprintf("%d", blog.Id)
	record, err := json.Marshal(blog)

	if err != nil {
		return 0, common.ErrDB(err)
	}

	if err := rdb.rdb.Set(ctx, key, string(record), 10*time.Minute).Err(); err != nil {
		return 0, common.ErrDB(err)
	}

	return 0, nil
}
