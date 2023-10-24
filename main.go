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

type Game struct {
	rocket *Rocket
	ball   *Ball
	field  *Field
}

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

type Rocket struct {
	x, y   float32
	width  float32
	height float32
	color  *color.Gray16
}

type Ball struct {
	x, y           float32
	speedX, speedY float32
	size           float32
	color          *color.Gray16
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("        Pong")

	game := newGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatalf("Ошибка при запуске игры: %v", err)
	}
}

func newGame() *Game {
	field := &Field{
		fieldLine: &FieldLine{
			x:       screenWidth / 2,
			y:       0,
			width:   2,
			spacing: 10,
			color:   &color.White,
		},
	}

	rocket := &Rocket{
		x:      10,
		y:      screenHeight / 2,
		width:  10,
		height: 50,
		color:  &color.White,
	}

	ball := &Ball{
		x:      screenWidth / 2,
		y:      screenHeight / 2,
		size:   20,
		speedX: 6,
		speedY: 6,
		color:  &color.White,
	}

	return &Game{
		field:  field,
		rocket: rocket,
		ball:   ball,
	}
}

func (fl *FieldLine) Draw(screen *ebiten.Image) {
	for y := fl.y; y < screenHeight; y += fl.spacing {
		vector.DrawFilledRect(screen, fl.x, y, fl.width, fl.width+fl.spacing/2, fl.color, false)
	}
}

func (f *Field) Draw(screen *ebiten.Image) {
	f.fieldLine.Draw(screen)
}

func (r *Rocket) Update() {
	_, mouseY := ebiten.CursorPosition()
	r.y = float32(mouseY) - r.height/2
}

func (r *Rocket) Draw(screen *ebiten.Image) {
	ry := r.y

	if ry < 0 {
		ry = 0
	} else if ry > screenHeight-r.height {
		ry = screenHeight - r.height
	}

	vector.DrawFilledRect(screen, r.x, ry, r.width, r.height, r.color, false)
}

func (b *Ball) HandleCollisions() {
	if b.x < 0 || b.x > screenWidth-b.size {
		b.speedX = -b.speedX
	}
	if b.y < 0 || b.y > screenHeight-b.size {
		b.speedY = -b.speedY
	}
}

func (b *Ball) Update() {
	b.x += b.speedX
	b.y += b.speedY

	b.HandleCollisions()
}

func (b *Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, b.x, b.y, b.size, b.size, b.color, false)
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("quit")
	}

	g.rocket.Update()
	g.ball.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.field.Draw(screen)
	g.rocket.Draw(screen)
	g.ball.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
