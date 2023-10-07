package service

import (
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/jwt"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/sid"
)

type Service struct {
	sid *sid.Sid
	jwt *jwt.JWT
}

func NewService(sid *sid.Sid, jwt *jwt.JWT) *Service {
	return &Service{
		sid: sid, // 雪花
		jwt: jwt,
	}
}
