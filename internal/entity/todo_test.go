package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTodo(t *testing.T) {
	todo, err := NewTodo("Test Todo", "This is a test description")
	assert.Nil(t, err)
	assert.Equal(t, todo.Title, "Test Todo")
	assert.Equal(t, todo.Description, "This is a test description")
}

func TestToogleDone(t *testing.T) {
	todo, err := NewTodo("Test Todo", "This is a test description")
	assert.Nil(t, err)
	assert.Equal(t, todo.Done, false)
	todo.ToogleDone()
	assert.Equal(t, todo.Done, true)
}

func TestIsValid(t *testing.T) {
	todo, err := NewTodo("Test Todo", "This is a test description")
	assert.Nil(t, err)
	assert.Equal(t, todo.isValid(), true)
}

func TestNotIsValid(t *testing.T) {
	todo, err := NewTodo("", "")
	assert.Error(t, err)
	assert.Nil(t, todo)
}
