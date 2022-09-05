package nc

import (
	"github.com/gin-gonic/gin"
	"go-starter/src/common"
	"net/http"
)

func ListFileOrFolderHandler(c *gin.Context) {
	g := LoginCloud()
	defer g.Logout()
	list, err := g.WebDav().ReadDir(c.DefaultQuery("path", "/"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.DefaultFailed(err))
	}
	var result = make([]map[string]interface{}, 1)
	for _, v := range list {
		var m = make(map[string]interface{}, 6)
		m["name"] = v.Name()
		m["size"] = v.Size()
		m["mode"] = v.Mode()
		m["modTime"] = v.ModTime()
		m["isDir"] = v.IsDir()
		m["sys"] = v.Sys()
		result = append(result, m)
	}

	c.JSON(http.StatusOK, common.Ok(result))
}

func CreateFileOrFolderHandler(c *gin.Context) {
	var fileParams *FileParams
	err := c.ShouldBind(&fileParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.FailedByReqParam())
	}
	g := LoginCloud()
	defer g.Logout()
	err = g.WebDav().MkdirAll(fileParams.Path, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.DefaultFailed(err))
	}
	c.JSON(http.StatusOK, common.Ok(nil))
}

func DeleteFileOrFolderHandler(c *gin.Context) {
	var fileParams *FileParams
	err := c.ShouldBind(&fileParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.FailedByReqParam())
	}
	g := LoginCloud()
	defer g.Logout()
	err = g.WebDav().RemoveAll(fileParams.Path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.DefaultFailed(err))
	}
	c.JSON(http.StatusOK, common.Ok(nil))
}

func RenameFileOrFolderHandler(c *gin.Context) {
	var fileParams *FileParams
	err := c.ShouldBind(&fileParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.FailedByReqParam())
	}
	g := LoginCloud()
	defer g.Logout()
	err = g.WebDav().Rename(fileParams.OldPath, fileParams.NewPath, fileParams.Overwrite)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.DefaultFailed(err))
	}
	c.JSON(http.StatusOK, common.Ok(nil))
}
func MoveFileOrFolderHandler(c *gin.Context) {
	var fileParams *FileParams
	err := c.ShouldBind(&fileParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.FailedByReqParam())
	}
	g := LoginCloud()
	defer g.Logout()
	err = g.WebDav().Copy(fileParams.OldPath, fileParams.NewPath, fileParams.Overwrite)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.DefaultFailed(err))
	}
	c.JSON(http.StatusOK, common.Ok(nil))
}

func UploadFileHandler(c *gin.Context) {
	var fileParams *FileParams
	err := c.ShouldBind(&fileParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.FailedByReqParam())
	}
	f, err := c.FormFile("file")
	data, err := f.Open()
	if err != nil {
		c.JSON(http.StatusOK, common.Msg{
			Code: 500,
			Msg:  "上传文件异常，请重新上传",
			Data: err.Error(),
		})
	}
	g := LoginCloud()
	defer g.Logout()
	err = g.WebDav().WriteStream(fileParams.Path+"/"+f.Filename, data, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.DefaultFailed(err))
	}
	c.JSON(http.StatusOK, common.Ok(nil))
}

func DownloadFileHandler(c *gin.Context) {
	fileDir := c.Query("path")
	fileName := c.Query("fileName")

	g := LoginCloud()
	defer g.Logout()
	response, err := g.WebDav().Read(fileDir + "/" + fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.DefaultFailed(err))
	}

	//非空处理
	if fileDir == "" || fileName == "" || response == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "失败",
			"error":   "资源不存在",
		})
	}
	c.Header("Transfer-Encoding", "true")
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Data(200, "application/octet-stream", response)
}
