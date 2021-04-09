package todos

type TodoList struct {
	todos []Todo
}

func NewTodoList() TodoList {
	var todos []Todo

	return TodoList{
		todos: todos,
	}
}

func (tl TodoList) Get(index int) Todo {
	return tl.todos[index]
}

func (tl *TodoList) Add(text string) []Todo {
	var todo = NewTodo(text)
	return append(tl.todos, todo)
}

func (tl *TodoList) UpdateText(index int, text string) {
	var todo = tl.todos[index]
	todo.UpdateText(text)
}

func (tl *TodoList) Complete(index int) {
	var todo = tl.todos[index]
	todo.Complete()
}

func (tl *TodoList) Incomplete(index int) {
	var todo = tl.todos[index]
	todo.Incomplete()
}

func (tl *TodoList) Delete(index int) []Todo {
	return append(tl.todos[:index], tl.todos[index+1])
}

func (tl *TodoList) Clear() {
	var todos []Todo
	tl.todos = todos
}
