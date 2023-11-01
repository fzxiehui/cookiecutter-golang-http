package model

type {[.Name]} struct {
	Model
	Name string `gorm:"comment:'名称'" json:"name"` // 名称
	Sort int    `gorm:"comment:'排序'" json:"sort"` // 排序
}
