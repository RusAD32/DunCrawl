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

type PlayerStats struct {
	pl                                *Player
	hpX, hpY, hpW, hpH                int
	infoX, infoY                      int
	statusX, statusY                  int
	hpCol, effectCol, dmgCol, textCol color.Color
	dmgProcessing, healProcessing     string
}

func (p *PlayerStats) Draw(screen *ebiten.Image, font font.Face) {
	//THIS IS SPARTA!!!
	realW := p.hpW * p.pl.GetHP() / p.pl.GetMaxHP() / 2
	pethpW := p.hpW * p.pl.GetPet().GetHP() / p.pl.GetPet().GetMaxHP() / 2
	ebitenutil.DrawRect(screen, float64(p.hpX), float64(p.hpY), float64(realW), float64(p.hpH), p.hpCol)
	text.Draw(screen, fmt.Sprintf("%d/%d", p.pl.GetHP(), p.pl.GetMaxHP()), font, p.hpX, p.hpY+font.Metrics().Height.Ceil(), p.textCol)
	ebitenutil.DrawRect(screen, float64(p.hpX+p.hpW*11/20), float64(p.hpY), float64(pethpW), float64(p.hpH), p.hpCol)
	text.Draw(screen, fmt.Sprintf("%d/%d", p.pl.GetPet().GetHP(), p.pl.GetPet().GetMaxHP()), font, p.hpX+p.hpW*11/20, p.hpY+font.Metrics().Height.Ceil(), p.textCol)
	if p.healProcessing != "" {
		text.Draw(screen, p.healProcessing, font, p.infoX, p.infoY+font.Metrics().Height.Ceil(), p.hpCol)
	}
	if p.dmgProcessing != "" {
		text.Draw(screen, p.dmgProcessing, font, p.infoX, p.infoY+font.Metrics().Height.Ceil(), p.dmgCol)
	}
	for i, v := range *p.pl.GetEffects() {
		text.Draw(screen, v.GetInfo(), font, p.statusX/2-p.hpW/8*i, p.statusY, p.effectCol)
	}
	for i, v := range *p.pl.GetPet().GetEffects() {
		text.Draw(screen, v.GetInfo(), font, p.statusX-p.hpW/8*i, p.statusY, p.effectCol)
	}
}
