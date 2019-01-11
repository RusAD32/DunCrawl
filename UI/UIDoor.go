package UI

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

type UIDoor struct {
	DrawableClickable
	num     int
	visible bool
}

func NewUIDoor(x, y, w, h, num int, t *TexPreloader) *UIDoor {
	doorPics := []string{
		"resources/UIElements/DoorLft.png",
		"resources/UIElements/DoorFwd.png",
		"resources/UIElements/DoorRgt.png",
	}
	d := &UIDoor{
		num: num,
	}
	var pic *ebiten.Image
	if num <= 2 {
		pic = t.GetImgByPath(doorPics[num])
	} else {
		pic, _ = ebiten.NewImage(w, h, ebiten.FilterLinear)
		_ = pic.Fill(color.Black)
	}
	d.DCInit(x, y, w, h, 1, NewSprite(pic))
	return d
}

func (d *UIDoor) Draw(screen *ebiten.Image, col color.Color) {
	if d.visible {
		d.DrawImg(screen)
	}
	//ebitenutil.DrawRect(screen, float64(d.x), float64(d.y), float64(d.w), float64(d.h), col)
}
