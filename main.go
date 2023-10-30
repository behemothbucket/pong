package main

import (
	"errors"
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	field   *Field
	racket1 *Racket
	racket2 *Racket
	ball    *Ball
	score   []int
	scored  bool
	timeout int
}

func main() {
	ebiten.SetWindowSize(screenWidth*1.5, screenHeight*1.5)
	ebiten.SetWindowTitle("Pong")

	game := newGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func newGame() *Game {
	field := &Field{
		fieldLine: &FieldLine{
			x:       screenWidth / 2,
			y:       0,
			width:   2,
			spacing: 18,
			color:   &color.White,
		},
	}

	racket1 := &Racket{
		x:      10,
		y:      screenHeight / 2,
		width:  10,
		height: 50,
		color:  &color.White,
	}

	racket2 := &Racket{
		x:      screenWidth - racket1.width - 10,
		y:      screenHeight / 2,
		width:  10,
		height: 50,
		color:  &color.White,
	}

	ball := &Ball{
		radius: 10,
		color:  &color.White,
	}
	ball.ServeBall()

	score := make([]int, 2)
	scored := false
	timeout := 0

	return &Game{field, racket1, racket2, ball, score, scored, timeout}
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) || inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return errors.New("Quit game")
	}

	g.racket1.Update()
	g.racket2.Update()
	g.ball.Update(g)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.field.Draw(screen)
	g.racket1.Draw(screen)
	g.racket2.Draw(screen)
	g.ball.Draw(screen)
	g.DrawScore(screen)
}

func (g *Game) DrawScore(screen *ebiten.Image) {
	scoreText := fmt.Sprintf("FPS:%0.f\ndX:%0.f\ndY:%0.f\nspeed:%0.f", ebiten.ActualFPS(), g.ball.directionX, g.ball.directionY, g.ball.speed)
	ebitenutil.DebugPrint(screen, scoreText)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
