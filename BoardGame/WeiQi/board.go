package weiqi

import (
	"fmt"
)

type Board struct {
	Size            int
	Captures        [S_MAX]int
	Komi            float64
	StepCnt         int
	B               []int //棋盘位置对应颜色，大小size*size
	G               []int //区域同色且直连的棋子具有相同的gid，大小size*size
	LastGid         int
	LibCntWaterMark []bool
}

//Tool function

func BoardAtXY(b *Board, x, y int) int {
	return b.B[x+b.Size*y]
}

func BoardAt(b *Board, c *Coordinate) int {
	return BoardAtXY(b, c.X, c.Y)
}

func GroupAtXY(b *Board, x, y int) int {
	return b.G[x+b.Size*y]
}

func GroupAt(b *Board, c *Coordinate) int {
	return GroupAtXY(b, c.X, c.Y)
}

func XY2Pos(b *Board, x, y int) int {
	return x + b.Size*y
}

func Pos2X(b *Board, pos int) int {
	return pos % b.Size
}

func Pos2Y(b *Board, pos int) int {
	return pos / b.Size
}

func Coord2Pos(b *Board, c *Coordinate) int {
	return XY2Pos(b, c.X, c.Y)
}

// Board Function

func (b *Board) Reszie(size int) {
	b.Size = size
	b.B = make([]int, b.Size*b.Size)
	b.G = make([]int, b.Size*b.Size)
}
func BoardCopy(src *Board) (dest *Board) {
	dest = &Board{}
	dest.B = src.B
	dest.G = src.G
	return
}

//func (b *Board)Print(f *os.File)  {
//
//}
func (b *Board) Print() {
	fmt.Printf("Move: % 3d  Komi: %2.1f  Captures B: %d W: %d\n     ", b.StepCnt, b.Komi, b.Captures[S_BLACK], b.Captures[S_WHITE])
	char := string("ABCDEFGHJKLMNOPQRSTUVWXYZ")
	for x := 0; x < b.Size; x++ {
		fmt.Printf("%c", char[x])
	}
	fmt.Printf("\n   +-")
	for x := 0; x < b.Size; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n")
	for y := b.Size - 1; y >= 0; y-- {
		fmt.Printf("%2d | ", y+1)
		for x := 0; x < b.Size; x++ {
			fmt.Printf("%d ", BoardAtXY(b, x, y))
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("   +-")
	for x := 0; x < b.Size; x++ {
		fmt.Printf("--")
	}
	fmt.Printf("+\n\n")

}

func (b *Board) BoardPlayNoCheck(m *Move) int {
	gid := 0
	pos := Coord2Pos(b, &m.Coord)
	b.B[pos] = m.Color

	coords := [4]Coordinate{Coordinate{m.Coord.X - 1, m.Coord.Y},
		Coordinate{m.Coord.X, m.Coord.Y - 1},
		Coordinate{m.Coord.X + 1, m.Coord.Y},
		Coordinate{m.Coord.X, m.Coord.Y + 1}}
	for i := 0; i < 4; i++ {
		c := Coordinate{coords[i].X, coords[i].Y}
		if c.X < 0 || c.Y < 0 || c.X >= b.Size || c.Y >= b.Size {
			continue
		}
		if BoardAt(b, &c) == m.Color && GroupAt(b, &c) != gid {
			if gid <= 0 {
				gid = GroupAt(b, &c)
			} else { //合并区域类似连接对角下面的点
				for x := 0; x < b.Size; x++ {
					for y := 0; y < b.Size; y++ {
						if GroupAtXY(b, x, y) == GroupAt(b, &c) {
							b.G[XY2Pos(b, x, y)] = gid
						}
					}
				}
			}
		} else if BoardAt(b, &c) == StoneOther(m.Color) && b.BoardGroupLibs(GroupAt(b, &c)) == 0 { //提子
			b.BoardGroupCapture(GroupAt(b, &c))
		}
	}
	if gid <= 0 {
		b.LastGid++
		gid = b.LastGid
	}
	b.G[pos] = gid
	b.StepCnt++
	return gid
}
func (b *Board) BoardPlay(m *Move) int {
	if !b.BoardValidMove(m, 0) {
		return 0
	}
	return b.BoardPlayNoCheck(m)
}

func (b *Board) BoardNoValidMoves(color int) bool {
	for x := 0; x < b.Size; x++ {
		for y := 0; y < b.Size; y++ {
			c := Coordinate{x, y}
			m := Move{Coord: c, Color: color}
			if b.BoardValidMove(&m, 1) {
				return false
			}
		}
	}
	return true
}

func (b *Board) BoardValidMove(m *Move, sensible int) bool {

	if BoardAt(b, &m.Coord) != S_NONE {
		return false
	}
	tmpB := BoardCopy(b)
	tmpB.BoardPlayNoCheck(m)
	if tmpB.BoardGroupLibs(GroupAt(tmpB, &m.Coord)) <= sensible {
		return false
	}
	return true
}

func (b *Board) BoardLocalLibs(coord *Coordinate) int {
	l := 0
	coords := [4]Coordinate{Coordinate{coord.X - 1, coord.Y},
		Coordinate{coord.X, coord.Y - 1},
		Coordinate{coord.X + 1, coord.Y},
		Coordinate{coord.X, coord.Y + 1}}
	for i := 0; i < 4; i++ {
		c := Coordinate{coords[i].X, coords[i].Y}
		if c.X < 0 || c.Y < 0 || c.X >= b.Size || c.Y >= b.Size {
			continue
		}
		if b.LibCntWaterMark != nil {
			if b.LibCntWaterMark[XY2Pos(b, c.X, c.Y)] {
				continue
			}
			b.LibCntWaterMark[XY2Pos(b, c.X, c.Y)] = true
		}
		if BoardAt(b, &c) == S_NONE {
			l++
		}

	}
	return l

}

func (b *Board) BoardGroupLibs(gid int) int {
	l := 0
	b.LibCntWaterMark = make([]bool, b.Size*b.Size)
	for x := 0; x < b.Size; x++ {
		for y := 0; y < b.Size; y++ {
			c := Coordinate{x, y}
			if GroupAtXY(b, x, y) == gid {
				l += b.BoardLocalLibs(&c)
			}
		}
	}
	b.LibCntWaterMark = nil
	return l
}

func (b *Board) BoardGroupCapture(gid int) {
	for x := 0; x < b.Size; x++ {
		for y := 0; y < b.Size; y++ {
			c := Coordinate{x, y}
			if GroupAtXY(b, x, y) == gid {
				b.Captures[StoneOther(BoardAt(b, &c))]++
				b.B[XY2Pos(b, x, y)] = S_NONE
				b.G[XY2Pos(b, x, y)] = 0
			}
		}
	}
}
