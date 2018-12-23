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

func getRoomCoords(i, j, offset, roomW, roomH int) (float64, float64) {
	return float64(roomW*j + offset), float64(roomH*i + offset)
}

func DrawPlayer(screen *ebiten.Image, x, y float64, roomW, roomH, dir int, col color.Color) {
	playerW := float64(roomW) * 0.8
	playerWoffs := float64(roomW) * 0.1
	playerH := float64(roomH) * 0.8
	playerHoffs := float64(roomH) * 0.1
	switch dir {
	case 0:
		{
			ax := x + playerWoffs
			ay := y + playerHoffs
			bx := ax + playerW
			by := ay + playerH/2.0
			cx := ax
			cy := ay + playerH
			ebitenutil.DrawLine(screen, ax, ay, bx, by, col)
			ebitenutil.DrawLine(screen, bx, by, cx, cy, col)
			ebitenutil.DrawLine(screen, cx, cy, ax, ay, col)
		}
	case 3:
		{
			ax := x + playerWoffs
			ay := y + playerHoffs + playerH
			bx := ax + playerW/2
			by := ay - playerH
			cx := ax + playerW
			cy := ay
			ebitenutil.DrawLine(screen, ax, ay, bx, by, col)
			ebitenutil.DrawLine(screen, bx, by, cx, cy, col)
			ebitenutil.DrawLine(screen, cx, cy, ax, ay, col)
		}
	case 2:
		{
			ax := x + playerWoffs + playerW
			ay := y + playerHoffs
			bx := ax - playerW
			by := ay + playerH/2.0
			cx := ax
			cy := ay + playerH
			ebitenutil.DrawLine(screen, ax, ay, bx, by, col)
			ebitenutil.DrawLine(screen, bx, by, cx, cy, col)
			ebitenutil.DrawLine(screen, cx, cy, ax, ay, col)
		}
	case 1:
		{
			ax := x + playerWoffs
			ay := y + playerHoffs
			bx := ax + playerW/2.0
			by := ay + playerH
			cx := ax + playerW
			cy := ay
			ebitenutil.DrawLine(screen, ax, ay, bx, by, col)
			ebitenutil.DrawLine(screen, bx, by, cx, cy, col)
			ebitenutil.DrawLine(screen, cx, cy, ax, ay, col)
		}
	}
}

func DrawLabyrinth(screen *ebiten.Image, col color.Color) {
	offset := 5
	w, h := screen.Size()
	roomW := (w - offset*2) / l.GetWidth()
	roomH := (h - offset*2) / l.GetLength()
	rooms := l.GetRooms()
	for i := 0; i < l.GetWidth(); i++ {
		for j := 0; j < l.GetLength(); j++ {
			room := rooms[i*l.GetLength()+j]
			walls := room.GetNeighbours()
			//ebitenutil.DebugPrintAt(screen, strconv.Itoa(room.Num), roomW*j + offset, roomH*i + offset)
			roomX, roomY := getRoomCoords(i, j, offset, roomW, roomH)
			nextX := roomX + float64(roomW)
			nextY := roomY + float64(roomH)
			if !walls[int(Forward)].CanGoThrough() {
				//fmt.Println(room.Num, "fwd")
				ebitenutil.DrawLine(screen, roomX, roomY, roomX+float64(roomW), roomY, col)
			}
			if !walls[int(Left)].CanGoThrough() {
				//fmt.Println(room.Num, "lft")
				ebitenutil.DrawLine(screen, roomX, roomY, roomX, nextY, col)
			}
			if !walls[int(Right)].CanGoThrough() {
				//fmt.Println(room.Num, "right")
				ebitenutil.DrawLine(screen, nextX, roomY, nextX, nextY, col)
			}
			if !walls[int(Back)].CanGoThrough() {
				//fmt.Println(room.Num, "dwn")
				ebitenutil.DrawLine(screen, roomX, nextY, nextX, nextY, col)
			}
			if rooms[i*l.GetLength()+j] == l.GetCurrentRoom() {
				DrawPlayer(screen, roomX, roomY, roomW, roomH, l.GetPrevious(), col)
			}
		}
	}
}

func update(screen *ebiten.Image) error {
	err := screen.Fill(color.White)
	if err != nil {
		panic("can't fill the screen with color")
	}
	DrawLabyrinth(screen, color.Black)
	//ebitenutil.DebugPrint(screen, "Hello world!")
	return nil
}

func ebitenTest() {
	l = GenerateLabyrinth(10, 10)
	PrintLabyrinth(&l)
	go UI.EnterLabyrinth(&l)
	if err := ebiten.Run(update, 600, 480, 2, "Hello world!"); err != nil {
		panic(err.Error())
	}
}

func main() {
	ebitenTest()
}
