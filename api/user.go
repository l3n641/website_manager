package api

import (
	"github.com/gin-gonic/gin"
)

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	c.String(200, "Hello, Geektutu")
}
