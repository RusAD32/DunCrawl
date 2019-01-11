package UI

import (
	. "DunCrawl/Interfaces"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
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
	activePic, _, err := ebitenutil.NewImageFromFile(sk.GetIconPath(), ebiten.FilterLinear)
	if err != nil {
		panic(err)
	}
	disabledPic, _, err := ebitenutil.NewImageFromFile(sk.GetIconPath(), ebiten.FilterLinear)
	if err != nil {
		panic(err)
	}
	//text.Draw(disabledPic, sb.sk.GetName(), sb.font, 0, sb.font.Metrics().Height.Ceil(), color.Black)
	//text.Draw(activePic, sb.sk.GetName(), sb.font, 0, sb.font.Metrics().Height.Ceil(), color.Black)
	sb.DCInit(x, y, w, h, 2, NewSprite(disabledPic), NewSprite(activePic))
	sb.pic[0].opts.ColorM.Scale(0.5, 0.5, 0.5, 1)
	return sb
}

func (sb *SkillButton) Draw(screen *ebiten.Image) {
	sb.DrawImg(screen)
	//ebitenutil.DrawRect(screen, float64(sb.x), float64(sb.y), float64(sb.w), float64(sb.h), sb.GetImage())
}
