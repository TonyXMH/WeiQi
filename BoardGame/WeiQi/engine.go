package weiqi

type Engine struct {
	Name    string
	Comment string
	Genmove func(b *Board, color int) (coord *Coordinate)
}
