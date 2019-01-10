package UI

import (
	. "DunCrawl/Interfaces"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	_ "image/png"
	"runtime"
)

func (g *UIGame) Update() {
	if g.cd > 0 {
		g.cd--
		return
	}
	switch g.l.GetState() {
	case Roam:
		{
			for _, v := range getNewClicks() {
				if g.Light(v[0], v[1]) {
					return
				}
				if g.loot != nil {
					if g.loot.isClicked(v[0], v[1]) {
						g.loot = nil // TODO давать игроку то, что он нашел, в конце концов
					}
					return
				}
				if g.chest != nil {
					for _, click := range getNewClicks() {
						if g.chest.isClicked(click[0], click[1]) {
							loot, goodies := g.l.UnlockChest()
							g.loot = NewLootPopup(g.w/3, g.h/3, g.w/3, g.h/3, g.font, loot, goodies)
							g.chest = nil
							return
						}
					}
				}
				nextDoor := g.doorClicked(v[0], v[1])
				if nextDoor != -1 {
					g.chest = nil
					if g.l.GotoRoom(Direction(nextDoor)) {
						g.prepareForFight()
					} else {
						loot, goodies := g.l.GetCurrentRoom().GetValues()
						if len(loot) != 0 || len(goodies) != 0 {
							g.loot = NewLootPopup(g.w/3, g.h/3, g.w/3, g.h/3, g.font, loot, goodies)
						}
					}
					if g.l.GetCurrentRoom().HasChest() {
						pic, _, err := ebitenutil.NewImageFromFile("./resources/UIElements/chest_t.png", ebiten.FilterLinear)
						if err != nil {
							panic(err)
						}
						g.chest = NewDrawableClickable(g.w/3, g.h/3, g.w/3, g.w/3, 1, NewSprite(pic))
					}
					g.updateDoors()
					return
				}
			}
		}
	case Fight:
		{
			switch g.l.GetCurrentRoom().FightState {
			case AwaitingSelfSkill:
				g.submitSelfSkill()
			case AwaitingDmgSkill:
				g.submitDmgSkill()
			case ResolvingSkills:
				g.resolveSkill()
			case FightEnd:
				loot, values := g.l.GetValues() // TODO display this on screen
				g.loot = NewLootPopup(g.w/3, g.h/3, g.w/3, g.h/3, g.font, loot, values)
				g.curEnemies = make([]*UIEnemy, 0)
				g.selfSkButs = make([]*SkillButton, 0)
				g.dmgSkButs = make([]*SkillButton, 0)
			}
		}
	}
}

func (g *UIGame) Draw(screen *ebiten.Image) {
	switch g.l.GetState() {
	case Roam:
		{
			//this is the most memory-greedy function
			DrawLabyrinth(screen, g.l, g.consts.labXPos, g.consts.labYPos, g.consts.labW, g.consts.labH, color.Black)
			for _, v := range g.currentDoors {
				v.Draw(screen, color.Black)
			}
			if g.chest != nil {
				g.chest.DrawImg(screen)
			}
			if g.loot != nil {
				g.loot.Draw(screen)
			}
			g.light.DrawImg(screen)
		}
	case Fight:
		{
			for _, v := range g.curEnemies {
				v.Draw(screen, g.font)
			}
			for _, v := range append(g.selfSkButs, g.dmgSkButs...) {
				v.Draw(screen)
			}
			g.pl.Draw(screen, g.font)
			g.queue.Draw(screen)
			if g.resolvingSk != nil {
				g.resolvingSk.Draw(screen)
			}
		}
	}
	//ebitenutil.DebugPrintAt(screen, PrintMemUsage(), 0, 300)

}

func PrintMemUsage() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	res := fmt.Sprintf("Alloc = %v MiB\n", bToMb(m.Alloc))
	res += fmt.Sprintf("TotalAlloc = %v MiB\n", bToMb(m.TotalAlloc))
	res += fmt.Sprintf("Sys = %v MiB\n", bToMb(m.Sys))
	res += fmt.Sprintf("NumGC = %v\n\n", m.NumGC)
	return res
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
