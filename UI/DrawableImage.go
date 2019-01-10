package UI

import (
	"github.com/hajimehoshi/ebiten"
)

type Drawable interface {
	DrawImg(screen *ebiten.Image)
}

type DrawableImage struct {
	pic   []*Sprite
	opts  *ebiten.DrawImageOptions
	state int
}

func (d *DrawableImage) initImg(x, y, w, h, length int, imgs ...*Sprite) {
	d.opts = &ebiten.DrawImageOptions{}
	imgW, imgH := imgs[0].frames[0].Size()
	d.opts.GeoM.Scale(float64(w)/float64(imgW), float64(h)/float64(imgH))
	d.opts.GeoM.Translate(float64(x), float64(y))
	d.pic = imgs
}

func (d *DrawableImage) ChoosePic() *ebiten.Image {
	return d.pic[d.state].GetCurrentFrame()
}

func (d *DrawableImage) DrawImg(screen *ebiten.Image) {
	_ = screen.DrawImage(d.ChoosePic(), d.opts)
}
