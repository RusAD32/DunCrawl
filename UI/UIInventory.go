package UI

import (
	. "DunCrawl/Interfaces"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image/color"
)

type InventoryIcon struct {
	DrawableClickable
	loot Stack
}

type UIInventory struct {
	items []*InventoryIcon
	DrawableClickable
	p *Player
}

func NewInventoryIcon(x, y, w, h int, loot Stack, font font.Face) *InventoryIcon {
	p := &InventoryIcon{}
	pic, _ := ebiten.NewImage(w, h, ebiten.FilterLinear)
	if loot != nil {
		_ = pic.Fill(LightGreen)
		info1 := loot.GetName()
		info2 := fmt.Sprintf("(%d)", loot.GetAmount())
		text.Draw(pic, info1, font, 0, font.Metrics().Height.Ceil(), color.Black)
		text.Draw(pic, info2, font, 0, 2*font.Metrics().Height.Ceil(), color.Black)
	} else {
		_ = pic.Fill(LightGray)
	}
	p.DCInit(x, y, w, h, 1, NewSprite(pic))
	return p
}

func (b *UIInventory) onIconClicked(mouseX, mouseY int) bool {
	for _, v := range b.items {
		if v.isClicked(mouseX-b.x, mouseY-b.y) {
			v.loot.GetItem().Use(b.p)
			return true
		}
	}
	return false
}

func NewUIInventory(x, y, w, h int, font font.Face, items []Stack) *UIInventory {
	// that's crapcode, but for now to refresh the inventory on the screen you need to create a new object like that
	p := &UIInventory{}
	iconW := h / 3
	iconH := h / 3
	lootIconX := h / 10
	lootIconY := h / 10
	lootIconOffs := iconW + h/10
	pic, _ := ebiten.NewImage(w, h, ebiten.FilterLinear)
	_ = pic.Fill(color.RGBA{R: 50, G: 50, B: 50, A: 255})
	p.items = make([]*InventoryIcon, 0)
	for i, v := range items {
		p.items = append(p.items, NewInventoryIcon(lootIconX+(i%4)*lootIconOffs, lootIconY+(i/4)*lootIconOffs, iconW, iconH, v, font))
	}
	p.DCInit(x, y, w, h, 1, NewSprite(pic))
	return p
}

func (p *UIInventory) Draw(screen *ebiten.Image) {
	for _, v := range p.items {
		v.DrawImg(p.pic[0].GetCurrentFrame())
	}
	p.DrawImg(screen)
}
