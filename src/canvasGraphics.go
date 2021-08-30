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
	loaded bool;
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
	
	result := CanvasImage{ image: image, loaded: false };

	var loadedCallback js.Func;

	loaded := func(this js.Value, args []js.Value) interface{} {
		result.loaded = true;
		image.Call("removeEventListener", "load", loadedCallback);
		return nil;
	};
	
	loadedCallback = js.FuncOf(loaded);

	image.Call("addEventListener", "load", loadedCallback);
	image.Set("src", fileName);
	return &result;
}

func (image *CanvasImage) Loaded() bool {
	return image.loaded;	
}

func (image *CanvasImage) Width() int32 {
	return 0;	
}

func (image *CanvasImage) Height() int32 {
	return 0;	
}