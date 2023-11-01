package service

import (
	"context"
	"errors"

	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/model"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/pkg/request"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/pkg/responses"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/repository"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/reflect_tools"
)

type {[.Name]}Service interface {
	/* Basic */
	// c
	Create(ctx context.Context,
		req *request.Create{[.Name]}Request) (*model.{[.Name]}, error)
	// r
	Get(ctx context.Context,
		id uint) (*model.{[.Name]}, error)
	// u
	Update(ctx context.Context,
		id uint, req *request.Update{[.Name]}Request) (*model.{[.Name]}, error)
	// d
	Delete(ctx context.Context,
		id uint) error
	// q
	Query(ctx context.Context,
		req *request.PublicQueryListRequest) (*responses.PublicQueryListResponses, error)
}

type {[.LowerName]}Service struct {
	Repo repository.{[.Name]}Repository
	*Service
}

func New{[.Name]}Service(server *Service, repo repository.{[.Name]}Repository) {[.Name]}Service {
	return &{[.LowerName]}Service{
		Repo:    repo,
		Service: server,
	}
}
/* Basic */
// c
func (s *{[.LowerName]}Service) Create(ctx context.Context,
	req *request.Create{[.Name]}Request) (*model.{[.Name]}, error) {
	var md model.{[.Name]}
	reflect_tools.StructAssign(&md, req)
	return s.Repo.Create(ctx, &md)
}

// r
func (s *{[.LowerName]}Service) Get(ctx context.Context,
	id uint) (*model.{[.Name]}, error) {
	return s.Repo.Get(ctx, id)
}

// u
func (s *{[.LowerName]}Service) Update(ctx context.Context,
	id uint, req *request.Update{[.Name]}Request) (*model.{[.Name]}, error) {
	md, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, errors.New("资源不存在")
	}
	reflect_tools.StructAssign(md, req)
	return s.Repo.Update(ctx, md)
}

// d
func (s *{[.LowerName]}Service) Delete(ctx context.Context,
	id uint) error {
	return s.Repo.Delete(ctx, id)
}

// q
func (s *{[.LowerName]}Service) Query(ctx context.Context,
	req *request.PublicQueryListRequest) (*responses.PublicQueryListResponses, error) {
	return s.Repo.Query(ctx, req)
}
