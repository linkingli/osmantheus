package gin_server

import (
	"github.com/gin-gonic/gin"
	"osmantheus/gin-server/v1"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.POST("/file/upload", v1.UpLoadFile)
	router.GET("/file/download", v1.Dowloadfile)
	return router
}
