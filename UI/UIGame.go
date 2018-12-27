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

type UIGame struct {
	w            int
	h            int
	l            *Labyrinth
	state        GameState
	font         font.Face
	currentDoors []*UIDoor
	curEnemies   []*UIEnemy
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
					}
					g.state = Fight
				}
				g.updateDoors()
			}
		}
	case Fight:
		{

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
