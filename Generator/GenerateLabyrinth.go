package Generator

import (
	. "DunCrawl/Enemies"
	"DunCrawl/EnemySkills"
	. "DunCrawl/Equipment"
	. "DunCrawl/Interfaces"
	"DunCrawl/Pets"
	"DunCrawl/PlayerSkills"
	"math/rand"
	"time"
)

const FirstDirection = int(Back)
const LockProbability = 0.2

func GenerateLabyrinth(length, width int) *Labyrinth {
	lab := GetDefaultLabyrinth(length, width, FirstDirection, getCorners(length, width))
	rand.Seed(time.Now().UnixNano())
	generateCenter(lab, length, width, width/2*length+length/2)
	for i := range getCorners(length, width) {
		generateBossPath(lab, width, length, i)
	}
	lab.MarkInited()
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
	p.SetPet(new(Pets.DefaultPet).Init())
	lab.SetPlayer(p)
	return lab
}

func getCorners(length, width int) []int {
	return []int{0, length - 1, width*length - 1, (width - 1) * length}
}

func generateCenter(lab *Labyrinth, length, width, startingRoomNum int) {
	rooms := fillWithRooms(lab, length, width, Solid)
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
		enemies[i] = new(DefaultDog).Init(i)
		//enemies[i].name += fmt.Sprintf(" %d", i)
		//enemies[i].skills = append(enemies[i].skills, &bite)
	}
	dfs(lab.GetCurrentRoom(), 0)
}

func generateBossPath(lab *Labyrinth, width, length, dir int) {
	newRooms := fillWithRooms(lab, length, width, Solid)
	firstRoom := newRooms[getCorners(length, width)[rotateRoomNum(dir, 2)]]
	lab.AddSection(&newRooms)
	connectTo := lab.GetRooms()[getCorners(length, width)[dir]]
	BackTrackerLabGen(firstRoom, 0)
	ConnectSection(&newRooms, lab.GetSection(0), firstRoom, connectTo, Direction(rotateRoomNum(dir, 2)))
	//TODO mark as inner or outer based on distance
}

func fillWithRooms(lab *Labyrinth, length, width int, kind WallType) []*Room {
	rooms := make([]*Room, 0)
	for i := 0; i < width; i++ {
		for j := 0; j < length; j++ {
			rooms = append(rooms, GenerateRoom(lab, i*length+j))
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

func GenerateRoom(l *Labyrinth, num int) *Room {
	r := new(Room)
	enemies := make([]*Enemy, 0)
	if rand.Float32() < 0.3 {
		enemies = make([]*Enemy, 4)
		for i := range enemies {
			enemies[i] = new(DefaultDog).Init(i)
			//enemies[i].name += fmt.Sprintf(" %d", i)
			bite := EnemySkills.DogBite{}
			bite.Init(enemies[i])
			enemies[i].AddSkill(&bite)
			//enemies[i].skills = append(enemies[i].skills, &bite)
		}
	}
	if rand.Float32() < 0.8 {
		r.SetChest(GetDefaultChest())
	}
	r.DistFromCenter = -1
	r.Init(enemies, l)
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
