package main

import (
	"fmt"
	"time"

	game "app/game"
	types "app/types"
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
	graphics.FillRect(types.Rect{0, 0, 1024, 768}, Color{0, 0, 0, 255})
	graphics.FillRect(types.Rect{10, 10, 100, 100}, Color{255, 0, 0, 255})
	graphics.FillRect(types.Rect{110, 110, 100, 100}, Color{0, 0, 255, 255})
	graphics.DrawSprite(image, game.ImageOffsets[0], types.Rect{0, 0, game.ImageOffsets[0].Width, game.ImageOffsets[0].Height})
}
