package models

type User struct {
	ID       uint16 `json:"id" gorm:"primaryKey"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "users"
}

type AccountRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}
