package UI

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

type UIDoor struct {
	DrawableClickable
	num int
}

func (d *UIDoor) Init(x, y, w, h, num int) *UIDoor {
	d.initRect(x, y, w, h)
	d.num = num
	d.DCInit(x, y, w, h, 1, color.Black)
	return d
}

func (d *UIDoor) Draw(screen *ebiten.Image, col color.Color) {
	d.DrawImg(screen)
	//ebitenutil.DrawRect(screen, float64(d.x), float64(d.y), float64(d.w), float64(d.h), col)
}
