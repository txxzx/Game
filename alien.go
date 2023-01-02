package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

/**
    @date: 2023/1/2
**/

type Alien struct {
	image       *ebiten.Image
	width       int
	height      int
	x           float64
	y           float64
	speedFactor float64
}

func NewAlien(cfg *Config) *Alien {
	img, _, err := ebitenutil.NewImageFromFile("./alien.png")
	if err != nil {
		log.Fatal(err)
	}


	width, height := img.Size()
	return &Alien{
		image:       img,
		width:       width,
		height:      height,
		x:           0,
		y:           0,
		speedFactor: cfg.AlienSpeedFactor,
	}
}

func (alien *Alien) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(alien.x, alien.y)
	screen.DrawImage(alien.image, op)
}

func (alien *Alien) outOfScreen() bool {
	return alien.y < -float64(alien.height)
}