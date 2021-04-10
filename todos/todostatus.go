package todos

// Enum representing the status of a todo
type TodoStatus int

const (
	// Represents an incomplete todo
	Incomplete TodoStatus = iota
	// Represents a complete todo
	Complete
)

// Translates a TodoStatus to a string
func (ts TodoStatus) String() string {
	switch ts {
	case Complete:
		return "x"
	case Incomplete:
		return " "
	default:
		return ""
	}
}
