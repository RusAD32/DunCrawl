package main

import (
	"./EnemySkills"
	. "./Equipment"
	. "./Interfaces"
	"./PlayerSkills"
	"./UI"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"strconv"
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
		enemies[i] = GetDefaultEnemy(i)
		//enemies[i].name += fmt.Sprintf(" %d", i)
		bite := EnemySkills.DogBite{}
		bite.Init(enemies[i])
		enemies[i].AddSkill(&bite)
		//enemies[i].skills = append(enemies[i].skills, &bite)
	}

	//f := Room{&p, *[]*Enemy{&dog, &dog2, &dog3, &dog4}, PriorityQueue, 0{}}
	r := Room{}
	uiToBg := make(chan string)
	bgToUi := make(chan []SkillInfo)
	confirm := make(chan bool)
	//events := make(chan Event)
	r.Init([]*Enemy{}, bgToUi, uiToBg, confirm)
	r2 := Room{}
	r2.Init(make([]*Enemy, 0), bgToUi, uiToBg, confirm)
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
	UI.EnterLabyrinth(&l)
}

func labGenTest() {
	l := GenerateLabyrinth(9, 15)
	PrintLabyrinth(&l)
}

var l Labyrinth

func update(screen *ebiten.Image) error {
	screen.Fill(color.White)
	w, h := screen.Size()
	roomW := (w - 1) / l.GetWidth()
	roomH := (h - 1) / l.GetLength()
	rooms := l.GetRooms()
	for i := 0; i < l.GetWidth(); i++ {
		for j := 0; j < l.GetLength(); j++ {
			room := rooms[i*l.GetLength()+j]
			walls := room.GetNeighbours()
			ebitenutil.DebugPrintAt(screen, strconv.Itoa(room.Num), roomW*j, roomH*i)
			if !walls[int(Forward)].CanGoThrough() {
				//fmt.Println(room.Num, "fwd")
				ebitenutil.DrawLine(screen, float64(roomW*j), float64(roomH*i), float64(roomW*(j+1)), float64(roomH*i), color.Black)
			}
			if !walls[int(Left)].CanGoThrough() {
				//fmt.Println(room.Num, "lft")
				ebitenutil.DrawLine(screen, float64(roomW*j), float64(roomH*i), float64(roomW*j), float64(roomH*(i+1)), color.Black)
			}
			if !walls[int(Right)].CanGoThrough() {
				//fmt.Println(room.Num, "right")
				ebitenutil.DrawLine(screen, float64(roomW*(j+1)), float64(roomH*i), float64(roomW*(j+1)), float64(roomH*(i+1)), color.Black)
			}
			if !walls[int(Back)].CanGoThrough() {
				//fmt.Println(room.Num, "dwn")
				ebitenutil.DrawLine(screen, float64(roomW*j), float64(roomH*(i+1)), float64(roomW*(j+1)), float64(roomH*(i+1)), color.Black)
			}
			if rooms[i*l.GetLength()+j] == l.GetCurrentRoom() {
				//ebitenutil.DebugPrintAt(screen, "p", roomW*i, roomH*i)
			}
		}
	}
	//ebitenutil.DebugPrint(screen, "Hello world!")
	return nil
}

func ebitenTest() {
	l = GenerateLabyrinth(10, 10)
	PrintLabyrinth(&l)
	if err := ebiten.Run(update, 800, 600, 1, "Hello world!"); err != nil {
		panic(err.Error())
	}
}

func main() {
	ebitenTest()
}
