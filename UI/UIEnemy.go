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

const (
	enemyDefault = iota
	enemyAttacking
	enemyAttacked
	enemyDead
)

type UIEnemy struct {
	DrawableClickable
	col        color.Color
	enemy      *Enemy
	isTargeted bool
	skillUsed  Skill
}

func (e *UIEnemy) Init(x, y, w, h int, colDef, colAttacking, colAttacked, colDead color.Color, enemy *Enemy) *UIEnemy {
	e.col = colDef
	e.enemy = enemy
	e.DCInit(x, y, w, h, 4, colDef, colAttacking, colAttacked, colDead)
	return e
}

func (e *UIEnemy) Draw(screen *ebiten.Image, font font.Face) {
	if !e.enemy.IsAlive() {
		e.state = enemyDead
	}
	e.DrawImg(screen)
	//ebitenutil.DrawRect(screen, float64(e.x), float64(e.y), float64(e.w), float64(e.h), e.col)
	text.Draw(screen,
		fmt.Sprintf("%s\n%d/%d\n", e.enemy.GetName(), e.enemy.GetHP(), e.enemy.GetMaxHP()),
		font,
		e.x,
		e.y-font.Metrics().Height.Ceil(),
		e.col,
	)
	if e.isTargeted {
		ebitenutil.DrawRect(screen, float64(e.x), float64(e.y)+float64(e.h)*1.2, float64(e.w), 10, e.col)
	}
	if e.skillUsed != nil {
		text.Draw(screen,
			e.skillUsed.GetName(),
			font,
			e.x,
			e.y+font.Metrics().Height.Ceil(),
			color.Black,
		)
	}
	for _, v := range *e.enemy.GetEffects() {
		text.Draw(screen, v.GetInfo(), font, e.x, e.y+e.h*11/10, e.col)
	}
}
