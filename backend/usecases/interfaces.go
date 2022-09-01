package usecases

import "github.com/ruizajtruss/todo-app/backend/types"

type TodosRepository interface {
	GetAllTodos() ([]types.Todo, error)
}