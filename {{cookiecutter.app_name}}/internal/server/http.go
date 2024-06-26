package server

import (
	"net/http"

	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/docs"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/handler"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/pkg/middleware"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/jwt"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewServerHTTP(
	jwt *jwt.JWT,
	userHandler handler.UserHandler,
	filesHandler handler.FilesHandler,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	// router := gin.Default()
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.DocExpansion("none"),
	))
	docs.SwaggerInfo.Title = "{{cookiecutter.app_name}}"
	docs.SwaggerInfo.Version = version.Version
	docs.SwaggerInfo.Description = "赤诚勇敢，自洽欢喜。"


	/*
	 * Basic routing
	 */
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "靓仔，你好！")
	})

	router.GET("/version", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, version.Version)
	})

	/*
	 * Files
	 */
	router.POST("/upload/image", filesHandler.SaveImage)
	router.GET("/download/image/:bulk/:uid/:name", filesHandler.GetImage)


	router.POST("/register", userHandler.Register)

	router.GET("/auth_test", middleware.StrictAuth(jwt), func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "靓仔，你好！")
	})

	// v1 := router.Group("/v1").Use(middleware.StrictAuth(jwt))
	// {}

	return router
}

