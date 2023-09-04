package server

import (
	// "t/internal/handler"
	// "t/internal/pkg/middleware"
	// "t/pkg/helper/resp"
	// "t/pkg/jwt"
	// "t/pkg/log"

	"net/http"

	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version"
	"github.com/gin-gonic/gin"
)

func NewServerHTTP(
// jwt *jwt.JWT,
// userHandler handler.UserHandler,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	/*
	 * Basic routing
	 */
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "靓仔，你好！")
	})

	router.GET("/version", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, version.Version)
	})

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	return router
}
