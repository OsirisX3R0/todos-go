package todos

import "fmt"

// A list of todos
type TodoList []Todo

// Creates a new list of todos
func NewTodoList() TodoList {
	todolist := TodoList{}
	return todolist
}

// Gets a todo by index
func (tl TodoList) Get(index int) Todo {
	return tl[index]
}

// Get the length of the todo list
func (tl TodoList) Len() int {
	return len(tl)
}

// Adds a new todo
func (tl TodoList) Add(text string) {
	todo := NewTodo(text)
	tl = append(tl, todo)
}

// Updates a todos text by index
func (tl TodoList) UpdateText(index int, text string) {
	tl[index].UpdateText(text)
}

// Set a todos status to complete by index
func (tl TodoList) Complete(index int) {
	tl[index].Complete()
}

// Set a todos status to incomplete by index
func (tl TodoList) Incomplete(index int) {
	tl[index].Incomplete()
}

// Delete a todo
func (tl TodoList) Delete(index int) {
	tl = append(tl[:index], tl[index+1:]...)
}

// Clear all todos
func (tl TodoList) Clear() {
	tl = NewTodoList()
}

// Translates a todo list to strings
func (tl TodoList) String() string {
	if tl.Len() == 0 {
		return "No Todos"
	}

	todoStr := ""
	for _, t := range tl {
		str := fmt.Sprintf("\n%v", t)
		todoStr += str
	}

	return todoStr
}
