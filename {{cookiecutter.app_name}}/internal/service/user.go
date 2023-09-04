package service

import (
	"context"

	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/pkg/request"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/log"
)

type UserService interface {
	Register(ctx context.Context, req *request.RegisterRequest) error
}

type userService struct {
	// UserRepo repository.UserRepository
	*Service
}

func NewUserService(server *Service /* userRepo repository.UserRepository */) UserService {
	return &userService{
		Service: server,
	}
}

func (s *userService) Register(ctx context.Context, req *request.RegisterRequest) error {

	// fmt.Print("Register Service")
	log.Debug("Register Service")
	return nil
}
