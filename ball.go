package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Ball struct {
	x, y       float32
	directionX float32
	directionY float32
	radius     float32
	color      *color.Gray16
}

func (b *Ball) HandleCollisions(g *Game) {
	if b.x < -b.radius || b.x > screenWidth+b.radius {
		playerIndex := 0

		if b.x < 0 {
			playerIndex = 1
		}

		g.score[playerIndex]++
		g.scored = true
		g.timeout = ebiten.TPS()
	}

	if b.y < 0 || b.y > screenHeight-b.radius {
		b.directionY = -b.directionY
	}
}

func (b *Ball) Update(g *Game) {
	if g.timeout > 0 {
		g.timeout--
		return
	}

	if g.scored {
		b.RandomDirection()
		g.scored = false
	}

	b.x += b.directionX
	b.y += b.directionY

	b.HandleCollisions(g)
}

func (b *Ball) RandomDirection() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	b.x = screenWidth / 2
	b.y = float32(random.Intn(480) + 1)

	b.directionX = float32(4 - 8*random.Intn(2))
	b.directionY = float32(4 - 8*random.Intn(2))
}

func (b *Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, b.x, b.y, b.radius, b.color, false)
}
