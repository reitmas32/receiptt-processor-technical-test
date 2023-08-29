package config

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine {
	if router == nil {
		gin.SetMode(GIN_RUN_MODE)
		router = gin.Default()
	}
	return router
}

var router *gin.Engine
