package todos

// An array of todos
type Todos []Todo

// Creates a new array of todos
func NewTodos() Todos {
	var todos Todos
	return todos
}

func (t *Todos) Clear() Todos {
	return NewTodos()
}
