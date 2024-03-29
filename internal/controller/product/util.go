package product

import (
	"fmt"
	"mime/multipart"
	"path/filepath"

	errorCommon "github.com/aziemp66/byte-bargain/common/error"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *ProductController) saveFile(c *gin.Context, fileHeader *multipart.FileHeader) (filename string, err error) {
	if fileHeader.Size > MaxFileSize {
		return filename, errorCommon.NewInvariantError("file size exceeds the maximum limit")
	}

	ext := filepath.Ext(fileHeader.Filename)
	if ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
		return filename, errorCommon.NewInvariantError("only png, jpg or jpeg extension is supported")
	}

	filename = fmt.Sprintf("%s%s", uuid.NewString(), ext)
	fileLocation := BasePath + "/" + filename
	err = c.SaveUploadedFile(fileHeader, fileLocation)
	return filename, err
}
