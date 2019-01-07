package UI

import (
	. "DunCrawl/Interfaces"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"golang.org/x/image/font"
	"image/color"
)

type LootIcon struct {
	DrawableImage
	loot Lootable
}

type CarriableIcon struct {
	DrawableImage
	x, y, w, h int // TODO вынести это в какой-нибудь rectangle
	loot       Stack
	leaving    bool
}

type ConfirmLootButton struct {
	x, y, w, h int
	DrawableImage
}

type LootPopup struct {
	loot       []*LootIcon
	goodies    []*CarriableIcon
	button     *ConfirmLootButton
	x, y, w, h int
	DrawableImage
}

func (p *LootIcon) Init(x, y, w, h int, loot Lootable, font font.Face) *LootIcon {
	p.opts = &ebiten.DrawImageOptions{}
	p.pic = make([]*ebiten.Image, 1)
	p.opts.GeoM.Translate(float64(x), float64(y))
	p.loot = loot
	pic, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)
	_ = pic.Fill(Gray)
	//text.Draw(pic, fmt.Sprintf("%s (%dg)", loot.GetName(), loot.GetValue()), font, 0, 0, color.Black)
	ebitenutil.DebugPrint(pic, fmt.Sprintf("%s\n(%dg)", loot.GetName(), loot.GetValue()))
	p.pic[0] = pic
	return p
}

func (p *CarriableIcon) Init(x, y, w, h int, loot Stack, font font.Face) *CarriableIcon {
	p.opts = &ebiten.DrawImageOptions{}
	p.pic = make([]*ebiten.Image, 1)
	p.opts.GeoM.Translate(float64(x), float64(y))
	p.loot = loot
	pic, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)
	_ = pic.Fill(Gray)
	//text.Draw(pic, fmt.Sprintf("%s (%d)", loot.GetName(), loot.GetAmount()), font, 0, 0, color.Black)
	_ = ebitenutil.DebugPrint(pic, fmt.Sprintf("%s\n(%d)"))
	p.pic[0] = pic
	return p
}

func (b *ConfirmLootButton) Init(x, y, w, h int, font font.Face) *ConfirmLootButton {
	b.x = x
	b.y = y
	b.w = w
	b.h = h
	b.pic = make([]*ebiten.Image, 1)
	b.pic[0], _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	b.opts = &ebiten.DrawImageOptions{}
	b.opts.GeoM.Translate(float64(x), float64(y))
	_ = b.pic[0].Fill(color.RGBA{177, 177, 177, 255})
	_ = ebitenutil.DebugPrint(b.pic[0], "Confirm") // TODO понять, wtf с text.Draw()
	//text.Draw(b.pic[0], "Confirm", font, 0, 0, color.Black)
	return b
}

func (b *LootPopup) IsClicked(mouseX, mouseY int) bool {
	return !(mouseX < b.x+b.button.x || mouseX > b.x+b.button.x+b.button.w || mouseY < b.y+b.button.y || mouseY > b.y+b.button.y+b.h)
}

func (p *LootPopup) Init(x, y, w, h int, font font.Face, loot []Lootable, goodies []Stack) *LootPopup {
	p.x = x
	p.y = y
	p.w = w
	p.h = h
	iconW := h / 5
	iconH := h / 5
	lootIconX := h / 10
	lootIconY := h / 10
	lootIconOffs := iconW + h/10
	p.opts = &ebiten.DrawImageOptions{}
	p.opts.GeoM.Translate(float64(x), float64(y))
	p.pic = make([]*ebiten.Image, 1)
	p.pic[0], _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	_ = p.pic[0].Fill(color.RGBA{50, 50, 50, 255})
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
	p.button = new(ConfirmLootButton).Init(w*2/5, h*3/4, w/4, h/5, font)
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
