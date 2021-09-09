package main

import (
	"fmt"
	"time"

	"cards"
	"engine"
	"engine/graphics"
)

func main() {
	graphics := graphics.InitCanvasGraphics()
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

func draw(graphics graphics.GraphicsContext, image graphics.Image) {
	graphics.FillRect(engine.Rectangle{X: 0, Y: 0, Width: 1024, Height: 768}, engine.Color{R: 0, G: 0, B: 0, A: 255})
	graphics.FillRect(engine.Rectangle{X: 10, Y: 10, Width: 100, Height: 100}, engine.Color{R: 255, G: 0, B: 0, A: 255})
	graphics.FillRect(engine.Rectangle{X: 110, Y: 110, Width: 100, Height: 100}, engine.Color{R: 0, G: 0, B: 255, A: 255})

	rect := cards.CardImageRectangles[cards.Card{Suit: cards.Hearts, Value: cards.Three}]
	graphics.DrawSprite(image, engine.Rectangle(rect), engine.Rectangle{X: 0, Y: 0, Width: rect.Width, Height: rect.Height})
}
