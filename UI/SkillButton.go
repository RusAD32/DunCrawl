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
	disabledPic, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)
	_ = disabledPic.Fill(disabledCol)
	text.Draw(disabledPic, sb.sk.GetName(), sb.font, 0, sb.font.Metrics().Height.Ceil()*2, color.Black)
	activePic, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)
	_ = activePic.Fill(activeCol)
	text.Draw(activePic, sb.sk.GetName(), sb.font, 0, sb.font.Metrics().Height.Ceil()*2, color.Black)
	sb.DCInit(x, y, w, h, 2, NewSprite(disabledPic), NewSprite(activePic))
	return sb
}

func (sb *SkillButton) Draw(screen *ebiten.Image) {
	sb.DrawImg(screen)
	//ebitenutil.DrawRect(screen, float64(sb.x), float64(sb.y), float64(sb.w), float64(sb.h), sb.GetImage())
}
