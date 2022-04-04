package api

import "github.com/gin-gonic/gin"

func LoadRouter() *gin.Engine {
	r := gin.Default()
	r.POST("user/login", UserLogin)
	return r
}
