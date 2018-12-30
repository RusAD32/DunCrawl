package UI

import (
	. "DunCrawl/Interfaces"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image/color"
)

type SkillButton struct {
	x, y, w, h             int
	font                   font.Face
	isSelf, active         bool
	activeCol, disabledCol color.Color
	sk                     Skill
	picActive, picDisbled  *ebiten.Image
	opts                   *ebiten.DrawImageOptions
}

func (sb *SkillButton) Init(x, y, w, h int, sk Skill, activeCol, disabledCol color.Color, font font.Face) *SkillButton {
	sb.x = x
	sb.y = y
	sb.w = w
	sb.h = h
	sb.activeCol = activeCol
	sb.disabledCol = disabledCol
	sb.sk = sk
	sb.active = false
	sb.font = font
	switch sk.(type) {
	case PlayerSelfSkill:
		sb.isSelf = true
	default:
		sb.isSelf = false
	}
	sb.picDisbled, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	err := sb.picDisbled.Fill(disabledCol)
	if err != nil {
		panic(err)
	}
	sb.picActive, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	err = sb.picActive.Fill(activeCol)
	if err != nil {
		panic(err)
	}
	sb.opts = &ebiten.DrawImageOptions{}
	sb.opts.GeoM.Translate(float64(x), float64(y))
	return sb
}

func (sb *SkillButton) GetImage() *ebiten.Image {
	if sb.active {
		return sb.picActive
	}
	return sb.picDisbled
}

func (sb *SkillButton) Draw(screen *ebiten.Image) {
	err := screen.DrawImage(sb.GetImage(), sb.opts)
	if err != nil {
		panic(err)
	}
	//ebitenutil.DrawRect(screen, float64(sb.x), float64(sb.y), float64(sb.w), float64(sb.h), sb.GetImage())
	text.Draw(screen, sb.sk.GetName(), sb.font, sb.x, sb.y+sb.font.Metrics().Height.Ceil()*2, color.Black)
}

func (sb *SkillButton) isClicked(mouseX, mouseY int) bool {
	return !(mouseX < sb.x || mouseX > sb.x+sb.w || mouseY < sb.y || mouseY > sb.y+sb.h ||
		!sb.isSelf && sb.sk.(PlayerDmgSkill).GetUses() == 0 || !sb.active)
}
