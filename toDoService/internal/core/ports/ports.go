package ports

import (
	"github.com/google/uuid"
	"toDoService/internal/core/domain"
	"toDoService/internal/core/dtos"
	"toDoService/internal/core/dtos/Requests"
)

type ToDoService interface {
	GetTodoByUserId(request *Requests.GetUserTodosRequest) (*[]domain.Todo, *error)
	UpdateTodo(request *Requests.UpdateTodoRequest) *error
	CreateTodo(request *Requests.CreateTodoRequest) *error
	DeleteTodo(request *Requests.DeleteTodoRequest) *error
}

type ToDoRepository interface {
	FindTodosByUserId(id uuid.UUID) (*[]domain.Todo, *error)
	CreateTodo(todo *domain.Todo) *error
	UpdateTodo(todo *domain.Todo) *error
	DeleteTodo(todo *domain.Todo) *error
	FindTodoById(todoId uuid.UUID) (*domain.Todo, *error)
}

// does this make sense ?
// idk really.
// my inital thougt is maybe switching the messaging protocol between services from http to grpc or rmq
type CoomunactionHandler interface {
	AuthorizeUserByToken(token string) (*dtos.UserDTO, *error)
}
