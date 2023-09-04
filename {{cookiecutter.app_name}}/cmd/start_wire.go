//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/server"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/jwt"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func newHttpServe(*viper.Viper) (*server.Server, func(), error) {

	panic(wire.Build(
		server.NewServer,
		server.NewServerHTTP,
		jwt.NewJwt,
	))
}
