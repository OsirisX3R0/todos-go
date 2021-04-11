package todos

import "fmt"

// A list of todos
type TodoList struct {
	todos Todos
}

// Creates a new list of todos
func NewTodoList() TodoList {
	return TodoList{
		todos: NewTodos(),
	}
}

// Gets a todo by index
func (tl TodoList) Get(index int) Todo {
	return tl.todos[index]
}

// Get the length of the todo list
func (tl TodoList) Len() int {
	return len(tl.todos)
}

// Adds a new todo
func (tl *TodoList) Add(text string) {
	todo := NewTodo(text)
	tl.todos = append(tl.todos, todo)
}

// Updates a todos text by index
func (tl *TodoList) UpdateText(index int, text string) {
	tl.todos[index].UpdateText(text)
}

// Set a todos status to complete by index
func (tl *TodoList) Complete(index int) {
	tl.todos[index].Complete()
}

// Set a todos status to incomplete by index
func (tl *TodoList) Incomplete(index int) {
	tl.todos[index].Incomplete()
}

// Delete a todo
func (tl *TodoList) Delete(index int) {
	tl.todos = append(tl.todos[:index], tl.todos[index+1:]...)
	//return tl.todos
}

// Clear all todos
func (tl *TodoList) Clear() {
	tl.todos = NewTodos()
}

// Translates a todo list to strings
func (tl TodoList) String() string {
	if tl.Len() == 0 {
		return "No Todos"
	}

	todoStr := ""
	for _, t := range tl.todos {
		str := fmt.Sprintf("\n%v", t)
		todoStr += str
	}

	return todoStr
}
