package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/model"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/pkg/request"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/log"
	"gorm.io/gorm/clause"

)

type {[.Name]}Repository interface {
	/* CRUDQ */
	// c
	Create(ctx context.Context,
		md *model.{[.Name]}) (*model.{[.Name]}, error)
	// r
	Get(ctx context.Context,
		id uint) (*model.{[.Name]}, error)
	// u
	Update(ctx context.Context,
		md *model.{[.Name]}) (*model.{[.Name]}, error)
	// d
	Delete(ctx context.Context,
		id uint) error
	// q
	Query(ctx context.Context,
		req *request.PublicQueryListRequest) (*[]model.{[.Name]}, int64, error)

}

func New{[.Name]}Repository(r *Repository) {[.Name]}Repository {
	return &{[.LowerName]}Repository{
		Repository: r,
	}
}

type {[.LowerName]}Repository struct {
	*Repository
}

/* CRUDQ */
// c
func (r *{[.LowerName]}Repository) Create(ctx context.Context,
	md *model.{[.Name]}) (*model.{[.Name]}, error) {
	if err := r.db.Create(md).Error; err != nil {
		log.Debug(err.Error())
		return nil, errors.New("添加到数据库失败")
	}
	return md, nil
}

// r
func (r *{[.LowerName]}Repository) Get(ctx context.Context,
	id uint) (*model.{[.Name]}, error) {

	md := model.{[.Name]}{}
		tx := r.db.Where("id = ?", id).Preload(clause.Associations).Find(&md)
	// log.Debug(tx.Error)
	// log.Debug(tx.RowsAffected)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, errors.New("无数据")
	}
	return &md, nil
}

// u
func (r *{[.LowerName]}Repository) Update(ctx context.Context,
	md *model.{[.Name]}) (*model.{[.Name]}, error) {

	if err := r.db.Save(md).Error; err != nil {
		return nil, errors.New("更新失败")
	}
	return md, nil
}

// d
func (r *{[.LowerName]}Repository) Delete(ctx context.Context,
	id uint) error {
	if err := r.db.Delete(&model.{[.Name]}{}, id).Error; err != nil {
		log.Debug(err.Error())
		return errors.New("无法删除")
	}
	return nil
}

// q
func (r *{[.LowerName]}Repository) Query(ctx context.Context,
	req *request.PublicQueryListRequest) (*[]model.{[.Name]}, int64, error) {
	var data []model.{[.Name]}

	tx := r.db.WithContext(ctx)
	// 是否需要排序
	if req.Sort != "" {
		tx = tx.Order(req.Sort)
	}

	// page
	if req.Page > 0 {
		tx = tx.Offset((req.Page - 1) * req.PageSize)
	}
	// page size
	if req.PageSize > 0 {
		tx = tx.Limit(req.PageSize)
	}

	// columns and query
	if len(req.Columns) > 0 {
		for _, item := range req.Columns {
			q := fmt.Sprintf("%s", item.Query)
			c := "="
			if item.Conditions != "" {
				c = item.Conditions
			}
			if item.Conditions == "LIKE" {
				q = fmt.Sprintf("%%%s%%", item.Query)
			}
			if item.Exp == "" || item.Exp == "and" {
				tx = tx.Where(fmt.Sprintf("%s %s ?", item.Field, c), q)
				continue
			}
			if item.Exp == "or" {
				tx = tx.Or(fmt.Sprintf("%s %s ?", item.Field, c), q)
				continue
			}
			if item.Exp == "not" {
				tx = tx.Not(fmt.Sprintf("%s %s ?", item.Field, c), q)
			}
		}
	}

	// Find
	// tx = tx.Preload(clause.Associations).Find(&data)
	tx = tx.Find(&data)
	if len(data) < 1 {
		return nil, 0, errors.New("没有数据")
	}
	var total int64
	tx.Offset(-1).Limit(-1).Count(&total)

	return &data, total, nil
}
