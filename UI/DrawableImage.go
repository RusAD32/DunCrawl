package UI

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

type Drawable interface {
	DrawImg(screen *ebiten.Image)
}

type DrawableImage struct {
	pic   []*ebiten.Image
	opts  *ebiten.DrawImageOptions
	state int
}

func (d *DrawableImage) initImg(x, y, w, h, length int, cols ...color.Color) {
	d.opts = &ebiten.DrawImageOptions{}
	d.opts.GeoM.Translate(float64(x), float64(y))
	d.pic = make([]*ebiten.Image, length)
	for i := range d.pic {
		d.pic[i], _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
		_ = d.pic[i].Fill(cols[i])
	}

}

func (d *DrawableImage) ChoosePic() *ebiten.Image {
	return d.pic[d.state]
}

func (d *DrawableImage) DrawImg(screen *ebiten.Image) {
	_ = screen.DrawImage(d.ChoosePic(), d.opts)
}
