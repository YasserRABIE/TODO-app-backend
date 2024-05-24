package models

type User struct {
	ID       uint16 `json:"id" gorm:"primaryKey;column:id"`
	UserName string `json:"userName" gorm:"unique;column:userName"`
	Email    string `json:"email" gorm:"unique;column:email"`
	Password string `json:"password" gorm:"column:password"`
}

func (User) TableName() string {
	return "users"
}

type AccountRequest struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
