package todos

import (
	"testing"
)

func TestNewTodoList(t *testing.T) {
	equal := true
	defaultList := NewTodoList()
	newList := TodoList{
		todos: NewTodos(),
	}

	if ((defaultList.todos == nil) != (newList.todos == nil)) || defaultList.Len() != newList.Len() {
		equal = false
	}

	for i := range defaultList.todos {
		if defaultList.todos[i] != newList.todos[i] {
			equal = false
		}
	}

	if equal != true {
		t.Errorf("Should be equal")
	}
}
