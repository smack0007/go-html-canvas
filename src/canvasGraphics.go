package main

import (
	"fmt"
	"syscall/js"
)

type CanvasGraphicsContext struct {
	context js.Value;
}

func InitCanvasGraphics() *CanvasGraphicsContext {
	document := js.Global().Get("document");
	canvas := document.Call("getElementById", "host");
	context := canvas.Call("getContext", "2d");

	context.Set("fillStyle", "rgb(0, 0, 0)");

	return &CanvasGraphicsContext { context: context };
}

func (graphics *CanvasGraphicsContext) FillRect(r Rect, c Color) {
	graphics.context.Set("fillStyle", fmt.Sprintf("rgba(%d, %d, %d, %d)", c.r, c.g, c.b, c.a));
	graphics.context.Call("fillRect", r.x, r.y, r.width, r.height);
}
