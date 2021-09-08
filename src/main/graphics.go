package main

type Rectangle struct {
	X, Y, Width, Height int
}

type GraphicsContext interface {
	DrawSprite(image Image, source Rectangle, destination Rectangle)
	
	FillRect(r Rectangle, c Color)

	LoadImage(fileName string) Image
}

type Image interface {
	Loaded() bool

	Width() int

	Height() int
}
