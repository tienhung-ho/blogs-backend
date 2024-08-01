package accountsmiddlewares

import (
	"blogs/internal/common"
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FileUploadMiddileware() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInternal(err))
			c.Abort()
			return
		}

		src, err := file.Open()

		if err != nil {
			c.JSON(http.StatusInternalServerError, common.ErrInternal(err))
			c.Abort()
			return
		}

		defer src.Close()

		fileBuffer := new(bytes.Buffer)

		_, err = io.Copy(fileBuffer, src)

		if err != nil {
			c.JSON(http.StatusInternalServerError, common.ErrInternal(err))
			c.Abort()
			return
		}

		fileData := fileBuffer.Bytes()

		fileURL, err := common.UploadImageToS3(fileData, file.Filename)

		if err != nil {
			c.JSON(http.StatusInternalServerError, common.ErrInternal(err))
			c.Abort()
			return
		}

		c.Set("fileURL", fileURL)

		c.Next()
	}
}
