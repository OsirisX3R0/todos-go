package todos

import (
	"reflect"
	"testing"
)

func TestNewTodos(t *testing.T) {
	var slice []Todo
	todos := NewTodos()

	if reflect.DeepEqual(slice, todos) {
		t.Errorf("Should be equal")
	}
}
