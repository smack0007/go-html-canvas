package graphics

import (
	"engine"
)

type GraphicsContext interface {
	DrawSprite(image Image, source engine.Rectangle, destination engine.Rectangle)
	
	FillRect(r engine.Rectangle, c engine.Color)

	LoadImage(fileName string) Image
}

type Image interface {
	Loaded() bool

	Width() int

	Height() int
}
