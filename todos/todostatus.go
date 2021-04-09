package todos

type TodoStatus int

const (
	Incomplete TodoStatus = iota
	Complete
)

func (ts TodoStatus) String() string {
	switch ts {
	case Complete:
		return "true"
	case Incomplete:
		return "false"
	default:
		return ""
	}
}
