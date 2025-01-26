package shape

import "syscall/js"

type Shape interface {
	Draw(ctx js.Value)
}
