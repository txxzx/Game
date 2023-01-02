package main

import (
	"github.com/usthooz/gconf"
	"image/color"
	"log"
)

/**
    @date: 2023/1/2
**/

// json文件
type Config struct {
	ScreenWidth  int        `json:"screenWidth"`
	ScreenHeight int        `json:"screenHeight"`
	Title        string     `json:"title"`
	BgColor      color.RGBA `json:"bgColor"`
	// 飞船移动速度
	ShipSpeedFactor float64    `json:"shipSpeedFactor"`
	// 子弹的宽度
	BulletWidth int `json:"bulletWidth"`
	// 子弹的高度
	BulletHeight int `json:"bulletHeight"`
	// 子弹的射击速度
	BulletSpeedFactor float64 `json:"bulletSpeed_factor"`
	// 子弹颜色
	BulletColor  color.RGBA `json:"bulletColor"`
	// 同时发出多个子弹
	MaxBulletNum      int        `json:"maxBulletNum"`
    // 子弹发射的毫秒数
	BulletInterval    int64      `json:"bulletInterval"`
	SpeedFactor float64 `json:"speedFactor"`
	SmallFontSize int `json:"smallFontSize"`
	TitleFontSize int `json:"titleFontSize"`
	FontSize int `json:"fontSize"`
	AlienSpeedFactor float64 `json:"alienSpeedFactor"`
}
var cfg Config

func loadConfig() *Config {
	logicConf := gconf.NewConf(&gconf.Gconf{
		ConfPath: "./config.json",
		Subffix:  gconf.JsonSub,
	})
	if err := logicConf.GetConf(&cfg); err != nil {
		log.Fatalf("os.Open failed: %v\n", err)
	}

	return &cfg
}