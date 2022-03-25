package services

import (
	"github.com/peam1146/todo_api/src/databases"
	"github.com/peam1146/todo_api/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoServices interface {
	CreateTodo(todo models.Todo) (primitive.ObjectID, error)
	GetAllTodos() ([]models.Todo, error)
	EditTodo(id primitive.ObjectID, todo models.Todo) error
	DeleteTodo(id primitive.ObjectID) error
}

type todoServices struct {
	database databases.Databases
}

func NewDefaultTodoServices() TodoServices {
	return &todoServices{
		database: databases.GetClient(),
	}
}

func (t *todoServices) CreateTodo(todo models.Todo) (primitive.ObjectID, error) {
	return t.database.InsertTodo(bson.M{"title": todo.Title, "Done": todo.Done})
}

func (t *todoServices) GetAllTodos() ([]models.Todo, error) {
	todos := []models.Todo{}
	err := t.database.GetAllTodos(&todos)
	return todos, err
}

func (t *todoServices) EditTodo(id primitive.ObjectID, todo models.Todo) error {
	return t.database.UpdateTodo(id, bson.M{"title": todo.Title, "Done": todo.Done})
}
func (t *todoServices) DeleteTodo(id primitive.ObjectID) error {
	return t.database.DeleteTodo(id)
}
