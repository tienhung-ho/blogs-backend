package permissionhandler

import (
	permissionbusiness "blogs/internal/business/permission"
	"blogs/internal/common"
	filtermodel "blogs/internal/model/filter"
	permissionstorage "blogs/internal/repository/mysql/permission"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		c.JSON(http.StatusBadRequest, "13131231231")

		// var paging common.Paging

		// if err := c.ShouldBind(&paging); err != nil {
		// 	c.JSON(http.StatusBadRequest, common.ErrInternal(err))
		// 	return
		// }

		// var filter filtermodel.Filter

		// if err := c.ShouldBind(&filter); err != nil {
		// 	c.JSON(http.StatusBadRequest, common.ErrInternal(err))
		// 	return
		// }

		// store := permissionstorage.NewMysqlStorage(db)
		// biz := permissionbusiness.NewPermissionListItem(store)
		// log.Printf("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		// permissions, err := biz.ListPermissions(c.Request.Context(), map[string]interface{}{"deleted": false}, &paging, &filter)

		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, err)
		// 	return
		// }

		// c.JSON(http.StatusOK, common.NewSuccesResponse(permissions, paging, filter))
	}
}

func ListPermissions(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInternal(err))
			return
		}

		var filter filtermodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInternal(err))
			return
		}

		store := permissionstorage.NewMysqlStorage(db)
		biz := permissionbusiness.NewPermissionListItem(store)
		permissions, err := biz.ListPermissions(c.Request.Context(), map[string]interface{}{"deleted": false}, &paging, &filter)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewSuccesResponse(permissions, paging, filter))
	}
}
