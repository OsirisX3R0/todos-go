package todos

import "fmt"

type Todo struct {
	// The text of a todo
	text string
	// The status of a todo
	completed TodoStatus
}

// Create a new todo
func NewTodo(text string) Todo {
	return Todo{
		text:      text,
		completed: Incomplete,
	}
}

// Displays the text of a todo
func (t Todo) Text() string {
	return t.text
}

// Returns whether or not a todo is completed
func (t Todo) IsCompleted() bool {
	if t.completed == Complete {
		return true
	} else {
		return false
	}
}

// Updates the text of a todo
func (t *Todo) UpdateText(text string) {
	t.text = text
}

// Sets a todos status to complete
func (t *Todo) Complete() {
	t.completed = Complete
}

// Sets a todos status to incomplete
func (t *Todo) Incomplete() {
	t.completed = Incomplete
}

// Translates a todo to a string
func (t Todo) String() string {
	return fmt.Sprintf("%v %v", t.completed, t.text)
}
