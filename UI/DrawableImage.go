package UI

import "github.com/hajimehoshi/ebiten"

type Drawable interface {
	DrawImg(screen *ebiten.Image)
}

type DrawableImage struct {
	pic   []*ebiten.Image
	opts  *ebiten.DrawImageOptions
	state int
}

func (d *DrawableImage) ChoosePic() *ebiten.Image {
	return d.pic[d.state]
}

func (d *DrawableImage) DrawImg(screen *ebiten.Image) {
	_ = screen.DrawImage(d.ChoosePic(), d.opts)
}
