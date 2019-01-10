package UI

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

type UIDoor struct {
	DrawableClickable
	num int
}

func NewUIDoor(x, y, w, h, num int) *UIDoor {
	d := &UIDoor{
		num: num,
	}
	pic, _ := ebiten.NewImage(w, h, ebiten.FilterLinear)
	_ = pic.Fill(color.Black)
	d.DCInit(x, y, w, h, 1, NewSprite(pic))
	return d
}

func (d *UIDoor) Draw(screen *ebiten.Image, col color.Color) {
	d.DrawImg(screen)
	//ebitenutil.DrawRect(screen, float64(d.x), float64(d.y), float64(d.w), float64(d.h), col)
}
