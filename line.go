package main

import (
	"image/color"
	"log"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Line struct {
	Start, End, v pixel.Vec
	Color         color.RGBA
}

func NewLine(start, end pixel.Vec, c color.RGBA) (l Line) {
	l = Line{
		Start: start,
		End:   end,
		Color: c,
		v:     end.Sub(start).Scaled(float64(1) / 3),
	}
	//log.Printf("New line: (%s, %s)\n", start, end)
	return l
}

func (l Line) Koch() (lines [4]Line) {
	v := l.End.Sub(l.Start).Scaled(float64(1) / 3)

	p1 := l.Start.Add(v)
	lines[0] = NewLine(l.Start, p1, l.Color)

	p2 := p1.Add(v.Rotated(math.Pi / 3))
	log.Printf("p2=%s\n", p2)
	lines[1] = NewLine(p1, p2, l.Color)

	p3 := p2.Add(v.Rotated(-math.Pi / 3))
	lines[2] = NewLine(p2, p3, l.Color)

	lines[3] = NewLine(p3, l.End, l.Color)
	return
}

func (l Line) Square() (lines [5]Line) {
	p1 := l.Start.Add(l.v)
	lines[0] = NewLine(l.Start, p1, l.Color)

	p2 := p1.Add(l.v.Rotated(math.Pi / 2))
	lines[1] = NewLine(p1, p2, l.Color)

	p3 := p2.Add(l.v)
	lines[2] = NewLine(p2, p3, l.Color)

	p4 := p3.Add(l.v.Rotated(-math.Pi / 2))
	lines[3] = NewLine(p3, p4, l.Color)

	lines[4] = NewLine(p4, l.End, l.Color)
	return
}

func (l Line) Draw(imd *imdraw.IMDraw, width float64) {
	imd.Color = l.Color
	imd.Push(l.Start)
	imd.Push(l.End)
	imd.Line(width)
}

type LineFragmenter struct {
	Lines    []Line
	LinesOut chan Line

	imd *imdraw.IMDraw
}

func NewLineFragmenter(imd *imdraw.IMDraw, method string) (lf LineFragmenter) {
	lf = LineFragmenter{
		imd: imd,
	}
	return
}
