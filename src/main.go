package main

func main() {  
	graphics := InitCanvasGraphics();
	draw(graphics);
}

func draw(graphics GraphicsContext) {
	graphics.FillRect(Rect{0, 0, 1024, 768}, Color{0, 0, 0, 255});
	graphics.FillRect(Rect{10, 10, 100, 100}, Color{255, 0, 0, 255});
}