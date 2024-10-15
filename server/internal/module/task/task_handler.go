package task_handler

import (
	"todolist/database"
	"todolist/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func InitTaskHandler(app *fiber.App) {

	app.Get("/api/v1/tasks", func(c *fiber.Ctx) error {

		db := database.DB.Db

		var tasks = []models.Task{}

		db.Find(&tasks)

		return c.JSON(tasks)

	})

	app.Post("/api/v1/tasks", func(c *fiber.Ctx) error {

		newTask := &models.Task{}
		if err := c.BodyParser(newTask); err != nil {
			return err
		}

		db := database.DB.Db

		// Save the new task to the database
		if err := db.Create(&newTask).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Failed to create new task",
			})
		}

		return c.JSON(newTask)
	})

	app.Put("/api/v1/tasks/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		taskData := &models.Task{}

		if err := c.BodyParser(taskData); err != nil {
			return err
		}

		db := database.DB.Db

		var task models.Task

		db.Find(&task, "id = ?", id)

		if task.Id == uuid.Nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Task not found"})
		}

		if taskData.Title != "" {
			task.Title = taskData.Title
		}

		if taskData.Content != nil {
			task.Content = taskData.Content
		}

		db.Save(&task)

		return c.JSON(task)
	})

	app.Patch("/api/v1/tasks/:id/done", func(c *fiber.Ctx) error {
		id := c.Params("id")

		var task models.Task
		db := database.DB.Db
		db.Find(&task, "id = ?", id)

		if task.Id == uuid.Nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Task not found"})
		}

		task.Done = true

		db.Save(&task)

		return c.JSON(task)
	})

	app.Delete("/api/v1/tasks/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		var task models.Task
		db := database.DB.Db
		db.Find(&task, "id = ?", id)

		if task.Id == uuid.Nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Task not found"})
		}

		err := db.Delete(&task, "id = ?", id).Error

		if err != nil {
			return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Failed to delete task"})
		}

		return c.JSON(task)
	})
}
