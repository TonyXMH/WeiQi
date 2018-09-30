package main

import (
	"./WeiQi"
	"bufio"
	"os"
	"io"
)

func main() {
	b := &weiqi.Board{}
	e := weiqi.EngineRandomInit()
	inputReader := bufio.NewReader(os.Stdin)
	for true  {
		str,err := inputReader.ReadString('\n')
		if err == io.EOF{
			break
		}
		weiqi.GtpParse(b,e,str)
		b.Print()
	}

}
