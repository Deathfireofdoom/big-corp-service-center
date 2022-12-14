package router

import (
	"github.com/gin-gonic/gin"
)

func Configure() *gin.Engine {
	router := gin.Default()

	router.GET("/test", connectionTest)
	router.GET("/test-the-manager", theManagerTest)

	return router
}
