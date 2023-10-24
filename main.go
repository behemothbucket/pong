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
	racket  *Racket
	ball    *Ball
	score   []int
	scored  bool
	timeout int
}

func main() {
	ebiten.SetWindowSize(screenWidth*1.5, screenHeight*1.5)
	ebiten.SetWindowTitle("        Pong")

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
			spacing: 10,
			color:   &color.White,
		},
	}

	racket := &Racket{
		x:      10,
		y:      screenHeight / 2,
		width:  10,
		height: 50,
		color:  &color.White,
	}

	ball := &Ball{
		radius: 5,
		color:  &color.White,
	}
	ball.RandomDirection()

	score := make([]int, 2)
	scored := false
	timeout := 0

	return &Game{field, racket, ball, score, scored, timeout}
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) || inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return errors.New("Quit")
	}

	g.racket.Update()
	g.ball.Update(g)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.field.Draw(screen)
	g.racket.Draw(screen)
	g.ball.Draw(screen)
	g.DrawScore(screen)
}

func (g *Game) DrawScore(screen *ebiten.Image) {
	scoreText := fmt.Sprintf("%d:%d", g.score[0], g.score[1])
	ebitenutil.DebugPrint(screen, scoreText)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
