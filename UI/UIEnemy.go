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
	isTargeted bool
	skillUsed  Skill
	pic        *ebiten.Image
	opts       *ebiten.DrawImageOptions
}

func (e *UIEnemy) Init(x, y, w, h int, col color.Color, enemy *Enemy) *UIEnemy {
	e.x = x
	e.y = y
	e.w = w
	e.h = h
	e.col = col
	e.enemy = enemy
	e.pic, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	err := e.pic.Fill(col)
	if err != nil {
		panic(err)
	}
	e.opts = &ebiten.DrawImageOptions{}
	e.opts.GeoM.Translate(float64(x), float64(y))
	return e
}

func (e *UIEnemy) isClicked(mouseX, mouseY int) bool {
	return !(mouseX < int(e.x) || mouseX > int(e.x+e.w) || mouseY < int(e.y) || mouseY > int(e.y+e.h))
}

func (e *UIEnemy) Draw(screen *ebiten.Image, font font.Face) {
	err := screen.DrawImage(e.pic, e.opts)
	if err != nil {
		panic(err)
	}
	//ebitenutil.DrawRect(screen, float64(e.x), float64(e.y), float64(e.w), float64(e.h), e.col)
	text.Draw(screen,
		fmt.Sprintf("%s\n%d/%d\n", e.enemy.GetName(), e.enemy.GetCurHP(), e.enemy.GetMaxHP()),
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
}
