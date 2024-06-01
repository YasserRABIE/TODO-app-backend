package handlers

import (
	"net/http"

	"github.com/YasserRABIE/authentication-porject/database"
	"github.com/YasserRABIE/authentication-porject/models"
	"github.com/gin-gonic/gin"
)

func AddTask(c *gin.Context) {
	addReq := &models.AddTaskReq{}

	userIdInterface, _ := c.Get("id")
	userId, ok := userIdInterface.(uint16)
	if !ok {
		resBody := models.NewFailedResponse(http.StatusBadRequest, map[string]string{
			"error": "id is wrong",
		})

		c.JSON(http.StatusBadRequest, &resBody)
		return
	}

	if err := c.BindJSON(addReq); err != nil {
		resBody := models.NewFailedResponse(http.StatusNoContent, map[string]string{
			"error": "Invalid request! Please provide task data",
		})

		c.JSON(http.StatusNoContent, &resBody)
		return
	}

	_, err := database.CreateTask(addReq, userId)
	if err != nil {
		resBody := models.NewFailedResponse(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})

		c.JSON(http.StatusBadRequest, &resBody)
		return
	}

	resBody := models.NewSuccessResponse(http.StatusOK, map[string]interface{}{
		"message": "task is added successfully",
	})

	c.JSON(http.StatusOK, &resBody)
}

func RemoveTask(c *gin.Context) {
	removeReq := &models.RemoveTaskReq{}

	if err := c.BindJSON(removeReq); err != nil {
		resBody := models.NewFailedResponse(http.StatusNoContent, map[string]string{
			"error": "Invalid request! Please provide task title",
		})

		c.JSON(http.StatusNoContent, &resBody)
		return
	}

	if err := database.RemoveTask(removeReq); err != nil {
		resBody := models.NewFailedResponse(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})

		c.JSON(http.StatusBadRequest, &resBody)
		return
	}

	resBody := models.NewSuccessResponse(http.StatusOK, map[string]interface{}{
		"message": "task is removed successfully",
	})

	c.JSON(http.StatusOK, &resBody)
}

func GetAllTasks(c *gin.Context) {
	userIdInterface, _ := c.Get("id")
	userId, ok := userIdInterface.(uint16)
	if !ok {
		resBody := models.NewFailedResponse(http.StatusBadRequest, map[string]string{
			"error": "id is wrong",
		})

		c.JSON(http.StatusBadRequest, &resBody)
		return
	}

	tasks, err := database.GetTasks(userId)
	if err != nil {
		resBody := models.NewFailedResponse(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})

		c.JSON(http.StatusBadRequest, &resBody)
		return
	}

	resBody := models.NewSuccessResponse(http.StatusOK, map[string]interface{}{
		"tasks": tasks,
	})

	c.JSON(http.StatusOK, &resBody)
}
