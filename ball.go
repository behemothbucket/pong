package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Ball struct {
	color      *color.Gray16
	x          float32
	y          float32
	directionX float32
	directionY float32
	speed      float32
	radius     float32
}

var randomSource = rand.NewSource(time.Now().UnixNano())

func (b *Ball) handleCollisions(g *Game) {
	b.checkWallCollision(g)
	b.checkRacketCollision(g, g.racket1)
	b.checkRacketCollision(g, g.racket2)
}

func (b *Ball) checkWallCollision(g *Game) {
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
		b.directionY *= -1
	}
}

// FIX: верхняя и нижняя граница не обрабатывется
func (b *Ball) checkRacketCollision(g *Game, r *Racket) {
	if b.x >= r.x && b.x <= r.x+r.width {
		if b.y >= r.y-(r.y/2) && b.y <= r.y+(r.height/2) {
			b.speed = getRandomSpeed(1, 3)
			b.directionX *= -1
		}
	}
}

func (b *Ball) Update(g *Game) {
	if g.timeout > 0 {
		g.timeout--
		return
	}

	if g.scored {
		b.ServeBall()
		g.scored = false
	}

	b.x += b.directionX * b.speed
	b.y += b.directionY * b.speed

	b.handleCollisions(g)
}

func (b *Ball) ServeBall() {
	b.x = screenWidth / 2
	b.y = screenHeight / 2
	b.directionX, b.directionY = getRandomDirection()
	b.speed = getRandomSpeed(5, 6)
}

func getRandomSpeed(min, max float32) float32 {
	random := rand.New(randomSource)
	return min + random.Float32()*(max-min)
}

func getRandomDirection() (float32, float32) {
	random := rand.New(randomSource)

	directions := []float32{-1, 1}

	directionX := directions[random.Intn(2)]
	directionY := directions[random.Intn(2)]

	return directionX, directionY
}

func (b *Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, b.x, b.y, b.radius, b.radius, b.color, false)
}
