package blogcachestorage

import (
	"blogs/internal/common"
	blogmodel "blogs/internal/model/blog"
	"context"
	"encoding/json"
	"errors"
	"time"
)

// CreateBlog lưu trữ blog vào Redis
func (rdb *redisStorage) CreateBlog(ctx context.Context, data interface{}, morekeys ...string) (int, error) {
	if len(morekeys) == 0 {
		return 0, errors.New("missing cache key")
	}

	key := morekeys[0]

	var record []byte
	var err error

	switch v := data.(type) {
	case *blogmodel.BlogCreation:
		record, err = json.Marshal(v)
		if err != nil {
			return 0, common.ErrDB(err)
		}
	case []blogmodel.Blog:
		record, err = json.Marshal(v)
		if err != nil {
			return 0, common.ErrDB(err)
		}
	default:
		return 0, errors.New("unsupported data type")
	}

	if err := rdb.rdb.Set(ctx, key, string(record), 20*time.Second).Err(); err != nil {
		return 0, common.ErrDB(err)
	}

	return 0, nil
}
