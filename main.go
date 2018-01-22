package main

import (
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func fragment(imd *imdraw.IMDraw, lines []Line) []Line {
	result := make([]Line, len(lines)*4)
	for i := range lines {
		for _, l := range lines[i].Koch() {
			result = append(result, l)
			l.Draw(imd, 1)
		}
	}
	return result
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Line fragments",
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	lines := []Line{
		NewLine(
			pixel.V(100, 200),
			pixel.V(900, 200),
			colornames.Coral,
		),
		NewLine(
			pixel.V(500, 430.94),
			pixel.V(100, 200),
			colornames.Forestgreen,
		),
		NewLine(
			pixel.V(900, 200),
			pixel.V(500, 430.94),
			colornames.Mediumaquamarine,
		),

		//NewLine(pixel.V(800, 100), pixel.V(200, 100)),
		//NewLine(pixel.V(500, 700), pixel.V(800, 100)),
		//NewLine(pixel.V(200, 100), pixel.V(500, 700)),
	}

	imd := imdraw.New(nil)
	imd.Color = colornames.Blueviolet
	for _, l := range lines {
		l.Draw(imd, 1)
	}

	for !win.Closed() {
		if win.Pressed(pixelgl.KeyEscape) {
			log.Println("Exiting...")
			break
		}
		if win.JustPressed(pixelgl.KeySpace) {
			log.Printf("Fragmenting %d lines\n", len(lines))
			imd.Clear()
			lines = fragment(imd, lines)
		}

		win.Clear(colornames.Aliceblue)
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
