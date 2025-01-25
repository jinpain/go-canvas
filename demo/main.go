package main

import (
	"fmt"

	"github.com/jinpain/go-canvas/pkg/canvas"
)

func main() {
	canvas, err := canvas.New(true)
	if err != nil {
		fmt.Println(err)
		return
	}

	canvas.Resize(800, 600)
}
