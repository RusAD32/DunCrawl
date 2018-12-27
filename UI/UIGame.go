package UI

import (
	. "DunCrawl/Interfaces"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"golang.org/x/image/font"
	"image/color"
	"io/ioutil"
)

type GameState int

const (
	Roam GameState = iota
	Fight
)

type UIGame struct {
	w            int
	h            int
	l            *Labyrinth
	state        GameState
	font         font.Face
	currentDoors []*UIDoor
	curEnemies   []*UIEnemy
	nextEnemy    int
	skButs       []*SkillButton
	enemyNums    map[*Enemy]int
	pl           *PlayerStats
	cd           int
}

func (g *UIGame) Init(l *Labyrinth, w, h int) {
	g.l = l
	g.w = w
	g.h = h
	g.updateDoors()
	g.state = Roam
	g.curEnemies = make([]*UIEnemy, 0)
	fontData, err := ioutil.ReadFile("./resources/Roboto-Regular.ttf")
	if err != nil {
		panic("cant load font " + err.Error())
	}
	f, err := truetype.Parse(fontData)
	options := truetype.Options{}
	g.font = truetype.NewFace(f, &options)
	g.skButs = make([]*SkillButton, 0)
	g.enemyNums = make(map[*Enemy]int)
	plst := PlayerStats{
		g.l.GetPlayer(),
		g.w / 10,
		g.h / 10 * 8,
		g.w / 10 * 8,
		g.h / 16,
		g.w / 3,
		g.h / 10 * 9,
		g.w / 10 * 8,
		g.h / 20 * 18,
		color.RGBA{0, 255, 0, 255},
		color.RGBA{0, 0, 255, 255},
		color.RGBA{255, 0, 0, 255},
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

func (g *UIGame) Draw(screen *ebiten.Image) {
	w, h := screen.Size()
	switch g.state {
	case Roam:
		{
			DrawLabyrinth(screen, g.l, 5, 5, w/5, h/5, color.Black)
			for _, v := range g.currentDoors {
				v.Draw(screen, color.Black)
			}
		}
	case Fight:
		{
			for _, v := range g.curEnemies {
				v.Draw(screen, g.font)
			}
			for _, v := range g.skButs {
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
							g.w/4*i + g.w/16,
							g.h / 4,
							g.w / 8,
							g.h / 4,
							color.RGBA{200, 0, 100, 255},
							v,
							false,
							nil,
						}
						g.enemyNums[v] = i
						g.curEnemies = append(g.curEnemies, &enemy)
					}
					g.skButs = make([]*SkillButton, 0)
					for i, v := range g.l.GetPlayer().GetSelfSkillList() {
						button := SkillButton{
							float64(g.w)/16 + float64(g.w)/8*float64(i),
							float64(g.h) * 0.66,
							float64(g.w) / 10,
							float64(g.h) / 10,
							true,
							true,
							v,
						}
						g.skButs = append(g.skButs, &button)
					}
					for i, v := range g.l.GetPlayer().GetDmgSkillList() {
						button := SkillButton{
							float64(g.w)/2 + float64(g.w)/16 + float64(g.w)/8*float64(i),
							float64(g.h) * 0.66,
							float64(g.w) / 10,
							float64(g.h) / 10,
							false,
							false,
							v,
						}
						g.skButs = append(g.skButs, &button)
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
				for _, v := range g.skButs {
					if v.isSelf {
						v.active = true
					}
				}
				for _, v := range g.curEnemies {
					v.col = color.RGBA{200, 0, 100, 255}
				}
				g.pl.dmgJustTaken = ""
				g.pl.healJustTaken = ""
				if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
					for _, v := range g.skButs {
						if v.isClicked(ebiten.CursorPosition()) && v.isSelf {
							g.nextEnemy = 0
							g.l.GetCurrentRoom().SubmitSelfSkill(v.sk.(PlayerSelfSkill))
							g.curEnemies[0].isTargeted = true
							for _, v := range g.skButs {
								if !v.isSelf {
									v.active = v.sk.(PlayerDmgSkill).GetUses() > 0
								} else {
									v.active = false
								}
							}
							return
						}
					}
				}
			case AwaitingDmgSkill:
				if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
					for _, v := range g.skButs {
						if v.isClicked(ebiten.CursorPosition()) && !v.isSelf {
							var curEn *UIEnemy
							for _, v := range g.curEnemies {
								if v.isTargeted {
									curEn = v
								}
							}
							skill := v.sk.(PlayerDmgSkill)
							skill.SetTarget(curEn.enemy)
							curEn.skillUsed = skill
							for _, v := range g.curEnemies {
								if !v.isTargeted && v.skillUsed == nil {
									v.isTargeted = true
									break
								}
							}
							curEn.isTargeted = false
							g.nextEnemy++
							g.l.GetCurrentRoom().SubmitDmgSkill(skill)
							if skill.GetUses() == 0 {
								v.active = false
							}
							return
						}
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
				for _, v := range g.skButs {
					v.active = false
				}
				for _, v := range g.curEnemies {
					v.col = color.RGBA{200, 0, 100, 255}
				}
				sk := g.l.GetCurrentRoom().GetNextSkillUsed()
				target := sk.GetTarget()
				switch sk.(type) {
				case PlayerDmgSkill:
					{
						en := g.curEnemies[g.enemyNums[target.(*Enemy)]]
						en.col = color.RGBA{255, 0, 0, 255}
						en.skillUsed = nil
						g.cd = 60
						g.pl.dmgJustTaken = sk.GetRes()
						g.pl.healJustTaken = ""
					}
				case EnemySkill:
					{
						en := g.curEnemies[g.enemyNums[sk.GetWielder().(*Enemy)]]
						en.col = color.RGBA{255, 0, 100, 255}
						g.cd = 60
						g.pl.dmgJustTaken = sk.GetRes()
						g.pl.healJustTaken = ""
					}
				case PlayerSelfSkill:
					{
						g.pl.dmgJustTaken = ""
						g.pl.healJustTaken = sk.GetRes()
						g.cd = 60
					}
				}
			case FightEnd:
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
				float64(g.w)*0.2 + float64(i*g.w/4),
				float64(g.h) * 0.2,
				float64(g.w) * 0.15,
				float64(g.h) * 0.5,
				i,
			}
			g.currentDoors = append(g.currentDoors, &door)
		}
	}
	if neighbours[3] {
		door := UIDoor{
			float64(g.w) * 0.2,
			float64(g.h) * 0.8,
			float64(g.w) * 0.66,
			float64(g.h) * 0.1,
			3,
		}
		g.currentDoors = append(g.currentDoors, &door)
	}
}
