package main

import (
	"fmt"

	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
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

	// Text
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(10, 550), basicAtlas)
	basicTxt4 := text.New(pixel.V(10, 500), basicAtlas)
	txt := text.New(pixel.V(10, 450), basicAtlas)
	basicTxt.Color = colornames.Red
	basicTxt4.Color = colornames.Green
	txt.Color = colornames.Blue
	fmt.Fprintln(basicTxt, "Hello, text!")
	fmt.Fprintln(basicTxt4, "Type some keys")

	imd := imdraw.New(nil)

	// Triangle
	imd.Color = pixel.RGB(1, 0, 0)
	imd.Push(pixel.V(100, 100))
	imd.Color = pixel.RGB(0, 1, 0)
	imd.Push(pixel.V(700, 100))
	imd.Color = pixel.RGB(0, 0, 1)
	imd.Push(pixel.V(400, 500))
	imd.Polygon(0)

	// Line
	imd.Color = colornames.Blueviolet
	imd.Push(pixel.V(0, 0), pixel.V(800, 600))
	imd.Line(1)

	for !win.Closed() {
		txt.WriteString(win.Typed())
		if win.JustPressed(pixelgl.KeyEnter) || win.Repeated(pixelgl.KeyEnter) {
			txt.WriteRune('\n')
		}

		win.Clear(colornames.Darkgray)
		imd.Draw(win)

		basicTxt.Draw(win, pixel.IM)
		basicTxt4.Draw(win, pixel.IM.Scaled(basicTxt4.Orig, 4))
		txt.Draw(win, pixel.IM.Scaled(txt.Orig, 2))

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
