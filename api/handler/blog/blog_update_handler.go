package bloghandler

import (
	blogbusiness "blogs/internal/business/blog"
	"blogs/internal/common"
	blogmodel "blogs/internal/model/blog"
	blogstorage "blogs/internal/repository/mysql/blog"
	blogcachestorage "blogs/internal/repository/redis/blog"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func UpdateBlog(db *gorm.DB, rdb *redis.Client) func(c *gin.Context) {
	return func(c *gin.Context) {

		var data blogmodel.BlogUpdate

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		store := blogstorage.NewMysqlStorage(db)
		rdbStore := blogcachestorage.NewRedisStorage(rdb)
		biz := blogbusiness.NewUpdateBlogBiz(store, rdbStore)

		if err := biz.UpdateBlog(c.Request.Context(), map[string]interface{}{"id": id}, &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewDataResponse(true))
	}
}
