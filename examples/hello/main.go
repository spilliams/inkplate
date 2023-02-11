package main

import (
	"fmt"
	"os"

	"github.com/spilliams/inkplate/pkg/inkplate"
)

func main() {
	i, err := inkplate.New("/dev/ttyUSB0")
	defer i.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ok, err := i.IsOK()
	if err != nil {
		fmt.Println(err)
	}
	if ok {
		fmt.Println("Is OK")
	} else {
		fmt.Println("Is NOT OK")
	}
}
