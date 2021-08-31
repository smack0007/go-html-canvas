package main

import (
	types "app/types"
)

type GraphicsContext interface {
	DrawSprite(image Image, source types.Rect, destination types.Rect)
	
	FillRect(r types.Rect, c Color)

	LoadImage(fileName string) Image
}

type Image interface {
	Loaded() bool

	Width() int

	Height() int
}
