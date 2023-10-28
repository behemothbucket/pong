package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Ball struct {
	color  *color.Gray16
	x      float32
	y      float32
	speedX float32
	speedY float32
	radius float32
}

func (b *Ball) HandleCollisions(g *Game) {
	checkWallCollision(b, g)
	checkRacketCollision(g, b, g.racket1)
	checkRacketCollision(g, b, g.racket2)
}

func checkWallCollision(b *Ball, g *Game) {
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
		b.speedY *= -1
	}
}

func checkRacketCollision(g *Game, b *Ball, racket *Racket) {
	if b.x-b.radius < racket.x+racket.width && b.x+b.radius > racket.x &&
		b.y+b.radius > racket.y && b.y-b.radius < racket.y+racket.height {
		random := getRandomSource()

		if racket == g.racket1 {
			b.speedX, _ = getRandomSpeed(random, 4, 8)
		} else if racket == g.racket2 {
			b.speedX, _ = getRandomSpeed(random, -8, -4)
		}
	}
}

func (b *Ball) Update(g *Game) {
	if g.timeout > 0 {
		g.timeout--
		return
	}

	if g.scored {
		b.ServeTheBall()
		g.scored = false
	}

	b.x += b.speedX
	b.y += b.speedY

	b.HandleCollisions(g)
}

func (b *Ball) ServeTheBall() {
	random := getRandomSource()

	b.x = screenWidth / 2
	b.y = float32(random.Intn(460) + 1)

	b.speedX, b.speedY = getRandomSpeed(random, 4, 8)
}

// FIX: направление мяча, случайная скорость
func getRandomSpeed(random *rand.Rand, min, max float32) (float32, float32) {
	speeds := []float32{min, max}
	index1 := random.Intn(len(speeds))
	index2 := random.Intn(len(speeds))
	return speeds[index1], speeds[index2]
}

func getRandomSource() (r *rand.Rand) {
	source := rand.NewSource(time.Now().UnixNano())
	return rand.New(source)
}

func (b *Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, b.x, b.y, b.radius, b.color, false)
}
