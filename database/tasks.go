package database

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/YasserRABIE/authentication-porject/initializers"
	"github.com/YasserRABIE/authentication-porject/models"
)

var ctx = context.Background()

func CreateTask(r *models.AddTaskReq, id uint16) (*models.Task, error) {
	task := models.Task{
		Title:   r.Title,
		Filter:  r.Filter,
		Color:   r.Color,
		User_id: id,
	}

	initializers.DB.Create(&task)
	if task.ID == 0 {
		return nil, fmt.Errorf("failed to insert task")
	}

	// update cache
	var tasks []models.Task
	idString := strconv.FormatUint(uint64(id), 10)
	if cache, err := initializers.Cache.Get(ctx, "tasks:"+idString).Result(); err == nil {
		json.Unmarshal([]byte(cache), &tasks)
		tasks = append(tasks, task)
		jsonTasks, _ := json.Marshal(tasks)
		initializers.Cache.Set(ctx, "tasks:"+idString, jsonTasks, 0)
	}

	return &task, nil
}

func RemoveTask(r *models.RemoveTaskReq) error {
	// delete the task
	result := initializers.DB.
		Where("title = ?", r.Title).
		Where("filter = ?", r.Filter).
		First(&models.Task{}).
		Delete(&models.Task{})
	if result.RowsAffected == 0 {
		return fmt.Errorf("failed to delete task")
	}

	// delete all cache
	if err := initializers.Cache.FlushDB(ctx).Err(); err != nil {
		fmt.Println("failed to delete cache")
	}

	return nil
}

func GetTasks(id uint16) ([]models.Task, error) {
	var tasks []models.Task
	idString := strconv.FormatUint(uint64(id), 10)

	if cache, err := initializers.Cache.Get(ctx, "tasks:"+idString).Result(); err == nil {
		json.Unmarshal([]byte(cache), &tasks)
		return tasks, nil
	}

	if err := initializers.DB.
		Where("user_id = ?", idString).
		Find(&tasks).Error; err != nil {
		return nil, err
	}

	jsonTasks, _ := json.Marshal(tasks)
	initializers.Cache.Set(ctx, "tasks:"+idString, jsonTasks, 0).Err()

	return tasks, nil
}

func GetTasksByFilter(id uint16, filter string) ([]models.Task, error) {
	var tasks []models.Task
	idString := strconv.FormatUint(uint64(id), 10)

	if cache, err := initializers.Cache.Get(ctx, filter+":"+idString).Result(); err == nil {
		json.Unmarshal([]byte(cache), &tasks)
		return tasks, nil
	}

	if err := initializers.DB.
		Where("user_id = ? AND filter = ?", id, filter).
		Find(&tasks).Error; err != nil {
		return nil, err
	}

	jsonTasks, _ := json.Marshal(tasks)
	initializers.Cache.Set(ctx, filter+":"+idString, jsonTasks, 0).Err()

	return tasks, nil
}
