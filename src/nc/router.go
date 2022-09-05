package nc

import (
	"github.com/gin-gonic/gin"
	"go-starter/src/middleware"
)

func Routers(e *gin.Engine) {
	api := e.Group("/api", middleware.JWTAuthMiddleware())
	{
		user := api.Group("/user")
		{
			user.GET("/info/:userId", GetUserHandler)
			user.GET("/info/all", GetAllUserHandler)
			user.GET("/info/allDetail", GetAllUserDetailsHandler)
			user.POST("/create", CreateUserHandler)
			user.DELETE("/delete", DeleteUserHandler)
			user.PUT("/update", UpdateUserHandler)
			user.POST("/sendWelcomeEmail", SendWelcomeEmail)
			user.PUT("/enable", EnableUser)
			user.PUT("/disable", DisableUser)
			//TODO 企业微信监听用户新增/删除/更新
		}
		group := api.Group("/group")
		{
			group.POST("/inGroup", AddInGroupHandler)
			group.DELETE("/outGroup", RemoveOutGroupHandler)
			group.GET("/list", GetAllGroupHandler)
			group.GET("/detail/:groupName", GetGroupDetailHandler)
			group.GET("/list/:username", GetUserGroupHandler)
			group.POST("/create", CreateGroupHandler)
			group.DELETE("/delete/:groupName", DeleteGroupHandler)
		}

		file := api.Group("/webdav")
		{
			file.GET("/list", ListFileOrFolderHandler)
			file.POST("/create", CreateFileOrFolderHandler)
			file.DELETE("/delete", DeleteFileOrFolderHandler)
			file.PUT("/rename", RenameFileOrFolderHandler)
			file.PUT("/move", MoveFileOrFolderHandler)
			file.POST("/upload", UploadFileHandler)
			file.GET("/download", DownloadFileHandler)
		}

		share := api.Group("/share")
		{
			share.POST("/", ShareHandler)
			share.GET("/list", GetAllShareHandler)
			share.DELETE("/delete/:shareId", DeleteShareHandler)
			share.PUT("/update", UpdateShareHandler)
		}
	}

}
