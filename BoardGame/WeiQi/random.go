package weiqi

import "math/rand"

func RandomGenMove(b *Board, color int) *Coordinate {
	m := Move{Color: color}
	if b.BoardNoValidMoves(color) {
		return &Coordinate{-1, -1}
	}
	for !b.BoardValidMove(&m, 1) {
		m.Coord.X = rand.Intn(b.Size)
		m.Coord.Y = rand.Intn(b.Size)
	}
	return &Coordinate{m.Coord.X, m.Coord.Y}
}

func EngineRandomInit() (engine *Engine) {
	return &Engine{Name: "RandomMove Engine", Comment: "I just make random moves. I won't pass as long as there is a place on the board where I can play. When we both pass, I will consider all the stones on the board alive.", Genmove: RandomGenMove}
}
