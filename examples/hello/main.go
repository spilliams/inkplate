package main

import (
	"fmt"
	"os"

	"github.com/spilliams/inkplate/pkg/inkplate"
)

func main() {
	i, e := inkplate.New("/dev/ttyUSB0")
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	i.Clear()
	i.Print("Hello, world")
}
