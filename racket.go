package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Racket struct {
	color  *color.Gray16
	x      float32
	y      float32
	width  float32
	height float32
}

func (r *Racket) Update() {
	_, mouseY := ebiten.CursorPosition()
	r.y = float32(mouseY) - r.height/2
}

func (r *Racket) getLimitedY() float32 {
	if r.y < 0 {
		return 0
	}

	if r.y > screenHeight-r.height {
		return screenHeight - r.height
	}

	return r.y
}

func (r *Racket) Draw(screen *ebiten.Image) {
	y := r.getLimitedY()
	vector.DrawFilledRect(screen, r.x, y, r.width, r.height, r.color, false)
}
