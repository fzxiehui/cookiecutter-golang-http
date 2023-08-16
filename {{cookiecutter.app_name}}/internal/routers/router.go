package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	/*
	 * Public
	 */
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	return router
}
