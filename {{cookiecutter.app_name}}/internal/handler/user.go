package handler

import (
	"net/http"

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
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.userService.Register(ctx, req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Register Success"})
}

