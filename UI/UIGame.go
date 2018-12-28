package UI

import (
	. "DunCrawl/Interfaces"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"golang.org/x/image/font"
	"image/color"
)

type GameState int

const (
	Roam GameState = iota
	Fight
)

var (
	Red        = color.RGBA{R: 255, A: 255}
	Green      = color.RGBA{G: 255, A: 255}
	Blue       = color.RGBA{B: 255, A: 255}
	Firebrick  = color.RGBA{R: 205, G: 38, B: 38, A: 255}
	OrangeRed  = color.RGBA{R: 255, G: 69, A: 255}
	Violet     = color.RGBA{R: 208, G: 32, B: 144, A: 255}
	Gray       = color.RGBA{R: 200, G: 200, B: 200, A: 255}
	LightGreen = color.RGBA{R: 200, G: 255, B: 200, A: 255}
	LightBlue  = color.RGBA{R: 200, G: 200, B: 255, A: 255}
)

type UIGame struct {
	w            int
	h            int
	l            *Labyrinth
	state        GameState
	font         font.Face
	currentDoors []*UIDoor
	curEnemies   []*UIEnemy
	selfSkButs   []*SkillButton
	dmgSkButs    []*SkillButton
	enemyNums    map[*Enemy]int
	pl           *PlayerStats
	cd           int
	consts
}

type consts struct {
	// constants regarding position of different elements of UI
	labXPos, labYPos, labW, labH                             int
	doorX, doorXOff, doorY, doorH, doorW                     int
	backdoorX, backdoorY, backdoorH, backdoorW               int
	enemyX, enemyXOff, enemyY, enemyH, enemyW                int
	hpX, hpY, hpW, hpH, infoX, infoY, statusX, statusY       int
	selfSkButX, skButXOff, skButY, skButW, skButH, dmgSkButX int
	// end of constant declaration
}

func (g *UIGame) Init(l *Labyrinth, w, h int) {
	// Constants go into UIGame object and are united here!
	g.l = l
	g.w = w
	g.h = h

	g.consts.labXPos = 5
	g.consts.labYPos = 5
	g.consts.labW = w / 5
	g.consts.labH = h / 5

	g.consts.doorX = w * 2 / 10
	g.consts.doorXOff = w / 4
	g.consts.doorY = h * 2 / 10
	g.consts.doorW = w * 15 / 100
	g.consts.doorH = h / 2

	g.consts.enemyX = w / 16
	g.consts.enemyXOff = w / 4
	g.consts.enemyY = h / 4
	g.consts.enemyW = w / 8
	g.consts.enemyH = h / 4

	g.consts.hpX = w / 10
	g.consts.hpY = h * 8 / 10
	g.consts.hpW = w * 8 / 10
	g.consts.hpH = h / 16
	g.consts.infoX = w / 3
	g.consts.infoY = h * 9 / 10
	g.consts.statusX = w * 8 / 10
	g.consts.statusY = h * 9 / 10

	g.consts.backdoorX = w * 2 / 10
	g.consts.backdoorY = h * 8 / 10
	g.consts.backdoorW = w * 66 / 100
	g.consts.backdoorH = h / 10

	g.consts.selfSkButX = w / 16
	g.consts.dmgSkButX = w/2 + w/16
	g.consts.skButY = h * 66 / 100
	g.consts.skButXOff = w / 8
	g.consts.skButW = w / 10
	g.consts.skButH = h / 10

	// end of constant declaration
	g.updateDoors()
	g.state = Roam
	g.curEnemies = make([]*UIEnemy, 0)
	fontRaw, err := LoadResource("Roboto-Regular.ttf")
	if err != nil {
		panic(err)
	}
	fontData, err := truetype.Parse(fontRaw)
	g.font = truetype.NewFace(fontData, &truetype.Options{})
	g.selfSkButs = make([]*SkillButton, 0)
	g.dmgSkButs = make([]*SkillButton, 0)
	g.enemyNums = make(map[*Enemy]int)
	plst := PlayerStats{
		g.l.GetPlayer(),
		g.consts.hpX,
		g.consts.hpY,
		g.consts.hpW,
		g.consts.hpH,
		g.consts.infoX,
		g.consts.infoY,
		g.consts.statusX,
		g.consts.statusY,
		Green,
		Blue,
		Red,
		color.Black,
		"", "",
	}
	g.pl = &plst
}

func (g *UIGame) doorClicked(mouseX, mouseY int) int {
	for _, v := range g.currentDoors {
		if v.isClicked(mouseX, mouseY) {
			return v.num
		}
	}
	return -1
}

func (g *UIGame) selfSkillButtonClicked(mouseX, mouseY int) int {
	for i, v := range g.selfSkButs {
		if v.isClicked(mouseX, mouseY) {
			return i
		}
	}
	return -1
}

func (g *UIGame) dmgSkillButtonClicked(mouseX, mouseY int) int {
	for i, v := range g.dmgSkButs {
		if v.isClicked(mouseX, mouseY) {
			return i
		}
	}
	return -1
}

func (g *UIGame) Draw(screen *ebiten.Image) {
	switch g.state {
	case Roam:
		{
			DrawLabyrinth(screen, g.l, g.consts.labXPos, g.consts.labYPos, g.consts.labW, g.consts.labH, color.Black)
			for _, v := range g.currentDoors {
				v.Draw(screen, color.Black)
			}
		}
	case Fight:
		{
			for _, v := range g.curEnemies {
				v.Draw(screen, g.font)
			}
			for _, v := range append(g.selfSkButs, g.dmgSkButs...) {
				v.Draw(screen, g.font)
			}
			g.pl.Draw(screen, g.font)
		}
	}

}

func (g *UIGame) Update() {
	if g.cd > 0 {
		g.cd--
		return
	}
	switch g.state {
	case Roam:
		{
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
				nextDoor := g.doorClicked(ebiten.CursorPosition())
				go g.l.GoToRoom(Direction(nextDoor))
				event := <-g.l.GetEventsChan()
				if event == FightEvent {
					ens := g.l.GetCurrentRoom().GetEnemies()
					for i, v := range ens {
						enemy := UIEnemy{
							g.consts.enemyXOff*i + g.consts.enemyX,
							g.consts.enemyY,
							g.consts.enemyW,
							g.consts.enemyH,
							Violet,
							v,
							false,
							nil,
						}
						g.enemyNums[v] = i
						g.curEnemies = append(g.curEnemies, &enemy)
					}
					for i, v := range g.l.GetPlayer().GetSelfSkillList() {
						button := SkillButton{
							g.consts.selfSkButX + i*g.consts.skButXOff,
							g.consts.skButY,
							g.consts.skButW,
							g.consts.skButH,
							true,
							true,
							LightGreen,
							Gray,
							v,
						}
						g.selfSkButs = append(g.selfSkButs, &button)
					}
					for i, v := range g.l.GetPlayer().GetDmgSkillList() {
						button := SkillButton{
							g.consts.dmgSkButX + i*g.consts.skButXOff,
							g.consts.skButY,
							g.consts.skButW,
							g.consts.skButH,
							false,
							false,
							LightBlue,
							Gray,
							v,
						}
						g.dmgSkButs = append(g.dmgSkButs, &button)
					}
					g.state = Fight
					g.l.GetCurrentRoom().AtTurnStart()
				}
				g.updateDoors()
			}
		}
	case Fight:
		{
			switch g.l.GetCurrentRoom().FightState {
			case AwaitingSelfSkill:
				for _, v := range g.selfSkButs {
					v.active = true
				}
				for _, v := range g.curEnemies {
					v.col = Violet
				}
				g.pl.dmgProcessing = ""
				g.pl.healProcessing = ""
				if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
					butNum := g.selfSkillButtonClicked(ebiten.CursorPosition())
					if butNum >= 0 {
						g.l.GetCurrentRoom().SubmitSelfSkill(g.selfSkButs[butNum].sk.(PlayerSelfSkill))
						g.curEnemies[0].isTargeted = true
						for _, v := range g.selfSkButs {
							v.active = false
						}
						for _, v := range g.dmgSkButs {
							v.active = v.sk.(PlayerDmgSkill).GetUses() > 0
						}
						return
					}
				}
			case AwaitingDmgSkill:
				if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
					skNum := g.dmgSkillButtonClicked(ebiten.CursorPosition())
					if skNum >= 0 {
						var curEn *UIEnemy
						for _, v := range g.curEnemies {
							if v.isTargeted {
								curEn = v
							}
						}
						skill := g.dmgSkButs[skNum].sk.(PlayerDmgSkill)
						skill.SetTarget(curEn.enemy)
						curEn.skillUsed = skill
						for _, v := range g.curEnemies {
							if !v.isTargeted && v.skillUsed == nil {
								v.isTargeted = true
								break
							}
						}
						curEn.isTargeted = false
						g.l.GetCurrentRoom().SubmitDmgSkill(skill)
						g.dmgSkButs[skNum].active = skill.GetUses() != 0
						return
					}
					for _, v := range g.curEnemies {
						if v.isClicked(ebiten.CursorPosition()) && v.skillUsed == nil {
							for _, v := range g.curEnemies {
								v.isTargeted = false
							}
							v.isTargeted = true
							return
						}
					}
				}
			case ResolvingSkills:
				for _, v := range append(g.selfSkButs, g.dmgSkButs...) {
					v.active = false
				}
				for _, v := range g.curEnemies {
					v.col = Violet
				}
				sk := g.l.GetCurrentRoom().GetNextSkillUsed()
				target := sk.GetTarget()
				switch sk.(type) {
				case PlayerDmgSkill:
					{
						en := g.curEnemies[g.enemyNums[target.(*Enemy)]]
						en.col = Firebrick
						en.skillUsed = nil
						g.cd = 60
						g.pl.dmgProcessing = sk.GetRes()
						g.pl.healProcessing = ""
					}
				case EnemySkill:
					{
						en := g.curEnemies[g.enemyNums[sk.GetWielder().(*Enemy)]]
						en.col = OrangeRed
						g.cd = 60
						g.pl.dmgProcessing = sk.GetRes()
						g.pl.healProcessing = ""
					}
				case PlayerSelfSkill:
					{
						g.pl.dmgProcessing = ""
						g.pl.healProcessing = sk.GetRes()
						g.cd = 60
					}
				}
			case FightEnd:
				g.curEnemies = make([]*UIEnemy, 0)
				g.selfSkButs = make([]*SkillButton, 0)
				g.dmgSkButs = make([]*SkillButton, 0)
				g.state = Roam
			}
		}
	}
}

func (g *UIGame) updateDoors() {
	neighbours := g.l.GetSliceNeighbours()
	g.currentDoors = make([]*UIDoor, 0)
	for i := 0; i < 3; i++ {
		if neighbours[i] {
			door := UIDoor{
				g.consts.doorX + i*g.consts.doorXOff,
				g.consts.doorY,
				g.consts.doorW,
				g.consts.doorH,
				i,
			}
			g.currentDoors = append(g.currentDoors, &door)
		}
	}
	if neighbours[3] { // should always be true
		door := UIDoor{

			g.consts.backdoorX,
			g.consts.backdoorY,
			g.consts.backdoorW,
			g.consts.backdoorH,
			3,
		}
		g.currentDoors = append(g.currentDoors, &door)
	}
}
