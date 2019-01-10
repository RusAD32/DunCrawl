package UI

import (
	. "DunCrawl/Interfaces"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image/color"
)

const (
	butInactive = iota
	butActive
)

type SkillButton struct {
	font   font.Face
	isSelf bool
	sk     Skill
	DrawableClickable
}

func NewSkillButton(x, y, w, h int, sk Skill, activeCol, disabledCol color.Color, font font.Face) *SkillButton {
	sb := &SkillButton{
		sk:   sk,
		font: font,
	}
	switch sk.(type) {
	case PlayerSelfSkill:
		sb.isSelf = true
	default:
		sb.isSelf = false
	}
	sb.DCInit(x, y, w, h, 2, disabledCol, activeCol)
	text.Draw(sb.pic[butActive], sb.sk.GetName(), sb.font, 0, sb.font.Metrics().Height.Ceil()*2, color.Black)
	text.Draw(sb.pic[butInactive], sb.sk.GetName(), sb.font, 0, sb.font.Metrics().Height.Ceil()*2, color.Black)
	return sb
}

func (sb *SkillButton) Draw(screen *ebiten.Image) {
	sb.DrawImg(screen)
	//ebitenutil.DrawRect(screen, float64(sb.x), float64(sb.y), float64(sb.w), float64(sb.h), sb.GetImage())
}
