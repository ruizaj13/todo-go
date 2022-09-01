package ui

import "github.com/ruizajtruss/todo-app/backend/types"

type Service interface {
	GetAllTodos() ([]types.Todo, error)
}