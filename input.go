package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

/**
    @date: 2023/1/2
**/

// ebiten提供函数IsKeyPressed来判断某个键是否按下，同时内置了100多个键的常量定义，
//见源码keys.go文件。ebiten.KeyLeft表示左方向键，ebiten.KeyRight表示右方向键，ebiten.KeySpace表示空格。
type Input struct {
	// msg 字段
	msg string
	// 上次发射子弹的时间
	lastBulletTime time.Time
}

func (i *Input) Update(g *Game) {
	// 判断按键是否为左键
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.ship.X -= cfg.ShipSpeedFactor
		if g.ship.X < -float64(g.ship.Width)/2 {
			g.ship.X = -float64(g.ship.Width) / 2
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.ship.X += cfg.ShipSpeedFactor
		if g.ship.X > float64(cfg.ScreenWidth)-float64(g.ship.Width)/2 {
			g.ship.X = float64(cfg.ScreenWidth) - float64(g.ship.Width)/2

		}
	}else if ebiten.IsKeyPressed(ebiten.KeySpace) &&  time.Now().Sub(i.lastBulletTime).Milliseconds() > g.cfg.BulletInterval {
		if len(g.bullets) <cfg.MaxBulletNum {
			bullet := NewBullet(g)
			g.AddBullet(bullet)
			i.lastBulletTime = time.Now()
		}
	}
}

func (i *Input) IsKeyPressed() bool{
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeySpace){
		return true
	}
	return false
}