package UI

import (
	. "DunCrawl/Interfaces"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image/color"
)

type SkillIcon struct {
	font font.Face
	sk   SkillInfo
	x, y int
	DrawableImage
}

func NewSkillIcon(w, h int, sk SkillInfo, font font.Face, p *TexPreloader) *SkillIcon {
	sb := &SkillIcon{
		sk:   sk,
		font: font,
	}
	pic := p.GetImgByPath(sk.GetIconPath())
	sb.initImg(0, 0, w, h, 1, NewSprite(pic))
	return sb
}

func (sb *SkillIcon) Draw(screen *ebiten.Image) {
	sb.DrawImg(screen)
	//ebitenutil.DrawRect(screen, float64(sb.x), float64(sb.y), float64(sb.w), float64(sb.h), sb.GetImage())
	text.Draw(screen, sb.sk.GetName(), sb.font, sb.x, sb.y-sb.font.Metrics().Height.Ceil()*2, color.Black)
	text.Draw(screen, sb.sk.GetTarget().GetName(), sb.font, sb.x, sb.y-sb.font.Metrics().Height.Ceil(), color.Black)
}
