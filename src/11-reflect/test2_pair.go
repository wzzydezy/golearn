package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	//tty:pair<type:*os.File, "/dev/tty"文件描述符>
	tty, err := os.OpenFile("11-refflect/dev/tty", os.O_RDWR, 0)
	if err!= nil{
		fmt.Println("open file error", err)
		return
	}
    // r : pair<typr: , value:>
	var r io.Reader
	// r : tty:pair<type:*os.File, "/dev/tty"文件描述符>
	r = tty
    // w : pair<typr: , value:>
	var w io.Writer
	// w : tty:pair<type:*os.File, "/dev/tty"文件描述符>
	w = r.(io.Writer) //转换类型

	w.Write([]byte("HELLO TEST!!!"))
}