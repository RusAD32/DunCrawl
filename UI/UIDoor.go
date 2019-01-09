package UI

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

type UIDoor struct {
	ClickableRect
	num int
	DrawableImage
}

func (d *UIDoor) Init(x, y, w, h, num int) *UIDoor {
	d.x = x
	d.y = y
	d.w = w
	d.h = h
	d.num = num
	d.pic = make([]*ebiten.Image, 1)
	d.pic[0], _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	_ = d.pic[0].Fill(color.Black)
	d.opts = &ebiten.DrawImageOptions{}
	d.opts.GeoM.Translate(float64(x), float64(y))
	return d
}

func (d *UIDoor) Draw(screen *ebiten.Image, col color.Color) {
	d.DrawImg(screen)
	//ebitenutil.DrawRect(screen, float64(d.x), float64(d.y), float64(d.w), float64(d.h), col)
}
