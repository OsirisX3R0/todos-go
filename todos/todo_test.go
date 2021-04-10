package todos

import (
	"testing"
)

func TestCreateNewTodo(t *testing.T) {
	defaultTodo := NewTodo("New todo")
	newTodo := Todo{
		text:      "New todo",
		completed: Incomplete,
	}

	if defaultTodo != newTodo {
		t.Errorf("Todos not equal")
	}
}

func TestReturnText(t *testing.T) {
	todo := NewTodo("New todo")

	if todo.Text() != "New todo" {
		t.Errorf("Todo text: expected 'New todo'")
	}
}

func TestReturnIsCompleted(t *testing.T) {
	todo := NewTodo("New todo")

	if todo.IsCompleted() != false {
		t.Errorf("Todo completed: extected false")
	}
}

func TestUpdateText(t *testing.T) {
	todo := NewTodo("New todo")
	todo.UpdateText("New text")

	if todo.Text() != "New text" {
		t.Errorf("Todo completed: extected 'New text'")
	}
}

func TestStartAsIncomplete(t *testing.T) {
	todo := NewTodo("New todo")

	if todo.IsCompleted() != false {
		t.Errorf("Todo should initialize to Incomplete")
	}
}

func TestMarkCompete(t *testing.T) {
	todo := NewTodo("New todo")
	todo.Complete()

	if todo.IsCompleted() != true {
		t.Errorf("Todo should be Complete")
	}
}

func TestMarkIncomplete(t *testing.T) {
	todo := NewTodo("New todo")
	todo.Complete()
	todo.Incomplete()

	if todo.IsCompleted() != false {
		t.Errorf("Todo should be Incomplete")
	}
}
