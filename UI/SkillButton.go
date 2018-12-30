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
	x, y, w, h             int
	font                   font.Face
	isSelf                 bool
	activeCol, disabledCol color.Color
	sk                     Skill
	DrawableImage
}

func (sb *SkillButton) Init(x, y, w, h int, sk Skill, activeCol, disabledCol color.Color, font font.Face) *SkillButton {
	sb.x = x
	sb.y = y
	sb.w = w
	sb.h = h
	sb.activeCol = activeCol
	sb.disabledCol = disabledCol
	sb.sk = sk
	sb.state = butInactive
	sb.font = font
	switch sk.(type) {
	case PlayerSelfSkill:
		sb.isSelf = true
	default:
		sb.isSelf = false
	}
	sb.pic = make([]*ebiten.Image, 2)
	sb.pic[butInactive], _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	_ = sb.pic[butInactive].Fill(disabledCol)
	sb.pic[butActive], _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	_ = sb.pic[butActive].Fill(activeCol)
	sb.opts = &ebiten.DrawImageOptions{}
	sb.opts.GeoM.Translate(float64(x), float64(y))
	return sb
}

func (sb *SkillButton) Draw(screen *ebiten.Image) {
	err := sb.DrawImg(screen)
	if err != nil {
		panic(err)
	}
	//ebitenutil.DrawRect(screen, float64(sb.x), float64(sb.y), float64(sb.w), float64(sb.h), sb.GetImage())
	text.Draw(screen, sb.sk.GetName(), sb.font, sb.x, sb.y+sb.font.Metrics().Height.Ceil()*2, color.Black)
}

func (sb *SkillButton) isClicked(mouseX, mouseY int) bool {
	return !(mouseX < sb.x || mouseX > sb.x+sb.w || mouseY < sb.y || mouseY > sb.y+sb.h ||
		!sb.isSelf && sb.sk.(PlayerDmgSkill).GetUses() == 0 || sb.state == butInactive)
}
