package nc

import (
	"github.com/gin-gonic/gin"
	"go-starter/src/common"
	"net/http"
)

// GetUserHandler 获取用户信息
func GetUserHandler(c *gin.Context) {

	userId := c.Param("userId")

	g := LoginCloud()
	defer g.Logout()
	user, err := g.Users().Get(userId)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, user)

}

func GetAllUserHandler(c *gin.Context) {
	g := LoginCloud()
	defer g.Logout()
	user, err := g.Users().List()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, user)
}

func GetAllUserDetailsHandler(c *gin.Context) {
	g := LoginCloud()
	defer g.Logout()
	user, err := g.Users().ListDetails()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, user)
}

func CreateUserHandler(c *gin.Context) {
	var user *UserDetail
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.FailedByReqParam())
	}

	g := LoginCloud()
	defer g.Logout()

	err = g.Users().Create(user.Username, user.Password, user.UserDetails)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.DefaultFailed(err))
	}
	c.JSON(http.StatusOK, common.Ok(nil))
}

func DeleteUserHandler(c *gin.Context) {
	var user *UserDetail
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.FailedByReqParam())
	}

	g := LoginCloud()
	defer g.Logout()
	err = g.Users().Delete(user.Username)
	if err != nil {
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.DefaultFailed(err))
	}
	c.JSON(http.StatusOK, common.Ok(nil))
}

func UpdateUserHandler(c *gin.Context) {
	var user *UserDetail
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.FailedByReqParam())
	}
	g := LoginCloud()
	defer g.Logout()
	err = g.Users().Update(user.UserDetails)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.DefaultFailed(err))
	}

	c.JSON(http.StatusOK, common.Ok(nil))
}

func SendWelcomeEmail(c *gin.Context) {
	var user *UserDetail
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.FailedByReqParam())
	}
	g := LoginCloud()
	defer g.Logout()
	err = g.Users().SendWelcomeEmail(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.DefaultFailed(err))
	}
	c.JSON(http.StatusOK, common.Ok(nil))
}

func EnableUser(c *gin.Context) {
	var user *UserDetail
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.FailedByReqParam())
	}
	g := LoginCloud()
	defer g.Logout()
	err = g.Users().Enable(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.DefaultFailed(err))
	}
	c.JSON(http.StatusOK, common.Ok(nil))
}

func DisableUser(c *gin.Context) {
	var user *UserDetail
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.FailedByReqParam())
	}
	g := LoginCloud()
	defer g.Logout()
	err = g.Users().Disable(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.DefaultFailed(err))
	}
	c.JSON(http.StatusOK, common.Ok(nil))
}
