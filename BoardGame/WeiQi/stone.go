package weiqi

const (
	S_NONE  = 0
	S_BLACK = 1
	S_WHITE = 2
	S_MAX   = 3
)

func StoneOther(s int) int {
	switch s {
	case S_BLACK:
		return S_WHITE
	case S_WHITE:
		return S_BLACK
	default:
		return s
	}
}
