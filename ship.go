package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "golang.org/x/image/bmp"
	"log"
)

/**
    @date: 2023/1/2
**/

// 定义一个飞船的结构体
type Ship struct {
	Image  *ebiten.Image
	Width  int
	Height int
	// x坐标
	X float64
	// y坐标
	Y float64
}

// 定义一个飞船的方法
func NewShip(screenWidth,screenHeight int) *Ship {
	// 加载飞船的图片
	img, _, err := ebitenutil.NewImageFromFile("./b2f7e8dc94b4743526d95abff42b7412.bmp")
	if err != nil {
		log.Fatal(err)
	}

	// 加载飞船的宽高以及飞船的图片
	width, height := img.Size()
	ship := &Ship{
		Image:  img,
		X: float64(screenWidth-width) / 2,
		Y: float64(screenHeight- height),
	}
	return ship
}
// 给ship增加一个绘制自身的方法，传入屏幕宽度方便维护
func (ship *Ship) Draw(screen *ebiten.Image, cfg *Config) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(cfg.ScreenWidth-ship.Width)/2, float64(cfg.ScreenHeight-ship.Height))
	screen.DrawImage(ship.Image, op)
}