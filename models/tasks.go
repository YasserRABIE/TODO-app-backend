package models

type Task struct {
	ID      uint16 `json:"id" gorm:"primaryKey;autoIncrement"`
	Title   string `json:"title"`
	Filter  string `json:"filter"`
	Color   string `json:"color"`
	User_id uint16 `json:"user_id" gorm:"foreignKey:ID"`
}

func (Task) TableName() string {
	return "tasks"
}

type AddTaskReq struct {
	Title  string `json:"title"`
	Filter string `json:"filter"`
	Color  string `json:"color"`
}

type RemoveTaskReq struct {
	Title  string `json:"title"`
	Filter string `json:"filter"`
}
type TasksByFilterReq struct {
	Filter string `json:"filter"`
}
