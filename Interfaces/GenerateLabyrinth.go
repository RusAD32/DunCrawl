package Interfaces

import (
	"math/rand"
	"time"
)

const FirstDirection = int(Back)
const FirstRoomPath = -1
const MaxLocks = 3
const LockProbability = 0.2

func GenerateLabyrinth(length, width int) Labyrinth {
	//TODO remove weird constants, they are calculated on paper
	lab := Labyrinth{
		nil,
		make([]*Room, 0),
		make([]*[]*Room, 0),
		width/2*length + length/2, // center
		nil,
		FirstDirection,
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
	p := GetDefaultPlayer()
	lab.p = p
	lab.current.p = p
	return lab
}

func getCorners(length, width int) []int {
	return []int{0, length - 1, width*length - 1, (width - 1) * length}
}

func generateCenter(lab *Labyrinth, length, width int) {
	rooms := fillWithRooms(lab, length, width, Solid)
	current := rooms[lab.startingRoomNum]
	lab.rooms = rooms
	lab.current = current
	lab.sections = append(lab.sections, &rooms)
	backTrackerLabGen(lab.current, 0)
	for i, v := range lab.bossEntryRoomNums {
		markPath(lab.rooms[v], i+1)
	}
	lab.current.pathNum = FirstRoomPath
	openDoors(*lab.sections[0])
	for i, v := range lab.current.neighbours {
		UnockRooms(lab.current, v.leadsTo, i)
	}
	dfs(lab.current, 0)
}

func generateBossPath(lab *Labyrinth, width, length, dir int) {
	newRooms := fillWithRooms(lab, length, width, Solid)
	firstRoom := newRooms[getCorners(width, length)[rotateRoomNum(dir, 2)]]
	lab.sections = append(lab.sections, &newRooms)
	connectTo := lab.rooms[lab.bossEntryRoomNums[dir]]
	backTrackerLabGen(firstRoom, 0)
	ConnectSection(&newRooms, lab.sections[0], firstRoom, connectTo, Direction(rotateRoomNum(dir, 2)))
	//TODO mark as inner or outer based on distance
}

func backTrackerLabGen(room *Room, distFromStart int) {
	if room.seenInDfs {
		return
	}
	room.seenInDfs = true
	room.DistFromCenter = distFromStart
	availNeighbourNums := make([]int, 0)
	for i, v := range room.neighbours {
		if v.GetNextDoor() != nil {
			availNeighbourNums = append(availNeighbourNums, i)
		}
	}
	for _, v := range rand.Perm(len(availNeighbourNums)) {
		num := availNeighbourNums[v]
		if !room.neighbours[num].leadsTo.seenInDfs {
			UnockRooms(room, room.neighbours[num].leadsTo, num)
			backTrackerLabGen(room.neighbours[num].leadsTo, distFromStart+1)
		}
	}
}

func fillWithRooms(lab *Labyrinth, length, width int, kind WallType) []*Room {
	rooms := make([]*Room, 0)
	for i := 0; i < width; i++ {
		for j := 0; j < length; j++ {
			rooms = append(rooms, GenerateRoom(lab.fightBgToUi, lab.fightUiToBg, lab.fightConfirmChan, i*length+j))
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

func openDoors(rooms []*Room) {
	for _, v := range rooms {
		for i, d := range v.neighbours {
			if !d.CanGoThrough() && d.leadsTo != nil && rand.Float32() < LockProbability {
				UnockRooms(v, d.leadsTo, i)
			}
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

func rotateRoomNum(dir, rotAmount int) int {
	return (dir + DOORS_PER_ROOM/rotAmount) % DOORS_PER_ROOM
}

func ifLock(room, next *Room) bool {
	locked := room.GetLocks()
	nextLocked := next.GetLocks()
	return (!next.seenInDfs) &&
		room.pathNum != -1 &&
		locked < MaxLocks &&
		nextLocked < MaxLocks-1 &&
		(room.pathNum == 0 || room.pathNum != next.pathNum) &&
		rand.Float32() < LockProbability
}
