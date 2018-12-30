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
	LightRed   = color.RGBA{R: 255, G: 200, B: 200, A: 255}
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
	queue        SkQueue
	consts
}

func (g *UIGame) Init(l *Labyrinth, w, h int) {
	g.l = l
	g.w = w
	g.h = h

	g.consts = getConstants(w, h)

	g.updateDoors()
	g.state = Roam
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
		x:      w / 10,
		y:      h / 10,
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
				v.Draw(screen)
			}
			g.pl.Draw(screen, g.font)
			g.queue.Draw(screen)
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
	g.state = Fight
}

func (g *UIGame) updateQueue() {
	g.queue.skills = make([]*SkillIcon, 0)
	skQ := g.l.GetCurrentRoom().GetSkQueue()
	for _, v := range skQ {
		var col color.Color
		switch v.(type) {
		case PlayerSelfSkill:
			col = LightGreen
		case PlayerDmgSkill:
			col = LightBlue
		default:
			{
				col = LightRed
			}
		}
		g.queue.skills = append(g.queue.skills, new(SkillIcon).Init(g.consts.skButW, g.consts.skButH, v, col, g.font))
	}
}

func (g *UIGame) submitSelfSkill() {
	g.queue.skills = make([]*SkillIcon, 0)
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
			g.updateQueue()
			return
		}
	}
}

func (g *UIGame) submitDmgSkill() {
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
			g.updateQueue()
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
}

func (g *UIGame) resolveSkill() {
	for _, v := range append(g.selfSkButs, g.dmgSkButs...) {
		v.active = false
	}
	for _, v := range g.curEnemies {
		v.col = Violet
	}
	sk := g.l.GetCurrentRoom().GetNextSkillUsed()
	g.updateQueue()
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
				go g.l.GoToRoom(Direction(nextDoor)) // TODO this part should be stateful imho
				event := <-g.l.GetEventsChan()
				if event == FightEvent {
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
