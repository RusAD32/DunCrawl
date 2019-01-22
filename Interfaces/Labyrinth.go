package Interfaces

type LabyrinthState int

const (
	Initializing LabyrinthState = iota
	Roam
	Fight
	Exited
)

const DOORS_PER_ROOM = 4
const NEW_NEIGHBOUR_OFFSET = -1

//TODO написать NewTrigger для лабиринта
//Что именно я имел в виду? Триггеры перехода в комнату?

type Labyrinth struct {
	p                 *Player
	rooms             []*Room
	sections          []*[]*Room
	startingRoomNum   int
	current           *Room
	previous          int
	length            int
	width             int
	bossEntryRoomNums []int
	state             LabyrinthState
}

func NewLabyrinth(width, length, FirstDirection int, corners []int) *Labyrinth {
	return &Labyrinth{
		rooms:             make([]*Room, 0),
		sections:          make([]*[]*Room, 0),
		startingRoomNum:   width/2*length + length/2, // center
		previous:          FirstDirection,
		length:            length,
		width:             width,
		bossEntryRoomNums: corners,
		state:             Initializing,
	}
}

func (l *Labyrinth) MarkInited() {
	l.state = Roam
}

func (l *Labyrinth) switchRooms(direction Direction) bool {
	if l.current == nil {
		l.current = l.rooms[l.startingRoomNum]
	} else if int(direction) >= 0 {
		neighbourWall := l.current.GetNeighbours()[getNextDoorNum(int(direction), l.previous)]
		if !neighbourWall.CanGoThrough() {
			return false
		}
		l.current.p = nil
		if neighbourWall.kind == NextSection {
			l.rooms = *neighbourWall.nextSection
		}
		l.current = neighbourWall.leadsTo
		l.previous = (l.previous + int(direction) + DOORS_PER_ROOM + NEW_NEIGHBOUR_OFFSET) % len(l.current.neighbours)
	}
	return true
}

func (l *Labyrinth) GotoRoom(direction Direction) bool {
	// Returning true if the fight starts in that room
	if !l.switchRooms(direction) {
		return false
	}
	l.current.p = l.p
	if l.current.FightState == TurnStart {
		l.state = Fight
		l.current.AtTurnStart()
		return true
	}
	return false
}

func (l *Labyrinth) GetValues() ([]*Lootable, []Stack) {
	if l.state != Fight || l.current.FightState == FightEnd {
		l.state = Roam
		return l.current.GetValues()
	}
	return nil, nil
}

func (l *Labyrinth) GetState() LabyrinthState {
	return l.state
}

func (l *Labyrinth) Light() {
	if l.current.HasShadowEnemies() {
		l.state = Fight
	}
	l.current.Light()
}

func (l *Labyrinth) UnlockChest() {
	l.current.UnlockChest()
}

func (l *Labyrinth) GetNeighbours() map[string]bool {
	res := make(map[string]bool)
	for i, v := range l.current.GetNeighbours() {
		res[DirToStr[getRelativeDirection(i, l.previous)]] = v.CanGoThrough()
	}
	return res
}

func (l *Labyrinth) GetSliceNeighbours() []bool {
	res := make([]bool, 4)
	for i, v := range l.current.GetNeighbours() {
		res[int(getRelativeDirection(i, l.previous))] = v.CanGoThrough()
	}
	return res
}

func (l *Labyrinth) GetCurrentRoom() *Room {
	return l.current
}

func (l *Labyrinth) GetPlayer() *Player {
	return l.p
}

func (l *Labyrinth) GetWidth() int {
	return l.width
}

func (l *Labyrinth) GetLength() int {
	return l.length
}

func (l *Labyrinth) GetRooms() []*Room {
	return l.rooms
}

func getNextDoorNum(direction, previous int) int {
	return (direction + previous + 1) % DOORS_PER_ROOM
}

func getRelativeDirection(newDirection, prevDirection int) Direction {
	return Direction((-prevDirection + newDirection + DOORS_PER_ROOM + NEW_NEIGHBOUR_OFFSET) % DOORS_PER_ROOM)
}

func (l *Labyrinth) GetPrevious() int {
	return l.previous
}

func (l *Labyrinth) SetPlayer(p *Player) {
	if l.current != nil {
		l.current.p = p
	}
	l.p = p
}

func (l *Labyrinth) AddRooms(rooms []*Room) {
	l.rooms = append(l.rooms, rooms...)
}

func (l *Labyrinth) SetCurrentRoom(roomnum int) {
	l.current = l.rooms[roomnum]
}

func (l *Labyrinth) AddSection(sec *[]*Room) {
	l.sections = append(l.sections, sec)
}

func (l *Labyrinth) GetSection(num int) *[]*Room {
	return l.sections[num]
}

func (l *Labyrinth) GetResources() []string {
	resources := []string{
		"resources/UIElements/Background.png",
		"resources/UIElements/DoorFwd.png",
		"resources/UIElements/DoorLft.png",
		"resources/UIElements/DoorRgt.png",
		"resources/UIElements/lamp_t.png",
	}
	return append(resources, resourcesDFS(l.GetCurrentRoom(), 0, 5, []int{})...)

}

func resourcesDFS(start *Room, curDepth, maxDepth int, alreadySeen []int) []string {
	if curDepth > maxDepth || inSlice(alreadySeen, start.Num) {
		return []string{}
	}
	alreadySeen = append(alreadySeen, start.Num)
	res := make([]string, 0)
	for _, v := range start.enemies {
		// The texture preloader already filters out the duplicates, so it shouldn't be a big deal
		res = append(res, v.IdleImgsPath()...)
		res = append(res, v.AttackedImgsPath()...)
		res = append(res, v.SkillImgsPath()...)
		res = append(res, v.DeadImgsPath()...)
		for _, sk := range v.skills {
			res = append(res, sk.GetIconPath())
		}
	}
	for _, v := range start.GetNeighbours() {
		if v.CanGoThrough() {
			res = append(res, resourcesDFS(v.GetNextDoor(), curDepth+1, maxDepth, alreadySeen)...)
		}
	}
	return res
}

func inSlice(sl []int, num int) bool {
	for _, v := range sl {
		if v == num {
			return true
		}
	}
	return false
}
