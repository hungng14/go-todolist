package main

import (
	"fmt"
	"log"
	"todolist/database"
	task_handler "todolist/internal/module/task"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	fmt.Println("Welcome to Take Note application!")
	database.InitDB()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // This allows all origins
		AllowHeaders: "*", // Specify allowed headers
		AllowMethods: "*", // Specify allowed methods
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello There!")
	})

	app.Use("/api", func(c *fiber.Ctx) error {
		return c.Next()
	})

	task_handler.InitTaskHandler(app)

	log.Fatal(app.Listen(":8081"))
}
