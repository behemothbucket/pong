package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Ball struct {
	x, y           float32
	speedX, speedY float32
	radius         float32
	color          *color.Gray16
}

func (b *Ball) HandleCollisions(g *Game) {
	if b.x < 0 || b.x > screenWidth-b.radius {
		b.speedX = -b.speedX

		playerIndex := 0
		if b.x < 0 {
			playerIndex = 1
		}

		g.score[playerIndex]++
	}

	if b.y < 0 || b.y > screenHeight-b.radius {
		b.speedY = -b.speedY
	}
}

func (b *Ball) Update(g *Game) {
	b.x += b.speedX
	b.y += b.speedY

	b.HandleCollisions(g)
}

func (b *Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, b.x, b.y, b.radius, b.color, false)
}
