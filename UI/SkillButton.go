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
	x, y, w, h     float64
	isSelf, active bool
	sk             Skill
}

func (sb *SkillButton) Draw(screen *ebiten.Image, font font.Face) {
	col := color.RGBA{200, 200, 255, 255}
	if !sb.active {
		col = color.RGBA{200, 200, 200, 255}
	}
	ebitenutil.DrawRect(screen, sb.x, sb.y, sb.w, sb.h, col)
	text.Draw(screen, sb.sk.GetName(), font, int(sb.x), int(sb.y)+font.Metrics().Height.Ceil()*2, color.Black)
}

func (sb *SkillButton) isClicked(mouseX, mouseY int) bool {
	return !(mouseX < int(sb.x) || mouseX > int(sb.x+sb.w) || mouseY < int(sb.y) || mouseY > int(sb.y+sb.h) || !sb.isSelf && sb.sk.(PlayerDmgSkill).GetUses() == 0)
}
