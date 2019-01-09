package UI

type ClickableRect struct {
	x, y, w, h int
}

func (c *ClickableRect) isClicked(mouseX, mouseY int) bool {
	return !(mouseX < int(c.x) || mouseX > int(c.x+c.w) || mouseY < int(c.y) || mouseY > int(c.y+c.h))
}
