package main

import (
	"fmt"
	"syscall/js"
)

type CanvasGraphicsContext struct {
	document js.Value;
	canvas js.Value;
	context js.Value;
}

type CanvasImage struct {
	image js.Value;
}

func InitCanvasGraphics() *CanvasGraphicsContext {
	document := js.Global().Get("document");
	canvas := document.Call("getElementById", "host");
	context := canvas.Call("getContext", "2d");

	context.Set("fillStyle", "rgb(0, 0, 0)");

	return &CanvasGraphicsContext {
		document: document,
		canvas: canvas,
		context: context,
	};
}

func (graphics *CanvasGraphicsContext) FillRect(r Rect, c Color) {
	graphics.context.Set("fillStyle", fmt.Sprintf("rgba(%d, %d, %d, %d)", c.r, c.g, c.b, c.a));
	graphics.context.Call("fillRect", r.x, r.y, r.width, r.height);
}

func (graphics *CanvasGraphicsContext) LoadImage(fileName string) Image {
	image := js.Global().Get("Image").New();
	image.Set("src", fileName);
	return &CanvasImage{ image: image };
}

func (image *CanvasImage) Width() int32 {
	return 0;	
}

func (image *CanvasImage) Height() int32 {
	return 0;	
}