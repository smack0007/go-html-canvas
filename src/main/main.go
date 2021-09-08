package main

import (
	"fmt"
	"time"

	"cards"
)

func main() {
	graphics := InitCanvasGraphics()
	fmt.Println("Loading image...")
	image := graphics.LoadImage("cards.png")

	go (func() {
		for !image.Loaded() {
			time.Sleep(50)
		}

		fmt.Println("Images loaded.")
		
		draw(graphics, image)
	})()

	select {}
}

func draw(graphics GraphicsContext, image Image) {
	graphics.FillRect(Rectangle{X: 0, Y: 0, Width: 1024, Height: 768}, Color{0, 0, 0, 255})
	graphics.FillRect(Rectangle{X: 10, Y: 10, Width: 100, Height: 100}, Color{255, 0, 0, 255})
	graphics.FillRect(Rectangle{X: 110, Y: 110, Width: 100, Height: 100}, Color{0, 0, 255, 255})

	rect := cards.CardImageRectangles[cards.Card{Suit: cards.Hearts, Value: cards.Three}]
	graphics.DrawSprite(image, Rectangle(rect), Rectangle{X: 0, Y: 0, Width: rect.Width, Height: rect.Height})
}
