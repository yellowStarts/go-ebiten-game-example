package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	err        error
	background *ebiten.Image // 背景
	spaceShip  *ebiten.Image // 太空船
	playerOne  player        // 玩家
)

const (
	screenWidth, screebHeight = 640, 480
)

// player 玩家结构体
type player struct {
	image      *ebiten.Image
	xPos, yPos float64
	speed      float64
}

// init 初始化
func init() {
	// 背景
	background, _, err = ebitenutil.NewImageFromFile("assets/space.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	// 太空船
	spaceShip, _, err = ebitenutil.NewImageFromFile("assets/space_ship.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	// player
	playerOne = player{spaceShip, screenWidth / 2.0, screebHeight / 2.0, 4}
}

// movePlayer 玩家移动
func movePlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		playerOne.yPos -= playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		playerOne.yPos += playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		playerOne.xPos -= playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		playerOne.xPos += playerOne.speed
	}
}

// update 更新画布
func update(screen *ebiten.Image) error {
	// 更新移动方法
	movePlayer()

	if ebiten.IsDrawingSkipped() {
		return nil
	}
	op := &ebiten.DrawImageOptions{}
	// 屏幕起始坐标
	op.GeoM.Translate(0, 0)
	screen.DrawImage(background, op)

	// player
	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(playerOne.xPos, playerOne.yPos)
	screen.DrawImage(playerOne.image, playerOp)

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screebHeight, 1, "Hello World!"); err != nil {
		log.Fatal(err)
	}
}
