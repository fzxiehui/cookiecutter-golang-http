package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func StrictAuth(j *jwt.JWT) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" {
			tokenString, _ = ctx.Cookie("Authorization")
		}
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "请登录！"})
			ctx.Abort()
			return
		}
		claims, err := j.ParseToken(tokenString)
		if err != nil {
			// resp.HandleError(ctx, http.StatusUnauthorized, 1, err.Error(), nil)
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
		// recoveryLoggerFunc(ctx, logger)
		ctx.Next()
	}
}

func NoStrictAuth(j *jwt.JWT) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" {
			tokenString, _ = ctx.Cookie("accessToken")
		}
		if tokenString == "" {
			tokenString = ctx.Query("accessToken")
		}
		if tokenString == "" {
			ctx.Next()
			return
		}

		claims, err := j.ParseToken(tokenString)
		if err != nil {
			ctx.Next()
			return
		}

		ctx.Set("claims", claims)
		// recoveryLoggerFunc(ctx, logger)
		ctx.Next()
	}
}

func recoveryLoggerFunc(ctx *gin.Context, logger *log.Logger) {
	userInfo := ctx.MustGet("claims").(*jwt.MyCustomClaims)

	fmt.Print(userInfo)
}
