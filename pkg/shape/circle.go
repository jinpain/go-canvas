package shape

import "syscall/js"

type Circle struct {
	X      int
	Y      int
	Radius int
	Color  string
}

func (c *Circle) Draw(ctx js.Value) {
	ctx.Set("fillStyle", c.Color)
	ctx.Call("beginPath")
	ctx.Call("arc", c.X, c.Y, c.Radius, 0, 2*3.14)
	ctx.Call("fill")
}
