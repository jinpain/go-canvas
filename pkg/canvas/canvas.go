package canvas

import (
	"fmt"
	"syscall/js"

	"github.com/jinpain/go-canvas/pkg/shape"
)

type CanvasInterface interface {
	Clear()
	Resize(width, height int)
	GetContext() js.Value
	GetWidth() int
	GetHeight() int
	DrawShape(shape shape.Shape)
}

type canvas struct {
	window   js.Value
	document js.Value
	body     js.Value
	canvas   js.Value
	ctx      js.Value
	width    int
	height   int
}

func New(isCreate bool) (CanvasInterface, error) {
	cvs := &canvas{
		window:   js.Global(),
		document: js.Global().Get("document"),
		body:     js.Global().Get("document").Get("body"),
	}

	if isCreate {
		canvasElem := cvs.document.Call("createElement", "canvas")
		canvasElem.Set("id", "go-canvas")
		cvs.body.Call("appendChild", canvasElem)
		ctx := canvasElem.Call("getContext", "2d")
		if ctx.IsUndefined() {
			return nil, fmt.Errorf("failed to get 2D context for the canvas")
		}
		cvs.canvas = canvasElem
		cvs.ctx = ctx
	}
	return cvs, nil
}

func GetById(id string) (CanvasInterface, error) {
	cvs := &canvas{
		window:   js.Global(),
		document: js.Global().Get("document"),
		body:     js.Global().Get("document").Get("body"),
	}

	canvasElem := cvs.document.Call("getElementById", id)
	if canvasElem.IsUndefined() {
		return nil, fmt.Errorf("element with id '%s' not found", id)
	}
	ctx := canvasElem.Call("getContext", "2d")
	if ctx.IsUndefined() {
		return nil, fmt.Errorf("failed to get 2D context for the canvas with id '%s'", id)
	}
	cvs.canvas = canvasElem
	cvs.ctx = ctx

	return cvs, nil
}

func (c *canvas) Resize(width, height int) {
	c.width = width
	c.height = height
	c.canvas.Set("width", width)
	c.canvas.Set("height", height)
}

func (c *canvas) Clear() {
	c.ctx.Call("clearRect", 0, 0, c.width, c.height)
}

func (c *canvas) GetContext() js.Value {
	return c.ctx
}

func (c *canvas) GetWidth() int {
	return c.width
}

func (c *canvas) GetHeight() int {
	return c.height
}

func (c *canvas) DrawShape(shape shape.Shape) {
	shape.Draw(c.ctx)
}
