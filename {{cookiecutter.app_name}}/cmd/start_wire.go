//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/handler"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/repository"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/server"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/service"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/jwt"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/sid"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var RepositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewUserRepository,
)

func newHttpServe(*viper.Viper) (*server.Server, func(), error) {

	panic(wire.Build(
		RepositorySet,
		ServiceSet,
		HandlerSet,
		server.NewServer,
		server.NewServerHTTP,
		jwt.NewJwt,
		sid.NewSid,
	))
}

