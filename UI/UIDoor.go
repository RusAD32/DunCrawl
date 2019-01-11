package UI

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
)

type UIDoor struct {
	DrawableClickable
	num     int
	visible bool
}

func NewUIDoor(x, y, w, h, num int) *UIDoor {
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
		var err error
		pic, _, err = ebitenutil.NewImageFromFile(doorPics[num], ebiten.FilterLinear)
		if err != nil {
			panic(err)
		}
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
