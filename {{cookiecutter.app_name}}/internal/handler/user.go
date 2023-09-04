package handler

import (
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/pkg/request"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Register(ctx *gin.Context)
	// Login(ctx *gin.Context)
	// GetProfile(ctx *gin.Context)
	// UpdateProfile(ctx *gin.Context)
}

func NewUserHandler(handler *Handler, userService service.UserService) UserHandler {
	return &userHandler{
		Handler:     handler,
		userService: userService,
	}
}

type userHandler struct {
	*Handler
	userService service.UserService
}

func (h *userHandler) Register(ctx *gin.Context) {
	req := &request.RegisterRequest{}
	h.userService.Register(ctx, req)

}
