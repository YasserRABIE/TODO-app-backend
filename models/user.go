package models

type User struct {
	ID       uint16 `json:"id" gorm:"primaryKey;column:id"`
	Name     string `json:"name" gorm:"unique;column:name"`
	Email    string `json:"email" gorm:"unique;column:email"`
	Password string `json:"password" gorm:"column:password"`
}

func (User) TableName() string {
	return "users"
}

type AccountRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
