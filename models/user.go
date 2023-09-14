package models

type User struct {
	Id       int64  `json:"id" gorm:"primarykey"`
	Username string `json:"username" gorm:"varchar(300)"`
	Password string `json:"password" gorm:"varchar(300)"`
}
