package main

type GraphicsContext interface {
	FillRect(r Rect, c Color);
	
	LoadImage(fileName string) Image;
}

type Image interface {
	Width() int32;

	Height() int32;
}