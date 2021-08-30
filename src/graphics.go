package main

type GraphicsContext interface {
	FillRect(r Rect, c Color)

	LoadImage(fileName string) Image
}

type Image interface {
	Loaded() bool

	Width() int32

	Height() int32
}