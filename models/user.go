package models

type User struct {
	ID       uint16 `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "users"
}
