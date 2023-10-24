package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Field struct {
	fieldLine *FieldLine
}

type FieldLine struct {
	x       float32
	y       float32
	width   float32
	spacing float32
	color   *color.Gray16
}

func (fl *FieldLine) Draw(screen *ebiten.Image) {
	for y := fl.y; y < screenHeight; y += fl.spacing {
		vector.DrawFilledRect(screen, fl.x, y, fl.width, fl.width+fl.spacing/2, fl.color, false)
	}
}

func (f *Field) Draw(screen *ebiten.Image) {
	f.fieldLine.Draw(screen)
}
