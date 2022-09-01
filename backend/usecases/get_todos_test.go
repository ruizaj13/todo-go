package usecases_test

import (
	"fmt"
	"testing"

	"github.com/ruizajtruss/todo-app/backend/types"
	"github.com/ruizajtruss/todo-app/backend/usecases"
)

type MockTodosRepo struct {

}

func (MockTodosRepo) GetAllTodos() ([]types.Todo, error) {
	return nil, fmt.Errorf("Something went wrong")
}

func TestGetTodos(t *testing.T) {
	repo := new(MockTodosRepo)

	todos, err := usecases.GetTodos(repo)

	if err != usecases.ErrInternal {
		t.Fatalf("Expected ErrInternal; Got: %v", err)
	}

	if todos != nil {
		t.Fatalf("Expected todos to be nil; Got: %v", todos)
	}
}