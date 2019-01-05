package UI

import (
	. "DunCrawl/Interfaces"
	"fmt"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"golang.org/x/image/font"
	"image/color"
)

type GameState int

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
	LightRed   = color.RGBA{R: 255, G: 200, B: 200, A: 255}
)

type UIGame struct {
	w                                   int
	h                                   int
	l                                   *Labyrinth
	state                               GameState
	font                                font.Face
	currentDoors                        []*UIDoor
	curEnemies                          []*UIEnemy
	selfSkButs                          []*SkillButton
	dmgSkButs                           []*SkillButton
	enemyNums                           map[*Enemy]int
	pl                                  *PlayerStats
	cd                                  int
	queue                               SkQueue
	resolvingSk                         *SkillIcon
	turnStartActions, preResolveActions bool
	consts
}

func (g *UIGame) Init(l *Labyrinth, w, h int) {
	g.l = l
	g.w = w
	g.h = h

	g.consts = getConstants(w, h)

	g.updateDoors()
	g.curEnemies = make([]*UIEnemy, 0)
	fontRaw, err := LoadResource("Roboto-Regular.ttf")
	if err != nil {
		panic(err)
	}
	fontData, err := truetype.Parse(fontRaw)
	if err != nil {
		panic(err)
	}
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
	g.queue = SkQueue{
		x:      w / 7,
		y:      h / 12,
		xOffs:  g.consts.skButW * 10 / 8,
		skills: make([]*SkillIcon, 0),
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
	switch g.l.GetState() {
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
				v.Draw(screen)
			}
			g.pl.Draw(screen, g.font)
			g.queue.Draw(screen)
			if g.resolvingSk != nil {
				g.resolvingSk.Draw(screen)
			}
		}
	}

}

func (g *UIGame) prepareForFight() {
	ens := g.l.GetCurrentRoom().GetEnemies()
	for i, v := range ens {
		enemy := new(UIEnemy).Init(g.consts.enemyXOff*i+g.consts.enemyX,
			g.consts.enemyY,
			g.consts.enemyW,
			g.consts.enemyH,
			Violet,
			Firebrick,
			OrangeRed,
			color.Black,
			v)
		g.enemyNums[v] = i
		g.curEnemies = append(g.curEnemies, enemy)
	}
	for i, v := range g.l.GetPlayer().GetSelfSkillList() {
		button := new(SkillButton).Init(
			g.consts.selfSkButX+i*g.consts.skButXOff,
			g.consts.skButY,
			g.consts.skButW,
			g.consts.skButH,
			v,
			LightGreen,
			Gray,
			g.font)
		g.selfSkButs = append(g.selfSkButs, button)
	}
	for i, v := range g.l.GetPlayer().GetDmgSkillList() {
		button := new(SkillButton).Init(
			g.consts.dmgSkButX+i*g.consts.skButXOff,
			g.consts.skButY,
			g.consts.skButW,
			g.consts.skButH,
			v,
			LightBlue,
			Gray,
			g.font)
		g.dmgSkButs = append(g.dmgSkButs, button)
	}
	g.turnStartActions = true
}

func (g *UIGame) ConstructSkillIcon(skill SkillInfo, w, h int) *SkillIcon {
	var col color.Color
	switch skill.(type) {
	case PlayerDmgSkill:
		col = LightBlue
	case PlayerSelfSkill:
		col = LightGreen
	default:
		col = LightRed
	}
	return new(SkillIcon).Init(w, h, skill, col, g.font)
}

func (g *UIGame) updateQueue() {
	g.queue.skills = make([]*SkillIcon, 0)
	skQ := g.l.GetCurrentRoom().GetSkQueue()
	for _, v := range skQ {
		g.queue.skills = append(g.queue.skills, g.ConstructSkillIcon(v, g.consts.skButW, g.consts.skButH))
	}
}

func getNewClicks() [][]int {
	res := make([][]int, 0)
	for _, v := range inpututil.JustPressedTouchIDs() {
		click := make([]int, 2)
		x, y := ebiten.TouchPosition(v)
		click[0] = x
		click[1] = y
		res = append(res, click)
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		click := make([]int, 2)
		x, y := ebiten.CursorPosition()
		click[0] = x
		click[1] = y
		res = append(res, click)

	}
	return res
}

func (g *UIGame) submitSelfSkill() {
	if g.turnStartActions {
		g.resolvingSk = nil
		g.queue.skills = make([]*SkillIcon, 0)
		for _, v := range g.selfSkButs {
			v.state = butActive
		}
		for _, v := range g.curEnemies {
			v.state = enemyDefault
		}
		g.pl.dmgProcessing = ""
		g.pl.healProcessing = ""
		g.preResolveActions = true
		g.turnStartActions = false
	}
	for _, v := range getNewClicks() {
		butNum := g.selfSkillButtonClicked(v[0], v[1])
		if butNum >= 0 {
			g.l.GetCurrentRoom().SubmitSelfSkill(g.selfSkButs[butNum].sk.(PlayerSelfSkill))
			g.curEnemies[0].isTargeted = true
			for _, v := range g.selfSkButs {
				v.state = butInactive
			}
			for _, v := range g.dmgSkButs {
				if v.sk.(PlayerDmgSkill).GetUses() != 0 {
					v.state = butActive
				} else {
					v.state = butInactive
				}
			}
			g.updateQueue()
			return
		}
	}
}

func (g *UIGame) submitDmgSkill() {
	for _, touch := range getNewClicks() {
		skNum := g.dmgSkillButtonClicked(touch[0], touch[1])
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
			g.updateQueue()
			if skill.GetUses() != 0 {
				g.dmgSkButs[skNum].state = butActive
			} else {
				g.dmgSkButs[skNum].state = butInactive
			}
			return
		}
		for _, v := range g.curEnemies {
			if v.isClicked(touch[0], touch[1]) && v.skillUsed == nil {
				for _, v := range g.curEnemies {
					v.isTargeted = false
				}
				v.isTargeted = true
				return
			}
		}
	}
}

func (g *UIGame) resolveSkill() {
	if g.preResolveActions {
		for _, v := range append(g.selfSkButs, g.dmgSkButs...) {
			v.state = butInactive
		}
		g.preResolveActions = false
		g.turnStartActions = true
	}
	for _, v := range g.curEnemies {
		v.state = enemyDefault
	}
	sk := g.l.GetCurrentRoom().GetNextSkillUsed()
	g.resolvingSk = g.ConstructSkillIcon(sk, g.consts.skButW*4/3, g.consts.skButH*4/3)
	g.resolvingSk.x = 0
	g.resolvingSk.y = g.h / 14
	g.resolvingSk.opts.GeoM.Translate(float64(g.resolvingSk.x), float64(g.resolvingSk.y))
	g.updateQueue()
	target := sk.GetTarget()
	switch sk.(type) {
	case PlayerDmgSkill:
		{
			en := g.curEnemies[g.enemyNums[target.(*Enemy)]]
			en.state = enemyAttacked
			en.skillUsed = nil
			g.cd = 60
			g.pl.dmgProcessing = sk.GetRes()
			g.pl.healProcessing = ""
		}
	case NPCSkill:
		{
			en := g.curEnemies[g.enemyNums[sk.GetWielder().(*Enemy)]]
			en.state = enemyAttacking
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
}

func (g *UIGame) Update() {
	if g.cd > 0 {
		g.cd--
		return
	}
	switch g.l.GetState() {
	case Roam:
		{
			for _, v := range getNewClicks() {
				nextDoor := g.doorClicked(v[0], v[1])
				/*go g.l.GoToRoom(Direction(nextDoor)) // TODO this part should be stateful imho
				event := <-g.l.GetEventsChan()
				if event == FightEvent {
					g.prepareForFight()
					g.l.GetCurrentRoom().AtTurnStart()
				}*/
				if g.l.GoToRoomNew(Direction(nextDoor)) {
					g.prepareForFight()
					g.l.GetCurrentRoom().AtTurnStart()
				}
				g.updateDoors()
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
				fmt.Println(g.l.GetValues()) // TODO display this on screen
				g.curEnemies = make([]*UIEnemy, 0)
				g.selfSkButs = make([]*SkillButton, 0)
				g.dmgSkButs = make([]*SkillButton, 0)
			}
		}
	}
}

func (g *UIGame) updateDoors() {
	neighbours := g.l.GetSliceNeighbours()
	g.currentDoors = make([]*UIDoor, 0)
	for i := 0; i < 3; i++ {
		if neighbours[i] {
			door := new(UIDoor).Init(
				g.consts.doorX+i*g.consts.doorXOff,
				g.consts.doorY,
				g.consts.doorW,
				g.consts.doorH,
				i)
			g.currentDoors = append(g.currentDoors, door)
		}
	}
	if neighbours[3] { // should always be true
		door := new(UIDoor).Init(
			g.consts.backdoorX,
			g.consts.backdoorY,
			g.consts.backdoorW,
			g.consts.backdoorH,
			3)
		g.currentDoors = append(g.currentDoors, door)
	}
}
