package model

type User struct {
	Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null"`
}

// func (u *User) TableName() string {
// 	return "users"
// }
