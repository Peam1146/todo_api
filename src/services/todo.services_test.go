package services

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/peam1146/todo_api/src/databases"
	mock_databases "github.com/peam1146/todo_api/src/databases/mocks"
	"github.com/peam1146/todo_api/src/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_todoServices_GetAllTodos(t *testing.T) {
	type fields struct {
		database databases.Databases
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := mock_databases.NewMockDatabases(ctrl)
	db.EXPECT().GetAllTodos(gomock.Any()).Return(nil)
	tests := []struct {
		name    string
		fields  fields
		want    []models.Todo
		wantErr bool
	}{
		{
			name: "Get all todos",
			fields: fields{
				database: db,
			},
			want:    []models.Todo{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &todoServices{
				database: tt.fields.database,
			}
			got, err := tr.GetAllTodos() // got = []
			if (err != nil) != tt.wantErr {
				t.Errorf("todoServices.GetAllTodos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("todoServices.GetAllTodos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todoServices_CreateTodo(t *testing.T) {
	type fields struct {
		database databases.Databases
	}
	type args struct {
		todo models.Todo
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	todo := models.Todo{
		ID:    primitive.NewObjectID(),
		Title: "test",
		Done:  false,
	}

	db := mock_databases.NewMockDatabases(ctrl)
	db.EXPECT().InsertTodo(gomock.Any()).Return(todo.ID, nil)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    primitive.ObjectID
		wantErr bool
	}{
		{
			name: "Create todo",
			fields: fields{
				database: db,
			},
			args: args{
				todo: todo,
			},
			want:    todo.ID,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &todoServices{
				database: tt.fields.database,
			}
			got, err := tr.CreateTodo(tt.args.todo)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoServices.CreateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("todoServices.CreateTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todoServices_EditTodo(t *testing.T) {
	type fields struct {
		database databases.Databases
	}
	type args struct {
		id   primitive.ObjectID
		todo models.Todo
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	todo := models.Todo{
		ID:    primitive.NewObjectID(),
		Title: "test",
		Done:  false,
	}

	db := mock_databases.NewMockDatabases(ctrl)
	db.EXPECT().UpdateTodo(todo.ID, gomock.Any()).Return(nil)
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Edit todo",
			fields: fields{
				database: db,
			},
			args: args{
				id:   todo.ID,
				todo: todo,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &todoServices{
				database: tt.fields.database,
			}
			if err := tr.EditTodo(tt.args.id, tt.args.todo); (err != nil) != tt.wantErr {
				t.Errorf("todoServices.EditTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_todoServices_DeleteTodo(t *testing.T) {
	type fields struct {
		database databases.Databases
	}
	type args struct {
		id primitive.ObjectID
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	todo := models.Todo{
		ID:    primitive.NewObjectID(),
		Title: "test",
		Done:  false,
	}

	db := mock_databases.NewMockDatabases(ctrl)
	db.EXPECT().DeleteTodo(todo.ID).Return(nil)
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Delete todo",
			fields: fields{
				database: db,
			},
			args: args{
				id: todo.ID,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &todoServices{
				database: tt.fields.database,
			}
			if err := tr.DeleteTodo(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("todoServices.DeleteTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
