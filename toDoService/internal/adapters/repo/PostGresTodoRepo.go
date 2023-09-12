package repo

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"toDoService/internal/core/domain"
)

type TodoPostgresRepository struct {
	db *gorm.DB
}

func NewTodoPostgresRepository() *TodoPostgresRepository {
	host := "localhost"
	port := "5431"
	user := "postgres"
	password := "admin"
	dbname := "todo"

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password,
	)

	db, err := gorm.Open(postgres.Open(conn))
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&domain.Todo{})
	if err != nil {
		panic(err)
	}

	return &TodoPostgresRepository{
		db: db,
	}
}

func (s *TodoPostgresRepository) FindTodosByUserId(id uuid.UUID) (*[]domain.Todo, *error) {
	todos := &[]domain.Todo{}
	req := *s.db.Limit(10).Find(todos, "user_id = ?", id)
	if req.Error != nil {
		err := errors.New("unexpected error")
		return nil, &err // TODO
	}
	return todos, nil

}
func (s *TodoPostgresRepository) CreateTodo(todo *domain.Todo) *error {
	return nil
}
func (s *TodoPostgresRepository) UpdateTodo(todo *domain.Todo) *error {
	return nil
}
func (s *TodoPostgresRepository) DeleteTodo(todo *domain.Todo) *error {
	return nil
}
func (s *TodoPostgresRepository) FindTodoById(todoId uuid.UUID) (*domain.Todo, *error) {
	return nil, nil
}
