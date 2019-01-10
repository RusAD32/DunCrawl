package UI

type DrawableClickable struct {
	ClickableRect
	DrawableImage
}

func NewDrawableClickable(x, y, w, h, length int, imgs ...*Sprite) *DrawableClickable {
	return (&DrawableClickable{}).DCInit(x, y, w, h, length, imgs...)
}

func (c *DrawableClickable) DCInit(x, y, w, h, length int, imgs ...*Sprite) *DrawableClickable {
	c.initRect(x, y, w, h)
	c.initImg(x, y, w, h, length, imgs...)
	return c
}
