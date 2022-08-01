package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	err        error
	background *ebiten.Image
)

// init 初始化
func init() {
	background, _, err = ebitenutil.NewImageFromFile("assets/space.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

// update 更新画布
func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	op := &ebiten.DrawImageOptions{}
	// 屏幕起始坐标
	op.GeoM.Translate(0, 0)
	screen.DrawImage(background, op)

	return nil
}

func main() {
	if err := ebiten.Run(update, 640, 480, 1, "Hello World!"); err != nil {
		log.Fatal(err)
	}
}
