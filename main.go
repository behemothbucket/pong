package main

import (
	"errors"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

type Game struct{}

var (
	rocket *Rocket
	ball   *Ball
)

type Rocket struct {
	x, y   float32
	width  float32
	height float32
}

type Ball struct {
	x, y           float32
	speedX, speedY float32
	size           float32
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("quit")
	}

	_, mouseY := ebiten.CursorPosition()

	rocket.y = float32(mouseY) - rocket.height/2

	ball.x += ball.speedX
	ball.y += ball.speedY

	if ball.x < 0 || ball.x > screenWidth-ball.size {
		ball.speedX = -ball.speedX
	}
	if ball.y < 0 || ball.y > screenHeight-ball.size {
		ball.speedY = -ball.speedY
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// FIX StrokeLine -> DrawFilledRect
	var lineX float32 = screenWidth / 2
	var lineY1 float32 = 0
	var lineY2 float32 = screenHeight
	var lineSpacing float32 = 10

	lineColor := color.White
	var lineWidth float32 = 2

	for y := lineY1; y < lineY2; y += lineSpacing {
		vector.StrokeLine(screen, lineX, y, lineX, y+lineSpacing/2, lineWidth, lineColor, false)
	}

	vector.DrawFilledRect(screen, ball.x, ball.y, ball.size, ball.size, color.White, false)

	if rocket.y >= screenHeight-rocket.height {
		vector.DrawFilledRect(screen, rocket.x, screenHeight-rocket.height, rocket.width, rocket.height, lineColor, false)
	}

	if rocket.y <= 0 {
		vector.DrawFilledRect(screen, rocket.x, 0, rocket.width, rocket.height, lineColor, false)
	} else {
		vector.DrawFilledRect(screen, rocket.x, rocket.y, rocket.width, rocket.height, lineColor, false)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("        Pong")

	rocket = &Rocket{
		x:      10,
		y:      screenHeight / 2,
		width:  10,
		height: 50,
	}

	ball = &Ball{
		x:      screenWidth / 2,
		y:      screenHeight / 2,
		size:   20,
		speedX: 6,
		speedY: 6,
	}

	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
