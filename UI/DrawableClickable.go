package UI

import (
	"image/color"
)

type DrawableClickable struct {
	ClickableRect
	DrawableImage
}

func (c *DrawableClickable) DCInit(x, y, w, h, length int, col ...color.Color) *DrawableClickable {
	c.initRect(x, y, w, h)
	c.initImg(x, y, w, h, length, col...)
	return c
}
