package handler

import (
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}
func GetUserIdFromCtx(ctx *gin.Context) uint {
	v, exists := ctx.Get("claims")
	if !exists {
		return 0
	}
	return v.(*jwt.MyCustomClaims).UserId
}
