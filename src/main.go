package main

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/peam1146/todo_api/src/controllers"
	"github.com/peam1146/todo_api/src/databases"
	"github.com/peam1146/todo_api/src/utils"
)

func main() {
	defer databases.GetClient().Close()

	app := fiber.New(fiber.Config{
		AppName: "TODO API v1",
	})

	api := app.Group("/api")
	port := utils.Getenv("PORT", "3000")

	todo := api.Group("/todo")

	todoCtr := controller.NewDefaultTodoControllers()
	todo.Get("/", todoCtr.GetAllTodos)
	todo.Post("/create", todoCtr.CreateTodo)
	todo.Delete("/:id", todoCtr.DeleteTodo)
	todo.Post("/update", todoCtr.EditTodo)

	api.Get("/status", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Listen(":" + port)
}
