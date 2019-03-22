package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 800, 600),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Skyblue)

	imd := imdraw.New(nil)

	imd.Color = pixel.RGB(1, 0, 0)
	imd.Push(pixel.V(100, 100))
	imd.Color = pixel.RGB(0, 1, 0)
	imd.Push(pixel.V(700, 100))
	imd.Color = pixel.RGB(0, 0, 1)
	imd.Push(pixel.V(400, 500))
	imd.Polygon(0)

	for !win.Closed() {
		win.Clear(colornames.Aliceblue)
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
