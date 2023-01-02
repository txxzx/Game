package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

/**
    @date: 2023/1/2
**/

func main() {
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}