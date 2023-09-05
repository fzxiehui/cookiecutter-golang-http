//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/model"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/log"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Migrate struct {
	db *gorm.DB
}

func NewMigrate(db *gorm.DB) *Migrate {
	return &Migrate{
		db: db,
	}
}

func (m *Migrate) Run() {
	if err := m.db.AutoMigrate(&model.User{}); err != nil {
		log.Error("user migrate error")
	}
	log.Info("AutoMigrate End")
}

func newMigrate(*viper.Viper) (*Migrate, func(), error) {
	panic(wire.Build(
		RepositorySet,
		NewMigrate,
	))
}

