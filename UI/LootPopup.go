package UI

import (
	. "DunCrawl/Interfaces"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image/color"
)

type CarriableState int

const (
	Taking CarriableState = iota
	Leaving
)

type LootIcon struct {
	DrawableImage
	loot Lootable
}

type CarriableIcon struct {
	DrawableClickable
	loot Stack
}

type LootPopup struct {
	loot    []*LootIcon
	goodies []*CarriableIcon
	button  *DrawableClickable
	DrawableClickable
}

func NewLootIcon(x, y, w, h int, loot *Lootable, font font.Face) *LootIcon {
	p := &LootIcon{}
	pic, _ := ebiten.NewImage(w, h, ebiten.FilterLinear)
	_ = pic.Fill(LightGreen)
	info1 := loot.GetName()
	info2 := fmt.Sprintf("(%dg)", loot.GetValue())
	text.Draw(pic, info1, font, 0, font.Metrics().Height.Ceil(), color.Black)
	text.Draw(pic, info2, font, 0, 2*font.Metrics().Height.Ceil(), color.Black)
	p.initImg(x, y, w, h, 1, NewSprite(pic))
	return p
}

func NewCarriableIcon(x, y, w, h int, loot Stack, font font.Face) *CarriableIcon {
	p := &CarriableIcon{}
	p.state = int(Taking)
	pic, _ := ebiten.NewImage(w, h, ebiten.FilterLinear)
	_ = pic.Fill(LightGreen)
	pic2, _ := ebiten.NewImage(w, h, ebiten.FilterLinear)
	_ = pic2.Fill(Gray)
	info1 := loot.GetName()
	info2 := fmt.Sprintf("(%d)", loot.GetAmount())
	text.Draw(pic, info1, font, 0, font.Metrics().Height.Ceil(), color.Black)
	text.Draw(pic, info2, font, 0, 2*font.Metrics().Height.Ceil(), color.Black)
	text.Draw(pic2, info1, font, 0, font.Metrics().Height.Ceil(), color.Black)
	text.Draw(pic2, info2, font, 0, 2*font.Metrics().Height.Ceil(), color.Black)
	p.DCInit(x, y, w, h, 1, NewSprite(pic), NewSprite(pic2))
	return p
}

func (b *LootPopup) onCarriableClicked(mouseX, mouseY int) bool {
	for _, v := range b.goodies {
		if v.isClicked(mouseX-b.x, mouseY-b.y) {
			v.state = (v.state + 1) % 2
			return true
		}
	}
	return false
}

func (b *LootPopup) isClicked(mouseX, mouseY int) bool {
	return b.button.isClicked(mouseX-b.x, mouseY-b.y)
}

func NewLootPopup(x, y, w, h int, font font.Face, loot []*Lootable, goodies []Stack) *LootPopup {
	p := &LootPopup{}
	iconW := h / 5
	iconH := h / 5
	lootIconX := h / 10
	lootIconY := h / 10
	lootIconOffs := iconW + h/10
	pic, _ := ebiten.NewImage(w, h, ebiten.FilterLinear)
	_ = pic.Fill(color.RGBA{R: 50, G: 50, B: 50, A: 255})
	p.loot = make([]*LootIcon, 0)
	for i, v := range loot {
		p.loot = append(p.loot, NewLootIcon(lootIconX+i*lootIconOffs, lootIconY, iconW, iconH, v, font))
	}
	carIconX := h / 10
	carIconY := h * 2 / 5
	p.goodies = make([]*CarriableIcon, 0)
	for i, v := range goodies {
		p.goodies = append(p.goodies, NewCarriableIcon(carIconX+i*lootIconOffs, carIconY, iconW, iconH, v, font))
	}
	butPic, _ := ebiten.NewImage(w/4, h/5, ebiten.FilterLinear)
	_ = butPic.Fill(color.RGBA{R: 177, G: 177, B: 177, A: 255})
	text.Draw(butPic, "Confirm", font, 0, font.Metrics().Height.Ceil(), color.Black)
	p.button = NewDrawableClickable(w*2/5, h*3/4, w/4, h/5, 1, NewSprite(butPic))
	p.DCInit(x, y, w, h, 1, NewSprite(pic))
	return p
}

func (p *LootPopup) Draw(screen *ebiten.Image) {
	for _, v := range p.loot {
		v.DrawImg(p.pic[0].GetCurrentFrame())
	}
	for _, v := range p.goodies {
		v.DrawImg(p.pic[0].GetCurrentFrame())
	}
	p.button.DrawImg(p.pic[0].GetCurrentFrame())
	p.DrawImg(screen)
}
