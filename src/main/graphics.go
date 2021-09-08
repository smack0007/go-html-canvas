package main

import (
	"types"
)

type GraphicsContext interface {
	DrawSprite(image Image, source types.Rectangle, destination types.Rectangle)
	
	FillRect(r types.Rectangle, c Color)

	LoadImage(fileName string) Image
}

type Image interface {
	Loaded() bool

	Width() int

	Height() int
}
