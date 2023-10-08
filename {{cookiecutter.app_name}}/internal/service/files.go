package service

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"path"

	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/repository"
)

type FilesService interface {
	// 上传图片
	SaveImage(ctx context.Context, id uint, fh *multipart.FileHeader) (string, error)
	// 下载图片
	GetImage(ctx context.Context, bulkName string, fileName string) ([]byte, error)
}

type filesService struct {
	Repo repository.FilesRepository
	*Service
}

func NewFilesService(server *Service, repo repository.FilesRepository) FilesService {
	return &filesService{
		Repo:    repo,
		Service: server,
	}
}

func (s *filesService) SaveImage(ctx context.Context, id uint, fh *multipart.FileHeader) (string, error) {

	sid, err := s.sid.GenUint64()
	if err != nil {
		return "", errors.New("生成 sid 失败")
	}
	fileExt := path.Ext(fh.Filename)
	// fileName := fmt.Sprintf("/download/image/sys/%d/%d%s", id, sid, fileExt)
	fileName := fmt.Sprintf("image/%d/%d%s", id, sid, fileExt)
	data := s.Repo.SaveImage(ctx, fileName, fh)
	ret_fileName := fmt.Sprintf("/download/image/sys/%d/%d%s", id, sid, fileExt)
	return ret_fileName, data
}

// 下载图片
func (s *filesService) GetImage(ctx context.Context, bulkName string, fileName string) ([]byte, error) {
	return s.Repo.GetImage(ctx, bulkName, fileName)
}
