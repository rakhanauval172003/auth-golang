package models

type User struct {
	ID       uint   `json:"id" gorm:"primarykey"`
	NAME     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
