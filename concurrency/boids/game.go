package boid

import (
	_ "fmt"
	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2"
	_ "image/color"
	"log"
	_ "log"
)

const (
	screenWidth   = 640
	screenHeight  = 360
	boidCount     = 500
	viewRadius    = 13
	adjustingRate = 0.015
)

var (
	boids     [boidCount]*Boid
	boidArray [screenWidth + 1][screenHeight + 1]int
)

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	for _, boid := range boids {
		screen.Set(int(boid.position.x+1), int(boid.position.y), boid.boidColor)
		screen.Set(int(boid.position.x-1), int(boid.position.y), boid.boidColor)
		screen.Set(int(boid.position.x), int(boid.position.y-1), boid.boidColor)
		screen.Set(int(boid.position.x), int(boid.position.y+1), boid.boidColor)
	}
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return screenWidth, screenHeight
}

func RunGame() {

	for i, row := range boidArray {
		for j := range row {
			boidArray[i][j] = -1
		}
	}

	for i := 0; i < boidCount; i++ {
		CreateBoid(i)
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Boids")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
