package UI

import (
	. "DunCrawl/Interfaces"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image/color"
)

type SkillButton struct {
	x, y, w, h             int
	isSelf, active         bool
	activeCol, disabledCol color.Color
	sk                     Skill
}

func (sb *SkillButton) GetColor() color.Color {
	if sb.active {
		return sb.activeCol
	}
	return sb.disabledCol
}

func (sb *SkillButton) Draw(screen *ebiten.Image, font font.Face) {
	ebitenutil.DrawRect(screen, float64(sb.x), float64(sb.y), float64(sb.w), float64(sb.h), sb.GetColor())
	text.Draw(screen, sb.sk.GetName(), font, sb.x, sb.y+font.Metrics().Height.Ceil()*2, color.Black)
}

func (sb *SkillButton) isClicked(mouseX, mouseY int) bool {
	return !(mouseX < sb.x || mouseX > sb.x+sb.w || mouseY < sb.y || mouseY > sb.y+sb.h ||
		!sb.isSelf && sb.sk.(PlayerDmgSkill).GetUses() == 0 || !sb.active)
}
