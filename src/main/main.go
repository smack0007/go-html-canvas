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
		fmt.Println(fmt.Sprintf("Image Size: %d %d", image.Width(), image.Height()))

		fmt.Println(game.ImageOffsets[0]);
		draw(graphics, image)
	})()

	select {}
}

func draw(graphics GraphicsContext, image Image) {
	graphics.FillRect(types.Rect{X: 0, Y: 0, Width: 1024, Height: 768}, Color{0, 0, 0, 255})
	graphics.FillRect(types.Rect{X: 10, Y: 10, Width: 100, Height: 100}, Color{255, 0, 0, 255})
	graphics.FillRect(types.Rect{X: 110, Y: 110, Width: 100, Height: 100}, Color{0, 0, 255, 255})
	graphics.DrawSprite(image, game.ImageOffsets[0], types.Rect{X: 0, Y: 0, Width: game.ImageOffsets[0].Width, Height: game.ImageOffsets[0].Height})
}
