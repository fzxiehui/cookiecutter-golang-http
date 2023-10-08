package model

type {[.Name]} struct {
	Model
	Name              string           `gorm:"not null" json:"name"` // 名称
	Sort              int              `json:"sort"`                 // 排序
}

func (d *{[.Name]}) TableName() string {
	return "{[.LowerName]}s"
}
