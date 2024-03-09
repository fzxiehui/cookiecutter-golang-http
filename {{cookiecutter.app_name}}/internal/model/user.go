package model

type User struct {
	Model
	Username string `gorm:"column:Username; unique;not null"`
	Password string `gorm:"column:Password; not null"`
	Email    string `gorm:"column:Email; not null"`
}

func (u *User) TableName() string {
	return "User"
}
