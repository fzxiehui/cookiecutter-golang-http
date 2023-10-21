//go:build wireinject
// +build wireinject

package cmd

import (
	"context"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/model"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/log"
	"github.com/google/wire"
	"github.com/minio/minio-go/v7"
	// "github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"io/ioutil"
	"strings"
)

type Migrate struct {
	db *gorm.DB
	oss     *minio.Client
	config  *viper.Viper
	sqlfile []string
}

func NewMigrate(db *gorm.DB, oss *minio.Client, config *viper.Viper, sqlfile []string) *Migrate {
	return &Migrate{
		db:      db,
		oss:     oss,
		config:  config,
		sqlfile: sqlfile,
	}
}

func (m *Migrate) Run() {
	// oss Init
	// system bulck
	is_exist, _ := m.oss.BucketExists(context.Background(), "sys")
	if !is_exist {
		bucket := "sys"
		location := "us-east-1"
		opt := minio.MakeBucketOptions{Region: location}
		err := m.oss.MakeBucket(context.Background(), bucket, opt)
		if err != nil {
			log.Fatal("MakeBucket sys Error!")
			return
		}
	}

	tables := []interface{}{
		model.User{},    // 用户表(客户)
	}

	/*
	 * 删除已有的表, 创建新表
	 */
	log.Info("------------------ DropTable ------------------")
	for _, t := range tables {
		if m.db.Migrator().HasTable(&t) {
			_ = m.db.Migrator().DropTable(&t)
		}
	}

	log.Info("------------------ AutoMigrate ------------------")
	for _, t := range tables {
		err := m.db.AutoMigrate(&t)
		if err != nil {
			log.Info(err.Error)
			return
		}
	}


	/*
	 * 以下对 SQL 文件初始化
	 */

	log.Info(m.sqlfile)
	if len(m.sqlfile) > 0 {
		log.Info("=========================================================")
		log.Info("================Init SQL File============================")
		log.Info("=========================================================")
		for _, f := range m.sqlfile {
			log.Info("Sql File:", f)
			bytes, err := ioutil.ReadFile(f)
			if err != nil {
				log.Error(err)
				continue
			}
			sql_arr := strings.Split(string(bytes), ";")
			for _, sql := range sql_arr {

				if strings.TrimSpace(sql) == "" {
					continue
				}

				if err := m.db.Exec(string(sql)).Error; err != nil {
					log.Error(err)
				}
			}
		}
	}

	log.Info("AutoMigrate End")
}

func newMigrate(*viper.Viper, []string) (*Migrate, func(), error) {
	panic(wire.Build(
		RepositorySet,
		NewMigrate,
	))
}

