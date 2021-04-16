package todos

import (
	"testing"
)

func TestNewTodoList(t *testing.T) {
	equal := true
	defaultList := NewTodoList()
	newList := TodoList{}

	if ((defaultList == nil) != (newList == nil)) || defaultList.Len() != newList.Len() {
		equal = false
	}

	if equal {
		for i := range defaultList {
			if defaultList[i] != newList[i] {
				equal = false
			}
		}

	}
	if equal != true {
		t.Errorf("Should be equal")
	}
}
