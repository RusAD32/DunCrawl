package UI

import (
	. "DunCrawl/Interfaces"
	"fmt"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image/color"
	"io/ioutil"
	"strconv"
)

type GameState int

const (
	Roam GameState = iota
	Fight
)

type SkillButton struct {
	x, y, w, h float64
	isSelf     bool
	sk         Skill
}

func (sb *SkillButton) Draw(screen *ebiten.Image, font font.Face) {
	ebitenutil.DrawRect(screen, sb.x, sb.y, sb.w, sb.h, color.RGBA{200, 200, 200, 255})
	text.Draw(screen, sb.sk.GetName(), font, int(sb.x), int(sb.y)+font.Metrics().Height.Ceil()*2, color.Black)
}

func (sb *SkillButton) isClicked(mouseX, mouseY int) bool {
	return !(mouseX < int(sb.x) || mouseX > int(sb.x+sb.w) || mouseY < int(sb.y) || mouseY > int(sb.y+sb.h))
}

type UIGame struct {
	w            int
	h            int
	l            *Labyrinth
	state        GameState
	font         font.Face
	currentDoors []*UIDoor
	curEnemies   []*UIEnemy
	nextEnemy    int
	skButs       []SkillButton
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
			ebitenutil.DebugPrint(screen, strconv.Itoa(int(g.state)))
			DrawLabyrinth(screen, g.l, 5, 5, w/5, h/5, color.Black)
			for _, v := range g.currentDoors {
				v.Draw(screen, color.Black)
			}
		}
	case Fight:
		{
			ebitenutil.DebugPrint(screen, strconv.Itoa(g.l.GetPlayer().GetCurHP())+"\n"+strconv.Itoa(int(g.l.GetCurrentRoom().FightState)))
			for _, v := range g.curEnemies {
				v.Draw(screen, g.font)
			}
			for _, v := range g.skButs {
				v.Draw(screen, g.font)
			}
		}
	}

}

func (g *UIGame) Update() {
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
						}
						g.curEnemies = append(g.curEnemies, &enemy)
						g.skButs = make([]SkillButton, 0)
						for i, v := range g.l.GetPlayer().GetSelfSkillList() {
							button := SkillButton{
								float64(g.w)/16 + float64(g.w)/8*float64(i),
								float64(g.h) * 0.8,
								float64(g.w) / 10,
								float64(g.h) / 10,
								true,
								v,
							}
							g.skButs = append(g.skButs, button)
						}
						for i, v := range g.l.GetPlayer().GetDmgSkillList() {
							button := SkillButton{
								float64(g.w)/2 + float64(g.w)/16 + float64(g.w)/8*float64(i),
								float64(g.h) * 0.8,
								float64(g.w) / 10,
								float64(g.h) / 10,
								false,
								v,
							}
							g.skButs = append(g.skButs, button)
						}
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
				if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
					for _, v := range g.skButs {
						if v.isClicked(ebiten.CursorPosition()) && v.isSelf {
							g.nextEnemy = 0
							g.l.GetCurrentRoom().SubmitSelfSkill(v.sk.(PlayerSelfSkill))

						}
					}
				}
			case AwaitingDmgSkill:
				if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
					for _, v := range g.skButs {
						if v.isClicked(ebiten.CursorPosition()) && !v.isSelf {
							skill := v.sk.(PlayerDmgSkill)
							skill.SetTarget(g.curEnemies[g.nextEnemy].enemy)
							g.nextEnemy++
							g.l.GetCurrentRoom().SubmitDmgSkill(skill)

						}
					}
				}
			case ResolvingSkills:
				fmt.Println(g.l.GetCurrentRoom().GetNextSkillUsed().GetRes())
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
