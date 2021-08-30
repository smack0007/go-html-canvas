package main

import "fmt"

var signal = make(chan int)

func keepAlive() {
	for {
		<-signal
	}
}

func main() {
	graphics := InitCanvasGraphics()
	fmt.Println("Loading image...");
	image := graphics.LoadImage("cards.png")

	// for !image.Loaded() {
	// 	// fmt.Println("Image not loaded...");
	// }

	draw(graphics)

	keepAlive()
}

func draw(graphics GraphicsContext) {
	graphics.FillRect(Rect{0, 0, 1024, 768}, Color{0, 0, 0, 255})
	graphics.FillRect(Rect{10, 10, 100, 100}, Color{255, 0, 0, 255})
	graphics.FillRect(Rect{110, 110, 100, 100}, Color{0, 0, 255, 255})
}