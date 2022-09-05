package nc

import (
	"github.com/gin-gonic/gin"
	"go-starter/src/common"
	"net/http"
)

func GetAllGroupHandler(c *gin.Context) {
	g := LoginCloud()
	defer g.Logout()
	list, err := g.Groups().List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.DefaultFailed(err))
	}
	c.JSON(http.StatusOK, common.Ok(list))
}

func GetGroupDetailHandler(c *gin.Context) {
	g := LoginCloud()
	defer g.Logout()
	details, err := g.Groups().ListDetails(c.Param("groupName"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.DefaultFailed(err))
	}
	c.JSON(http.StatusOK, common.Ok(details))
}

func GetUserGroupHandler(c *gin.Context) {
	g := LoginCloud()
	defer g.Logout()
	result, err := g.Groups().Users(c.Param("username"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.DefaultFailed(err))
	}
	c.JSON(http.StatusOK, common.Ok(result))
}

func CreateGroupHandler(c *gin.Context) {
	var groupParams *GroupParams
	err := c.ShouldBind(&groupParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.FailedByReqParam())
	}
	g := LoginCloud()
	defer g.Logout()
	err = g.Groups().Create(groupParams.Group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.DefaultFailed(err))
	}
	c.JSON(http.StatusOK, common.Ok(nil))

}
func DeleteGroupHandler(c *gin.Context) {
	g := LoginCloud()
	defer g.Logout()
	err := g.Groups().Delete(c.Param("groupName"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.DefaultFailed(err))
	}
	c.JSON(http.StatusOK, common.Ok(nil))
}

func AddInGroupHandler(c *gin.Context) {
	var groupParams *GroupParams
	err := c.ShouldBind(&groupParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.FailedByReqParam())
	}
	g := LoginCloud()
	defer g.Logout()
	err = g.Users().GroupAdd(groupParams.Username, groupParams.Group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.DefaultFailed(err))
	}
	c.JSON(http.StatusOK, common.Ok(nil))
}

func RemoveOutGroupHandler(c *gin.Context) {
	var groupParams *GroupParams
	err := c.ShouldBind(&groupParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.FailedByReqParam())
	}
	g := LoginCloud()
	defer g.Logout()
	err = g.Users().GroupRemove(groupParams.Username, groupParams.Group)

	if err != nil {
		c.JSON(http.StatusInternalServerError, common.DefaultFailed(err))
	}
	c.JSON(http.StatusOK, common.Ok(nil))
}
