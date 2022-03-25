package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/peam1146/todo_api/src/models"
	"github.com/peam1146/todo_api/src/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoControllers interface {
	GetAllTodos(ctx *fiber.Ctx) error
	CreateTodo(ctx *fiber.Ctx) error
	EditTodo(ctx *fiber.Ctx) error
	DeleteTodo(ctx *fiber.Ctx) error
}

type todoControllers struct {
	todoServices services.TodoServices
}

func NewDefaultTodoControllers() TodoControllers {
	return &todoControllers{
		todoServices: services.NewDefaultTodoServices(),
	}
}

func (t *todoControllers) GetAllTodos(ctx *fiber.Ctx) error {
	todos, err := t.todoServices.GetAllTodos()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(todos)
}

func (t *todoControllers) CreateTodo(ctx *fiber.Ctx) error {
	todo := models.Todo{}
	if err := ctx.BodyParser(&todo); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	id, err := t.todoServices.CreateTodo(todo)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(id)
}

func (t *todoControllers) EditTodo(ctx *fiber.Ctx) error {
	todo := models.Todo{}
	if err := ctx.BodyParser(&todo); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := t.todoServices.EditTodo(todo.ID, todo); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(fiber.Map{"status": "success"})
}

func (t *todoControllers) DeleteTodo(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(400).JSON(fiber.Map{"error": "id is required"})
	}

	new_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if err := t.todoServices.DeleteTodo(new_id); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"status": "success"})
}
