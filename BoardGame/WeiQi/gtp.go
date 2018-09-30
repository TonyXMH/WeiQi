package weiqi

import (
	"strings"
	"strconv"
	"fmt"
)

func Str2Coord(str string)(Coordinate)  {
	return Coordinate{X:int(str[0]-'A'),Y:int(str[1]-'1')}
}
func GtpParse(b *Board, engine *Engine, str string) {//这个函数很不安全，但是是测试用，则先只做一些简单的判断
	cmds:=strings.Split(str," ")
	if cmds[0] == "boardsize"{
		size,_ := strconv.Atoi(cmds[1])
		b.Reszie(size)
	}else if cmds[0] == "komi"{
		komi,_:=strconv.ParseFloat(cmds[1],64)
		b.Komi = komi
	}else if cmds[0] == "play"{
		m:=Move{}
		if cmds[1] == "black"{
			m.Color = S_BLACK
		}else if cmds[1] == "white"{
			m.Color = S_WHITE
		}else{
			m.Color = S_NONE
		}
		m.Coord = Str2Coord(cmds[2])
		if b.BoardPlay(&m) == 0{
			fmt.Printf( "! ILLEGAL MOVE %d,%d,%d\n", m.Color, m.Coord.X, m.Coord.Y)
		}

	}else if cmds[0] == "genmove"{
		m:=Move{}
		if cmds[1] == "black"{
			m.Color = S_BLACK
		}else if cmds[1] == "white"{
			m.Color = S_WHITE
		}else{
			m.Color = S_NONE
		}
		c:=engine.Genmove(b,m.Color)
		m.Coord = *c
		b.BoardPlay(&m)
	}else{
		fmt.Println("unknown command")
	}
}
