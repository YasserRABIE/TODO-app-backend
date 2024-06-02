package database

import (
	"fmt"

	"github.com/YasserRABIE/authentication-porject/initializers"
	"github.com/YasserRABIE/authentication-porject/models"
)

func CreateTask(r *models.AddTaskReq, id uint16) (*models.Task, error) {
	task := &models.Task{
		Title:   r.Title,
		Filter:  r.Filter,
		Color:   r.Color,
		User_id: id,
	}

	initializers.DB.Create(task)
	if task.ID == 0 {
		return nil, fmt.Errorf("failed to insert task")
	}

	return task, nil
}

func RemoveTask(r *models.RemoveTaskReq) error {
	result := initializers.DB.
		Where("title = ?", r.Title).
		Where("filter = ?", r.Filter).
		First(&models.Task{}).
		Delete(&models.Task{})

	if result.RowsAffected == 0 {
		return fmt.Errorf("failed to delete task")
	}

	return nil
}

func GetTasks(id uint16) ([]models.Task, error) {
	var tasks []models.Task

	if err := initializers.DB.
		Where("user_id = ?", id).
		Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTasksByFilter(id uint16, filter string) ([]models.Task, error) {
	var tasks []models.Task

	if err := initializers.DB.
		Where("user_id = ? AND filter = ?", id, filter).
		Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}
