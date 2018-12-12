package Interfaces

import (
	"math/rand"
	"time"
)

func GenerateLabyrinth(length, width int) Labyrinth {
	//TODO remove weird constants, they are calculated on paper
	lab := Labyrinth{
		nil,
		make([]*Room, 0),
		make([]*[]*Room, 0),
		width/2*length + length/2,
		nil,
		3,
		make(chan bool),
		make(chan []SkillInfo),
		make(chan string),
		make(chan Event),
		length,
		width,
		getCorners(length, width),
	}
	rand.Seed(time.Now().UnixNano())
	generateCenter(&lab, length, width)
	for i := range lab.bossEntryRoomNums {
		generateBossPath(&lab, width, length, i)
	}
	return lab
}

func getCorners(length, width int) []int {
	return []int{0, length - 1, width*length - 1, (width - 1) * length}
}

func generateCenter(lab *Labyrinth, length, width int) {
	rooms := fillWithRooms(lab, length, width, Door)
	current := rooms[lab.startingRoomNum]
	lab.rooms = rooms
	lab.current = current
	lab.sections = append(lab.sections, &rooms)
	dfs(lab.current, 0)
	for i, v := range lab.bossEntryRoomNums {
		markPath(lab.rooms[v], i+1)
	}
	lab.current.pathNum = -1
	dfsCloseDoors(lab.current)
	dfs(lab.current, 0)
}

func generateBossPath(lab *Labyrinth, width, length, dir int) {
	newRooms := fillWithRooms(lab, length, width, Solid)
	firstRoom := newRooms[getCorners(width, length)[(dir+2)%len(lab.bossEntryRoomNums)]]
	lab.sections = append(lab.sections, &newRooms)
	connectTo := lab.rooms[lab.bossEntryRoomNums[dir]]
	ConnectSection(&newRooms, lab.sections[0], firstRoom, connectTo, Direction(dir))
	backTrackerLabGen(firstRoom, 0)
	// lab.rooms = append(lab.rooms, newRooms...)
	//TODO mark as inner or outer based on distance
}

func backTrackerLabGen(room *Room, distFromStart int) {
	if room.seenInDfs {
		return
	}
	room.seenInDfs = true
	room.DistFromCenter = distFromStart
	availNeighbours := make([]*Wall, 0)
	availNeighbourNums := make([]int, 0)
	for i, v := range room.neighbours {
		if !v.GetNextDoor().seenInDfs {
			availNeighbours = append(availNeighbours, v)
			availNeighbourNums = append(availNeighbourNums, i)
		}
	}
	for _, v := range rand.Perm(len(availNeighbours)) {
		if !availNeighbours[v].leadsTo.seenInDfs {
			UnockRooms(room, availNeighbours[v].leadsTo, availNeighbourNums[v])
			backTrackerLabGen(room, distFromStart+1)
		}
	}
}

func fillWithRooms(lab *Labyrinth, length, width int, kind WallType) []*Room {
	rooms := make([]*Room, 0)
	for i := 0; i < width; i++ {
		for j := 0; j < length; j++ {
			rooms = append(lab.rooms, GenerateRoom(lab.fightBgToUi, lab.fightUiToBg, lab.fightConfirmChan, i*length+j))
			if i > 0 {
				ConnectRooms(lab.rooms[i*length+j], lab.rooms[(i-1)*length+j], Left, kind)
			}
			if j > 0 {
				ConnectRooms(lab.rooms[i*length+j], lab.rooms[i*length+j-1], Down, kind)
			}
		}
	}
	return rooms
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
			if !v.GetNextDoor().seenInDfs && room.pathNum != -1 && locked < 3 && (room.pathNum == v.GetNextDoor().pathNum && rand.Float32() < 0.1 || rand.Float32() < 0.6) {
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
