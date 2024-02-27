package model

type {[.Name]} struct {
	Model
	Name string `gorm:"comment:'名称'" json:"Name"` // 名称
	Sort int    `gorm:"comment:'排序'" json:"Sort"` // 排序
}
