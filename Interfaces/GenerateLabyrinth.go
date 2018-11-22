package Interfaces

import (
	"math/rand"
	"time"
)

func GenerateLabyrinth(length, width int) Labyrinth {
	//TODO remove weird constants, they are calculated on paper
	lab := Labyrinth{
		nil,
		make([]*Room, length*width),
		width/2*length + length/2,
		nil,
		3,
		make(chan bool),
		make(chan []SkillInfo),
		make(chan string),
		make(chan Event),
		length,
		width,
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < width; i++ {
		for j := 0; j < length; j++ {
			//fmt.Println(i*length+j)
			lab.rooms[i*length+j] = GenerateRoom(lab.fightBgToUi, lab.fightUiToBg, lab.fightConfirmChan, i*length+j)
			if i > 0 {
				ConnectRooms(lab.rooms[i*length+j], lab.rooms[(i-1)*length+j], Left)
			}
			if j > 0 {
				ConnectRooms(lab.rooms[i*length+j], lab.rooms[i*length+j-1], Down)
			}
		}
	}
	lab.current = lab.rooms[lab.startingRoomNum]
	dfs(lab.current, 0)
	for i, v := range []int{0, length - 1, (width - 1) * length, width*length - 1} {
		markPath(lab.rooms[v], i+1)
	}
	//dfsCloseDoors(lab.current)
	/*for _, v := range lab.rooms {
		for i, w := range v.neighbours {
			if w.CanGoThrough() && v.pathNum != w.leadsTo.pathNum && v != lab.current && w.leadsTo != lab.current && rand.Float32() < 0.35 {
				LockRooms(v, w.leadsTo, i)
			}
		}
		v.DistFromCenter = -1
	}*/

	//техническая информация, убрать
	dfs(lab.current, 0)
	return lab
}

func GenerateRoom(bgToUi chan []SkillInfo, uiToBg chan string, confirm chan bool, num int) *Room {
	r := new(Room)
	//TODO доделать генерацию врагов и лута в комнату
	r.DistFromCenter = -1
	r.Init([]*Enemy{}, bgToUi, uiToBg, confirm)
	r.Num = num
	return r
}

func dfs(room *Room, dist int) {
	room.DistFromCenter = dist
	for _, v := range room.GetNeighbours() {
		if v.CanGoThrough() {
			next := v.leadsTo
			if next.DistFromCenter < 0 || next.DistFromCenter > dist+1 {
				dfs(next, dist+1)
			}
		}
	}
}

func dfsCloseDoors(room *Room) {
	room.seenInDfs = true
	room.DistFromCenter = -1
	locked := 0
	for i, v := range room.GetNeighbours() {
		if v.CanGoThrough() {
			if locked < 3 && (room.pathNum == v.GetNextDoor().pathNum && rand.Float32() < 0.1 || rand.Float32() < 0.6) {
				locked++
				LockRooms(room, v.GetNextDoor(), i)
			} else if !v.GetNextDoor().seenInDfs {
				dfsCloseDoors(v.GetNextDoor())
			}
		} else {
			locked++
		}
	}
}

func markPath(room *Room, pathNum int) {
	if room.DistFromCenter == 0 {
		return
	}
	room.pathNum = pathNum
	for _, v := range room.neighbours {
		if v.CanGoThrough() && v.leadsTo.pathNum == 0 && v.leadsTo.DistFromCenter < room.DistFromCenter {
			markPath(v.leadsTo, pathNum)
			return
		}
	}
}

func PrintLab(l Labyrinth) int {
	locks := 0
	availRooms := 0
	for i := 0; i < l.width; i++ {
		for j := 0; j < l.length; j++ {
			for _, v := range l.rooms[i*l.length+j].neighbours {
				if !v.CanGoThrough() {
					locks++
				}
			}
			if l.rooms[i*l.length+j].DistFromCenter < 0 {
				availRooms++
			}
		}
	}
	return availRooms
}

/*
4 0 0 0 0
3 0 0 0 0
2 1 0 0 0
3 2 1 0 0
4 3 2 3 4
*/
