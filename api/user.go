package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"website_manager/service"
)

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var userLoginSrv service.UserLoginService
	if err := c.ShouldBindJSON(&userLoginSrv); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	} else {
		r := userLoginSrv.UserLogin()
		c.JSON(http.StatusOK, r)
	}
}
