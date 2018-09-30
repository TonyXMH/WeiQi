package weiqi

type Coordinate struct {
	X int
	Y int
}

type Move struct {
	Coord Coordinate
	Color int
}
