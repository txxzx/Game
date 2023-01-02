package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

/**
    @date: 2023/1/2
**/

type Bullet struct {
	image       *ebiten.Image
	width       int
	height      int
	x           float64
	y           float64
	speedFactor float64
}

func NewBullet(g *Game) *Bullet {
	rect := image.Rect(0, 0, g.cfg.BulletWidth, g.cfg.BulletHeight)
	img := ebiten.NewImageWithOptions(rect, nil)
	img.Fill(g.cfg.BulletColor)

	return &Bullet{
		image:       img,
		width:       cfg.BulletWidth,
		height:      cfg.BulletHeight,
		x:           g.ship.X + float64(g.ship.Width-cfg.BulletWidth)/2,
		y:           float64(cfg.ScreenHeight - g.ship.Height - cfg.BulletHeight),
		speedFactor: cfg.BulletSpeedFactor,
	}
}

// 增加子弹的绘制方法
func (bullet *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(bullet.x, bullet.y)
	screen.DrawImage(bullet.image, op)
}

// 判断子弹是否在屏幕外面
func (bullet *Bullet) outOfScreen() bool {
	return bullet.y < -float64(bullet.height)
}
