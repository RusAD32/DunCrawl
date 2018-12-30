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

func (sb *SkillIcon) Init(w, h int, sk SkillInfo, col color.Color, font font.Face) *SkillIcon {
	sb.sk = sk
	sb.font = font
	sb.pic = make([]*ebiten.Image, 1)
	sb.pic[0], _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	_ = sb.pic[0].Fill(col)
	sb.opts = &ebiten.DrawImageOptions{}
	return sb
}

func (sb *SkillIcon) Draw(screen *ebiten.Image) {
	sb.DrawImg(screen)
	//ebitenutil.DrawRect(screen, float64(sb.x), float64(sb.y), float64(sb.w), float64(sb.h), sb.GetImage())
	text.Draw(screen, sb.sk.GetName(), sb.font, sb.x, sb.y+sb.font.Metrics().Height.Ceil()*3/2, color.Black)
	text.Draw(screen, sb.sk.GetTarget().GetName(), sb.font, sb.x, sb.y+sb.font.Metrics().Height.Ceil()*3, color.Black)
}
