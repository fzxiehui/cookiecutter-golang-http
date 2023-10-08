package repository

import (
	"context"
	"errors"
	"io"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
)

type FilesRepository interface {
	// 上传图片
	SaveImage(ctx context.Context, fileName string, fh *multipart.FileHeader) error
	// 下载图片
	GetImage(ctx context.Context, bulkName string, fileName string) ([]byte, error)
}

func NewFilesRepository(r *Repository) FilesRepository {
	return &filesRepository{
		Repository: r,
	}
}

type filesRepository struct {
	*Repository
}

func (r *filesRepository) SaveImage(ctx context.Context, fileName string, fh *multipart.FileHeader) error {
	file, err := fh.Open()
	if err != nil {
		return errors.New("打开文件失败")
	}
	_, err = r.oss.PutObject(ctx, "sys", fileName, file, fh.Size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return errors.New("文件上传oss失败")
	}

	return nil
}

func (r *filesRepository) GetImage(ctx context.Context, bulkName string, fileName string) ([]byte, error) {

	obj, err := r.oss.GetObject(ctx, bulkName, fileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, errors.New("无法获取oss对像")
	}
	buf, err := io.ReadAll(obj)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
