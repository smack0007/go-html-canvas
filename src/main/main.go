package main

import (
	"fmt"
	"time"

	"game"
	"types"
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
	graphics.FillRect(types.Rect{X: 0, Y: 0, Width: 1024, Height: 768}, Color{0, 0, 0, 255})
	graphics.FillRect(types.Rect{X: 10, Y: 10, Width: 100, Height: 100}, Color{255, 0, 0, 255})
	graphics.FillRect(types.Rect{X: 110, Y: 110, Width: 100, Height: 100}, Color{0, 0, 255, 255})

	rect := game.CardImageRectangles[game.Card{Suit: game.Hearts, Value: game.Three}]
	graphics.DrawSprite(image, rect, types.Rect{X: 0, Y: 0, Width: rect.Width, Height: rect.Height})
}
