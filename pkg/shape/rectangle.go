package shape

import "syscall/js"

type Rectangle struct {
	X      int
	Y      int
	Width  int
	Height int
	Color  string
}

func (r *Rectangle) Draw(ctx js.Value) {
	ctx.Set("fillStyle", r.Color)
	ctx.Call("fillRect", r.X, r.Y, r.Width, r.Height)
}
