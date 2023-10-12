package services

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
	"toDoService/internal/core/customErrors"
	"toDoService/internal/core/domain"
	"toDoService/internal/core/dtos"
	"toDoService/internal/core/dtos/Requests"
	"toDoService/internal/core/ports"
)

type TodoService struct {
	db         ports.ToDoRepository
	comHandler ports.CoomunactionHandler
}

func NewTodoService(db ports.ToDoRepository, comHandler ports.CoomunactionHandler) *TodoService {
	return &TodoService{db: db, comHandler: comHandler}
}

func (s *TodoService) authorizeUser(token string) (*dtos.UserDTO, error) {
	userCredantials, err := s.comHandler.AuthorizeUserByToken(token)
	if err != nil {
		return nil, err
	}
	return userCredantials, nil
}

func (s *TodoService) GetTodoByUserId(token string) (*[]domain.Todo, error) { // should change the response type

	userCredantials, err := s.authorizeUser(token)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	parsedId, err := uuid.Parse(userCredantials.ID)

	if err != nil {
		return nil, err
	}

	res, err := s.db.FindTodosByUserId(parsedId)
	if err != nil {
		return nil, err
	}
	return res, nil

}
func (s *TodoService) UpdateTodo(request *Requests.UpdateTodoRequest, token string) error {
	userCredantials, err := s.authorizeUser(token)
	if err != nil {
		return err
	}
	parsedId, parseErr := uuid.Parse(userCredantials.ID)

	if parseErr != nil {
		return parseErr
	}

	requestTodoId, err := uuid.Parse(request.TodoId)
	if err != nil {
		return err
	}
	existingTodo, err := s.db.FindTodoById(requestTodoId)
	if err != nil {
		return err
	}

	if existingTodo.UserId != parsedId {
		var newErr error = &customErrors.UnauthorizedError{} // TODO make these same customErrors as a new file
		return newErr
	}

	existingTodo.Title = request.Title
	existingTodo.Description = request.Description
	if !existingTodo.IsDone && request.IsDone {
		existingTodo.IsDone = true
		nowTime := time.Now()
		existingTodo.CompletionDate = &nowTime // is this wrong to do ?
	}

	err = s.db.UpdateTodo(existingTodo)
	return err

}
func (s *TodoService) CreateTodo(request *Requests.CreateTodoRequest, token string) error {
	userCredantials, err := s.authorizeUser(token)
	if err != nil {
		return err
	}
	parsedId, parseErr := uuid.Parse(userCredantials.ID)

	if parseErr != nil {
		return parseErr
	}
	newTodo := &domain.Todo{
		ID:             uuid.New(),
		UserId:         parsedId,
		Title:          request.Title,
		Description:    request.Description,
		CreateDate:     time.Now(),
		IsDone:         false,
		CompletionDate: nil,
	}
	err = s.db.CreateTodo(newTodo)
	return err

}
func (s *TodoService) DeleteTodo(todoId string, token string) error {
	userCredantials, err := s.authorizeUser(token)
	if err != nil {
		return err
	}
	parsedUserId, parseErr := uuid.Parse(userCredantials.ID)

	if parseErr != nil {
		return parseErr
	}

	parsedTodoId, parseErr := uuid.Parse(todoId)
	if parseErr != nil {
		return parseErr
	}

	existingToDo, err := s.db.FindTodoById(parsedTodoId)

	if err != nil {
		return err
	}
	if existingToDo.UserId != parsedUserId {
		newErr := errors.New("Unauthorized Request")
		return newErr
	}
	err = s.db.DeleteTodo(existingToDo)
	return err
}
