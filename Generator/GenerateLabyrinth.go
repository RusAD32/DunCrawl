package Generator

import (
	. "DunCrawl/Enemies"
	. "DunCrawl/Equipment"
	. "DunCrawl/Interfaces"
	"DunCrawl/Pets"
	. "DunCrawl/PlayerSkills"
	"math/rand"
	"time"
)

const FirstDirection = int(Back)
const LockProbability = 0.2

func GenerateLabyrinth(length, width int) *Labyrinth {
	lab := NewLabyrinth(length, width, FirstDirection, getCorners(length, width))
	rand.Seed(time.Now().UnixNano())
	generateCenter(lab, length, width, width/2*length+length/2)
	for i := range getCorners(length, width) {
		generateBossPath(lab, width, length, i)
	}
	lab.MarkInited()
	p := NewPlayer()
	h := NewHatchet()
	p.Equip(h, MainHand)
	heal := NewHeal(p)
	p.AddSelfSkill(heal)
	cntr := NewCounterSk(p)
	p.AddSelfSkill(cntr)
	atk := NewSimpleAttack(p)
	p.AddDmgSkill(atk)
	stn := NewStunningBlow(p)
	p.AddDmgSkill(stn)
	p.SetPet(Pets.NewDefaultPet())
	lab.SetPlayer(p)
	return lab
}

func getCorners(length, width int) []int {
	return []int{0, length - 1, width*length - 1, (width - 1) * length}
}

func generateCenter(lab *Labyrinth, length, width, startingRoomNum int) {
	rooms := fillWithRooms(length, width, Solid, startingRoomNum)
	lab.AddRooms(rooms)
	lab.SetCurrentRoom(startingRoomNum)
	lab.AddSection(&rooms)
	BackTrackerLabGen(lab.GetCurrentRoom(), 0)
	openDoors(*lab.GetSection(0))
	for i, v := range lab.GetCurrentRoom().GetNeighbours() {
		UnockRooms(lab.GetCurrentRoom(), v.GetNextDoor(), i)
	}
	enemies := make([]*Enemy, 4)
	for i := range enemies {
		enemies[i] = NewDefaultDog(i)
		//enemies[i].name += fmt.Sprintf(" %d", i)
		//enemies[i].skills = append(enemies[i].skills, &bite)
	}
	dfs(lab.GetCurrentRoom(), 0)
}

func generateBossPath(lab *Labyrinth, width, length, dir int) {
	newRooms := fillWithRooms(length, width, Solid, -1)
	firstRoom := newRooms[getCorners(length, width)[rotateRoomNum(dir, 2)]]
	lab.AddSection(&newRooms)
	connectTo := lab.GetRooms()[getCorners(length, width)[dir]]
	BackTrackerLabGen(firstRoom, 0)
	ConnectSection(&newRooms, lab.GetSection(0), firstRoom, connectTo, Direction(rotateRoomNum(dir, 2)))
	//TODO mark as inner or outer based on distance
}

func fillWithRooms(length, width int, kind WallType, startingRoomNum int) []*Room {
	rooms := make([]*Room, 0)
	for i := 0; i < width; i++ {
		for j := 0; j < length; j++ {
			num := i*length + j
			if num == startingRoomNum {
				rooms = append(rooms, GenerateStartingRoom(num))
			} else {
				rooms = append(rooms, GenerateRoom(num))
			}
			if i > 0 {
				ConnectRooms(rooms[i*length+j], rooms[(i-1)*length+j], Forward, kind)
			}
			if j > 0 {
				ConnectRooms(rooms[i*length+j], rooms[i*length+j-1], Left, kind)
			}
		}
	}
	return rooms
}

func GenerateStartingRoom(num int) *Room {
	r := NewRoom([]*Enemy{}, []*Enemy{}, []*Lootable{}, []*Lootable{}, []Stack{}, []Stack{}, nil)
	r.DistFromCenter = -1
	r.Num = num
	return r
}

func GenerateRoom(num int) *Room {
	enemies := make([]*Enemy, 0)
	shEnemies := make([]*Enemy, 0)
	shLoot := make([]*Lootable, 0)
	if rand.Float32() < 0.3 {
		enemies = make([]*Enemy, 4)
		for i := range enemies {
			enemies[i] = NewDefaultDog(i)
		}
	}
	if rand.Float32() < 0.2 {
		shEnemies = make([]*Enemy, 4)
		for i := range shEnemies {
			shEnemies[i] = NewDefaultDog(i)
		}
	}
	if rand.Float32() < 0.2 {
		shLoot = make([]*Lootable, 2)
		for i := range shLoot {
			shLoot[i] = NewLootable("Shadow thingy", 100)
		}
	}
	var chest *Chest
	if rand.Float32() < 0.1 {
		chest = NewChest()
	}
	r := NewRoom(enemies, shEnemies, []*Lootable{}, shLoot, []Stack{}, []Stack{}, chest)
	r.DistFromCenter = -1
	r.Num = num
	return r
}

func dfs(room *Room, dist int) {
	room.DistFromCenter = dist
	for _, v := range room.GetNeighbours() {
		if v.CanGoThrough() {
			next := v.GetNextDoor()
			if next.DistFromCenter < 0 || next.DistFromCenter > dist+1 {
				dfs(next, dist+1)
			}
		}
	}
}

func openDoors(rooms []*Room) {
	for _, v := range rooms {
		for i, d := range v.GetNeighbours() {
			if !d.CanGoThrough() && d.GetNextDoor() != nil && rand.Float32() < LockProbability {
				UnockRooms(v, d.GetNextDoor(), i)
			}
		}
	}
}

func rotateRoomNum(dir, rotAmount int) int {
	return (dir + DOORS_PER_ROOM/rotAmount) % DOORS_PER_ROOM
}
