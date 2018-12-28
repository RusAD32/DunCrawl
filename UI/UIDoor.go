package UI

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
)

type UIDoor struct {
	x, y, w, h int
	num        int
}

func (d *UIDoor) isClicked(mouseX, mouseY int) bool {
	return !(mouseX < d.x || mouseX > d.x+d.w || mouseY < d.y || mouseY > d.y+d.h)
}

func (d *UIDoor) Draw(screen *ebiten.Image, col color.Color) {
	ebitenutil.DrawRect(screen, float64(d.x), float64(d.y), float64(d.w), float64(d.h), col)
}
