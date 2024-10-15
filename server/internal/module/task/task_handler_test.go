package task_handler_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"todolist/database"
	"todolist/internal/models"
	task_handler "todolist/internal/module/task"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"net/http"
	"net/http/httptest"
)

func setupTestApp() *fiber.App {

	app := fiber.New()

	database.InitDB()

	task_handler.InitTaskHandler(app)

	return app
}

func TestGetTasks(t *testing.T) {
	app := setupTestApp()

	db := database.DB.Db

	content := "Test content"
	db.Create(&models.Task{Title: "Test title", Content: &content})

	req := httptest.NewRequest(http.MethodGet, "/api/v1/tasks", nil)
	resp, _ := app.Test(req)

	// Validate the response
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var tasks []models.Task
	json.NewDecoder(resp.Body).Decode(&tasks)

	fmt.Println("tasks", tasks)

	var newTask models.Task

	for i, t := range tasks {
		if t.Title == "Test title" {
			newTask = tasks[i]
			break
		}
	}

	// Validate the response data
	assert.Equal(t, "Test title", newTask.Title)
}
