package models

type User struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" gorm:"varchar(300)"`
	Password  string    `json:"password" gorm:"varchar(300)"`
}
