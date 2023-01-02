package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
)

/**
    @date: 2023/1/2
**/

// 游戏的方法
type Game struct {
	// 控制飞船移动
	input  *Input
	ship   *Ship
	// 配置文件
	cfg   *Config
	// 构建子弹
	bullets map[*Bullet]struct{}
	// 外星人
	aliens map[*Alien]struct{}
	mode  Mode
	failCount int // 被外星人碰撞和移出屏幕的外星人数量之和
}

func NewGame() *Game {
	// 加载json的配置文件
	cfg := loadConfig()
	// 设置屏幕的大小
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	// 设置屏幕的标题
	ebiten.SetWindowTitle(cfg.Title)
	g := &Game{
		input:   &Input{},
		ship:    NewShip(cfg.ScreenWidth,cfg.ScreenHeight),
		cfg:     cfg,
		bullets: make(map[*Bullet]struct{}),
		aliens: make(map[*Alien]struct{}),
	}
	// 创建一组外星人
	g.CreateAliens()
	g.init()
	return g
}

// 创建外星人方法
func (g *Game) CreateAliens() {
	alien := NewAlien(g.cfg)
	// 左右各留一个外星人宽度的空间：
	availableSpaceX := g.cfg.ScreenWidth - 2*alien.width
	// 两个外星人之间留一个外星人宽度的空间。所以一行可以创建的外星人的数量为：
	numAliens := availableSpaceX / (2 * alien.width)
	for row := 0;row < 2;row ++ {
		for i := 0; i < numAliens; i++ {
			alien = NewAlien(g.cfg)
			alien.x = float64(alien.width + 2*alien.width*i)
			alien.y = float64(alien.height*row) * 1.5
			g.AddAliens(alien)
		}
	}
}
// Game结构的Update方法中，我们需要调用Input的Update方法触发按键的判断：
func (g *Game) Update() error {
	switch g.mode {
	case ModeTitle:
		if g.input.IsKeyPressed() {
			g.mode = ModeGame
		}
	case ModeGame:
		for bullet := range g.bullets {
			bullet.y -= bullet.speedFactor
		}

		for alien := range g.aliens {
			alien.y += alien.speedFactor
		}

		g.input.Update(g)

		g.CheckCollision()

		for bullet := range g.bullets {
			if bullet.outOfScreen() {
				delete(g.bullets, bullet)
			}
		}
	case ModeOver:
		if g.input.IsKeyPressed() {
			g.init()
			g.mode = ModeTitle
		}
	}

	return nil
}

// 绘制方法
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.cfg.BgColor)

	var titleTexts []string
	var texts []string
	switch g.mode {
	case ModeTitle:
		titleTexts = []string{"ALIEN INVASION"}
		texts = []string{"", "", "", "", "", "", "", "PRESS SPACE KEY", "", "OR LEFT MOUSE"}
	case ModeGame:
		g.ship.Draw(screen,g.cfg)
		for bullet := range g.bullets {
			bullet.Draw(screen)
		}
		for alien := range g.aliens {
			alien.Draw(screen)
		}
	case ModeOver:
		texts = []string{"", "GAME OVER!"}
	}

	for i, l := range titleTexts {
		x := (g.cfg.ScreenWidth - len(l)*g.cfg.TitleFontSize) / 2
		text.Draw(screen, l, titleArcadeFont, x, (i+4)*g.cfg.TitleFontSize, color.White)
	}
	for i, l := range texts {
		x := (g.cfg.ScreenWidth - len(l)*g.cfg.FontSize) / 2
		text.Draw(screen, l, arcadeFont, x, (i+4)*g.cfg.FontSize, color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.cfg.ScreenWidth / 2, g.cfg.ScreenHeight / 2
}

// 添加子弹的方法
func (g *Game) AddBullet(bullet *Bullet) {
	g.bullets[bullet] = struct{}{}
}

// 添加外星人
func (g *Game) AddAliens(alien *Alien) {
	g.aliens[alien] = struct{}{}
}

// 将碰撞的子弹和外星人删除。
func (g *Game) CheckCollision() {
	for alien := range g.aliens {
		for bullet := range g.bullets {
			if CheckCollision(bullet, alien) {
				delete(g.aliens, alien)
				delete(g.bullets, bullet)
			}
		}
	}
}

func (g *Game) init() {
	g.CreateAliens()
	g.CreateFonts()
}


var (
	titleArcadeFont font.Face
	arcadeFont      font.Face
	smallArcadeFont font.Face
)

// 创建字体
func (g *Game) CreateFonts() {
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	titleArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.cfg.TitleFontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	arcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.cfg.FontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	smallArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.cfg.SmallFontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}