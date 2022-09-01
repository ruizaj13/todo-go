package ui_test

import (
	"fmt"
	"testing"
	"net/http"
	"net/http/httptest"

	"github.com/ruizajtruss/todo-app/backend/types"
	"github.com/ruizajtruss/todo-app/backend/ui"
)



type MockService struct {
	err error
	todos []types.Todo
}

func (s MockService) GetAllTodos() ([]types.Todo, error) {
	if s.err != nil {
		return nil, s.err
	}

	return s.todos, nil
}

func TestHTTP(t *testing.T) {
	service := &MockService{
		err: fmt.Errorf("Something bad happened"),
	}

	w:= httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "http://mywebsite.com/", nil)

	server := ui.NewHTTP()

	server.UseService(service)

	server.ServeHTTP(w, r)
}