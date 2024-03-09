package model

type {[.Name]} struct {
	Model
	Name string `gorm:"column:Name" json:"Name"` // 名称
	Sort int    `gorm:"column:Sort" json:"Sort"` // 排序
}
