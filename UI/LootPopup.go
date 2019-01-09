package UI

import (
	. "DunCrawl/Interfaces"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image/color"
)

type LootIcon struct {
	DrawableImage
	loot Lootable
}

type CarriableIcon struct {
	DrawableClickable
	loot    Stack
	leaving bool
}

type LootPopup struct {
	loot    []*LootIcon
	goodies []*CarriableIcon
	button  *DrawableClickable
	DrawableClickable
}

func (p *LootIcon) Init(x, y, w, h int, loot Lootable, font font.Face) *LootIcon {
	p.initImg(x, y, w, h, 1, Gray)
	info1 := loot.GetName()
	info2 := fmt.Sprintf("(%dg)", loot.GetValue())
	text.Draw(p.pic[0], info1, font, 0, font.Metrics().Height.Ceil(), color.Black)
	text.Draw(p.pic[0], info2, font, 0, 2*font.Metrics().Height.Ceil(), color.Black)
	return p
}

func (p *CarriableIcon) Init(x, y, w, h int, loot Stack, font font.Face) *CarriableIcon {
	p.DCInit(x, y, w, h, 1, Gray)
	info1 := loot.GetName()
	info2 := fmt.Sprintf("(%d)", loot.GetAmount())
	text.Draw(p.pic[0], info1, font, 0, font.Metrics().Height.Ceil(), color.Black)
	text.Draw(p.pic[0], info2, font, 0, 2*font.Metrics().Height.Ceil(), color.Black)
	return p
}

func (b *LootPopup) isClicked(mouseX, mouseY int) bool {
	return b.button.isClicked(mouseX-b.x, mouseY-b.y)
}

func (p *LootPopup) Init(x, y, w, h int, font font.Face, loot []Lootable, goodies []Stack) *LootPopup {
	iconW := h / 5
	iconH := h / 5
	lootIconX := h / 10
	lootIconY := h / 10
	lootIconOffs := iconW + h/10
	p.DCInit(x, y, w, h, 1, color.RGBA{R: 50, G: 50, B: 50, A: 255})
	p.loot = make([]*LootIcon, 0)
	for i, v := range loot {
		p.loot = append(p.loot, new(LootIcon).Init(lootIconX+i*lootIconOffs, lootIconY, iconW, iconH, v, font))
	}
	carIconX := h / 10
	carIconY := h * 2 / 5
	p.goodies = make([]*CarriableIcon, 0)
	for i, v := range goodies {
		p.goodies = append(p.goodies, new(CarriableIcon).Init(carIconX+i*lootIconOffs, carIconY, iconW, iconH, v, font))
	}
	p.button = new(DrawableClickable).DCInit(w*2/5, h*3/4, w/4, h/5, 1, color.RGBA{R: 177, G: 177, B: 177, A: 255})
	text.Draw(p.button.pic[0], "Confirm", font, 0, w/8-font.Metrics().Height.Ceil(), color.Black)
	return p
}

func (p *LootPopup) Draw(screen *ebiten.Image) {
	for _, v := range p.loot {
		v.DrawImg(p.pic[0])
	}
	for _, v := range p.goodies {
		v.DrawImg(p.pic[0])
	}
	p.button.DrawImg(p.pic[0])
	p.DrawImg(screen)
}
