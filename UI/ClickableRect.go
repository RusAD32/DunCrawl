package UI

type ClickableRect struct {
	x, y, w, h int
}

func (c *ClickableRect) initRect(x, y, w, h int) {
	c.x = x
	c.y = y
	c.w = w
	c.h = h
}

func (c *ClickableRect) isClicked(mouseX, mouseY int) bool {
	return !(mouseX < int(c.x) || mouseX > int(c.x+c.w) || mouseY < int(c.y) || mouseY > int(c.y+c.h))
}
