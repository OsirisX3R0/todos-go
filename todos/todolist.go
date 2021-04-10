package todos

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

// Adds a new todo
func (tl *TodoList) Add(text string) Todos {
	var todo = NewTodo(text)
	return append(tl.todos, todo)
}

// Updates a todos text by index
func (tl *TodoList) UpdateText(index int, text string) {
	var todo = tl.todos[index]
	todo.UpdateText(text)
}

// Set a todos status to complete by index
func (tl *TodoList) Complete(index int) {
	var todo = tl.todos[index]
	todo.Complete()
}

// Set a todos status to incomplete by index
func (tl *TodoList) Incomplete(index int) {
	var todo = tl.todos[index]
	todo.Incomplete()
}

// Delete a todo
func (tl *TodoList) Delete(index int) []Todo {
	return append(tl.todos[:index], tl.todos[index+1])
}

// Clear all todos
func (tl *TodoList) Clear() {
	tl.todos = NewTodos()
}

// Translates a todo list to strings
