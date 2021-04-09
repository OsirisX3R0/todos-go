package todos

type Todo struct {
	text      string
	completed TodoStatus
}

func NewTodo(text string) Todo {
	return Todo{
		text:      text,
		completed: Incomplete,
	}
}

func (t Todo) Text(text string) string {
	return t.text
}

func (t Todo) IsCompleted() bool {
	if t.completed == Complete {
		return true
	} else {
		return false
	}
}

func (t *Todo) UpdateText(text string) {
	t.text = text
}

func (t *Todo) Complete() {
	t.completed = Complete
}

func (t *Todo) Incomplete() {
	t.completed = Incomplete
}
