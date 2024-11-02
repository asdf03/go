package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/asdf03/go/internal/domain"
)

type mockTodoUsecase struct {
	todos []domain.Todo
}

func (m *mockTodoUsecase) GetAll() ([]domain.Todo, error) {
	return m.todos, nil
}

func (m *mockTodoUsecase) Create(todo *domain.Todo) error {
	todo.ID = len(m.todos) + 1
	m.todos = append(m.todos, *todo)
	return nil
}

func TestGetTodos(t *testing.T) {
	mockUsecase := &mockTodoUsecase{
		todos: []domain.Todo{
			{ID: 1, Title: "Test Todo", Completed: false},
		},
	}

	handler := NewTodoHandler(mockUsecase)
	// Todo
}
