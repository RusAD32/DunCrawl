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

func loadSpriteFromPics(paths []string) *Sprite {
	pics := make([]*ebiten.Image, 0)
	for _, v := range paths {
		pic, _, err := ebitenutil.NewImageFromFile(v, ebiten.FilterDefault)
		if err != nil {
			panic(err)
		}
		pics = append(pics, pic)
	}
	return NewSprite(pics...)
}

func NewUIEnemy(x, y, w, h int, colDef, colAttacking, colAttacked, colDead color.Color, enemy *Enemy) *UIEnemy {
	e := &UIEnemy{
		col:   colDef,
		enemy: enemy,
	}
	spriteIdle := loadSpriteFromPics(enemy.IdleImgsPath())
	spriteSkill := loadSpriteFromPics(enemy.SkillImgsPath())
	spriteAttacked := loadSpriteFromPics(enemy.AttackedImgsPath())
	spriteDead := loadSpriteFromPics(enemy.DeadImgsPath())
	spriteDead.noLoop = true
	picDef, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)
	_ = picDef.Fill(colDef)
	picAttacking, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)
	_ = picAttacking.Fill(colAttacking)
	picAttacked, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)
	_ = picAttacked.Fill(colAttacked)
	picDead, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)
	_ = picDead.Fill(colDead)
	e.DCInit(x, y, w, h, 4, spriteIdle, spriteSkill, spriteAttacked, spriteDead)
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

func (e *UIEnemy) GetCurAnimLen() int {
	return e.pic[enemyAttacked].GetAnimationLength()
}
