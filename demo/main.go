package main

import (
	"fmt"

	"github.com/jinpain/go-canvas/pkg/canvas"
	"github.com/jinpain/go-canvas/pkg/shape"
)

func main() {
	canvas, err := canvas.New(true)
	if err != nil {
		fmt.Println(err)
		return
	}

	canvas.Resize(800, 600)

	rect := &shape.Rectangle{X: 10, Y: 10, Width: 100, Height: 50, Color: "red"}
	circle := &shape.Circle{X: 200, Y: 200, Radius: 50, Color: "blue"}

	canvas.DrawShape(rect)
	canvas.DrawShape(circle)
}
