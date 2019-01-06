package main

import (
	. "DunCrawl/Enemies"
	"DunCrawl/EnemySkills"
	. "DunCrawl/Equipment"
	. "DunCrawl/Generator"
	. "DunCrawl/Interfaces"
	"DunCrawl/PlayerSkills"
	"DunCrawl/UI"
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"math"
)

func walkTest() {
	p := GetDefaultPlayer()
	h := Hatchet{}
	h.Init()
	p.Equip(Equippable(h), MainHand)
	heal := PlayerSkills.Heal{}
	heal.Init(p)
	p.AddSelfSkill(&heal)
	cntr := PlayerSkills.Counter{}
	cntr.Init(p)
	p.AddSelfSkill(&cntr)
	atk := PlayerSkills.SimpleAttack{}
	atk.Init(p)
	p.AddDmgSkill(&atk)
	stn := PlayerSkills.StunningBlow{}
	stn.Init(p)
	p.AddDmgSkill(&stn)

	//dog := GetDefaultEnemy(index)
	enemies := make([]*Enemy, 4)
	for i := range enemies {
		enemies[i] = new(DefaultDog).Init(i)
		//enemies[i].name += fmt.Sprintf(" %d", i)
		bite := EnemySkills.DogBite{}
		bite.Init(enemies[i])
		enemies[i].AddSkill(&bite)
		//enemies[i].skills = append(enemies[i].skills, &bite)
	}

	//f := Room{&p, *[]*Enemy{&dog, &dog2, &dog3, &dog4}, PriorityQueue, 0{}}
	r := Room{}
	//events := make(chan Event)
	r.Init([]*Enemy{}, l)
	r2 := Room{}
	r2.Init(make([]*Enemy, 0), l)
	r2.AddLoot(GenerateLootable("Stuff", 10))
	ch := GetDefaultChest()
	r2.SetChest(ch)
	r.AddShadowLoot(GenerateLootable("Other stuff", 200))
	l := GenerateLabyrinth(3, 3)
	/*rooms := make([]*Room, 2)
	rooms[0] = &r
	rooms[1] = &r2
	ConnectRooms(&r, &r2, Left)
	l.Init(p, rooms, confirm, bgToUi, uiToBg, events)*/
	//	UI.TextFight(&r)
	UI.EnterLabyrinth(l)
}

func labGenTest() {
	l := GenerateLabyrinth(9, 15)
	PrintLabyrinth(l)
}

var l *Labyrinth
var g UI.UIGame

func update(screen *ebiten.Image) error {
	//UI.MoveThroughLabyrinth(l)
	_ = screen.Fill(color.White)
	g.Update()
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	g.Draw(screen)
	//ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%f", scale), 300, 200)
	//w, h := screen.Size()
	//UI.DrawLabyrinth(screen, &l,5, 5, w/5, h/5, color.Black)
	return nil
}

func ebitenTest() {
	l = GenerateLabyrinth(10, 10)
	g.Init(l, 600, 480)
	PrintLabyrinth(l)
	//go UI.EnterLabyrinth(&l)
	x, y := ebiten.ScreenSizeInFullscreen()
	//scale := ebiten.DeviceScaleFactor()
	scale := math.Max(ebiten.DeviceScaleFactor(), 1200/float64(x))
	scale = math.Max(scale, 960/float64(y))
	if err := ebiten.Run(update, 600, 480, 2.0/scale, "DunCrawl"); err != nil {
		panic(err.Error())
	}
}

func main() {
	ebitenTest()
}
