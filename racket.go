package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Racket struct {
	x, y   float32
	width  float32
	height float32
	color  *color.Gray16
}

func (r *Racket) Update() {
	_, mouseY := ebiten.CursorPosition()
	r.y = float32(mouseY) - r.height/2
}

func (r *Racket) Draw(screen *ebiten.Image) {
	ry := r.y

	if ry < 0 {
		ry = 0
	} else if ry > screenHeight-r.height {
		ry = screenHeight - r.height
	}

	vector.DrawFilledRect(screen, r.x, ry, r.width, r.height, r.color, false)
}
