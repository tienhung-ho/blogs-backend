package bloghandler

import (
	blogbusiness "blogs/internal/business/blog"
	"blogs/internal/common"
	filtermodel "blogs/internal/model/filter"
	blogstorage "blogs/internal/repository/mysql/blog"
	blogcachestorage "blogs/internal/repository/redis/blog"
	"github.com/redis/go-redis/v9"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListBlog(db *gorm.DB, rdb *redis.Client) func(c *gin.Context) {
	return func(c *gin.Context) {

		condition := map[string]interface{}{
			"deleted": false,
			"status":  []string{"pending", "active"},
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInternal(err))
			return
		}

		paging.Process()

		var filter filtermodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInternal(err))
			return
		}

		store := blogstorage.NewMysqlStorage(db)
		rdbStore := blogcachestorage.NewRedisStorage(rdb)
		biz := blogbusiness.NewListItemBlogBiz(store, rdbStore)
		data, err := biz.ListItem(c.Request.Context(), condition, &paging, &filter)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewDataResponse(data))
	}
}
