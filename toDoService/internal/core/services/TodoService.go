package services

import (
	"errors"
	"github.com/google/uuid"
	"time"
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

func (s *TodoService) authorizeUser(token string) (*dtos.UserDTO, *error) {
	userCredantials, err := s.comHandler.AuthorizeUserByToken(token)
	if err != nil {
		return nil, err
	}
	return userCredantials, nil
}

func (s *TodoService) GetTodoByUserId(request *Requests.GetUserTodosRequest) (*[]domain.Todo, *error) {

	userCredantials, err := s.authorizeUser(request.Token)
	if err != nil {
		return nil, err
	}
	parsedId, parseErr := uuid.Parse(userCredantials.ID)

	if parseErr != nil {
		return nil, &parseErr
	}

	res, err := s.db.FindTodosByUserId(parsedId)
	if err != nil {
		return nil, err
	}
	return res, nil

}
func (s *TodoService) UpdateTodo(request *Requests.UpdateTodoRequest) *error {
	userCredantials, err := s.authorizeUser(request.Token)
	if err != nil {
		return err
	}
	parsedId, parseErr := uuid.Parse(userCredantials.ID)

	if parseErr != nil {
		return &parseErr
	}

	existingTodo, err := s.db.FindTodoById(request.TodoId)
	if err != nil {
		return err
	}

	if existingTodo.UserId != parsedId {
		newErr := errors.New("Unauthorized Request!") // TODO make these same errors as a new file
		return &newErr
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
func (s *TodoService) CreateTodo(request *Requests.CreateTodoRequest) *error {
	userCredantials, err := s.authorizeUser(request.Token)
	if err != nil {
		return err
	}
	parsedId, parseErr := uuid.Parse(userCredantials.ID)

	if parseErr != nil {
		return &parseErr
	}
	newTodo := &domain.Todo{
		ID:             uuid.UUID{},
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
func (s *TodoService) DeleteTodo(request *Requests.DeleteTodoRequest) *error {
	userCredantials, err := s.authorizeUser(request.Token)
	if err != nil {
		return err
	}
	parsedUserId, parseErr := uuid.Parse(userCredantials.ID)

	if parseErr != nil {
		return &parseErr
	}

	existingToDo, err := s.db.FindTodoById(request.TodoId)

	if err != nil {
		return err
	}
	if existingToDo.UserId != parsedUserId {
		newErr := errors.New("Unauthorized Request")
		return &newErr
	}
	err = s.db.DeleteTodo(existingToDo)
	return err
}
