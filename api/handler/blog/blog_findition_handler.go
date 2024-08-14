package bloghandler

import (
	blogbusiness "blogs/internal/business/blog"
	"blogs/internal/common"
	blogstorage "blogs/internal/repository/mysql/blog"
	blogcachestorage "blogs/internal/repository/redis/blog"
	"github.com/redis/go-redis/v9"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBlog(db *gorm.DB, rdb *redis.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		store := blogstorage.NewMysqlStorage(db)
		rdbStore := blogcachestorage.NewRedisStorage(rdb)
		biz := blogbusiness.NewBlogBiz(store, rdbStore)

		data, err := biz.GetBlog(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewDataResponse(data))
	}
}
