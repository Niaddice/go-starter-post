package nc

import (
	"github.com/gin-gonic/gin"
	"gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud/types"
	"go-starter/src/common"
	"net/http"
	"strconv"
)

func GetAllShareHandler(c *gin.Context) {
	g := LoginCloud()
	defer g.Logout()
	list, err := g.Shares().List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.DefaultFailed(err))
	}
	c.JSON(http.StatusOK, common.Ok(list))
}

func ShareHandler(c *gin.Context) {
	var shareParams *ShareParams
	err := c.ShouldBind(&shareParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.FailedByReqParam())
	}
	g := LoginCloud()
	defer g.Logout()
	result, err := g.Shares().Create(shareParams.Path, shareParams.ShareType, shareParams.Permissions, shareParams.ShareWith, shareParams.PublicUpload, shareParams.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.DefaultFailed(err))
	}
	c.JSON(http.StatusOK, common.Ok(result))
}

func DeleteShareHandler(c *gin.Context) {
	shareId := c.Param("shareId")
	if shareId == "" {
		c.JSON(http.StatusInternalServerError, common.FailedByReqParam())
	}
	g := LoginCloud()
	defer g.Logout()
	_shareId, err := strconv.Atoi(shareId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.FailedByReqParam())
	}
	g.Shares().Delete(_shareId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.DefaultFailed(err))
	}
	c.JSON(http.StatusOK, common.Ok(nil))
}

func UpdateShareHandler(c *gin.Context) {
	var shareParams *ShareParams
	err := c.ShouldBind(&shareParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.FailedByReqParam())
	}
	g := LoginCloud()
	defer g.Logout()
	var shareUpdate = types.ShareUpdate{
		ShareID:      shareParams.ShareId,
		Permissions:  shareParams.Permissions,
		Password:     shareParams.Password,
		PublicUpload: shareParams.PublicUpload,
		ExpireDate:   shareParams.ExpireDate,
	}
	err = g.Shares().Update(shareUpdate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.DefaultFailed(err))
	}
	c.JSON(http.StatusOK, common.Ok(nil))
}
