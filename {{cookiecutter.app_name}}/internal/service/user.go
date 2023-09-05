package service

import (
	"context"
	"errors"

	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/model"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/repository"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/pkg/request"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, req *request.RegisterRequest) error
}

type userService struct {
	userRepo repository.UserRepository
	*Service
}

func NewUserService(server *Service, userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
		Service:  server,
	}
}

func (s *userService) Register(ctx context.Context, req *request.RegisterRequest) error {

	// fmt.Print("Register Service")
	// log.Debug("Register Service")
	// 检查用户名是否已存在
	if user, err := s.userRepo.GetByUsername(ctx, req.Username); err == nil && user != nil {
		return errors.New("username already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}
	user := &model.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Email:    req.Email,
	}

	if err = s.userRepo.Create(ctx, user); err != nil {
		return errors.New("failed to create user")
	}
	return nil
}

