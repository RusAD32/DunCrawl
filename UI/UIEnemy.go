package UI

import (
	. "DunCrawl/Interfaces"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image/color"
)

type UIEnemy struct {
	x, y, w, h int
	col        color.Color
	enemy      *Enemy
}

func (e *UIEnemy) Draw(screen *ebiten.Image, font font.Face) {
	ebitenutil.DrawRect(screen, float64(e.x), float64(e.y), float64(e.w), float64(e.h), e.col)
	text.Draw(screen,
		fmt.Sprintf("%s\n%d/%d\n", e.enemy.GetName(), e.enemy.GetCurHP(), e.enemy.GetMaxHP()),
		font,
		e.x,
		e.y-font.Metrics().Height.Ceil(),
		e.col,
	)
}
