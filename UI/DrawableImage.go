package UI

import "github.com/hajimehoshi/ebiten"

type Drawable interface {
	DrawImg(screen *ebiten.Image) error
}

type DrawableImage struct {
	pic  *ebiten.Image
	opts *ebiten.DrawImageOptions
}

func (d *DrawableImage) DrawImg(screen *ebiten.Image) error {
	return screen.DrawImage(d.pic, d.opts)
}
