package main

import (
	"fmt"

	"github.com/OsirisX3R0/todos-go/todos"
)

func main() {
	var todos = todos.NewTodoList()
	todos.Add("Take out the trash")
	todos.Add("Cook the dinner")
	todos.Add("Wash the dishes")
	todos.Add("Go to bed")

	index := 1
	fmt.Println("Initial Todos: ", todos)
	fmt.Println("\n\nTodo at index ", index, ": ", todos.Get(1))

	todos.Complete(0)
	todos.Complete(1)
	todos.Incomplete(1)
	todos.Delete(3)
	todos.UpdateText(1, "Grill the burgers")

	var todo = todos.Get(0)
	fmt.Println("\n\nUpdated Todos: ", todos)
	fmt.Println("\n\nNew todo at index ", index, ": ", todo)

	todos.Clear()
	fmt.Println("\n\nCleared todos: ", todos)
}
