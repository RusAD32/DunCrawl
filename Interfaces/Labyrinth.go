package Interfaces

type Event int

const (
	NoEvent Event = iota
	FightEvent
)

const DOORS_PER_ROOM = 4
const NEW_NEIGHBOUR_OFFSET = -1

//TODO давать плееру то, что он в комнатах нашел
//TODO написать Init для лабиринта

type Labyrinth struct {
	p                 *Player
	rooms             []*Room
	sections          []*[]*Room
	startingRoomNum   int
	current           *Room
	previous          int
	fightConfirmChan  chan bool
	fightBgToUi       chan []SkillInfo
	fightUiToBg       chan string
	eventsChannel     chan Event
	length            int
	width             int
	bossEntryRoomNums []int
}

func (l *Labyrinth) Init(p *Player, rooms []*Room, fightConfirm chan bool, fightBgToUi chan []SkillInfo, fightUiToBg chan string, events chan Event) {
	l.p = p
	l.rooms = rooms
	l.fightConfirmChan = fightConfirm
	l.fightBgToUi = fightBgToUi
	l.fightUiToBg = fightUiToBg
	l.eventsChannel = events
}

func (l *Labyrinth) GoToRoom(direction Direction) (int, []Stack) {
	if l.current == nil {
		l.current = l.rooms[l.startingRoomNum]
	} else if int(direction) >= 0 {
		neighbourWall := l.current.GetNeighbours()[getNextDoorNum(int(direction), l.previous)]
		if !neighbourWall.CanGoThrough() {
			return 0, nil
		}
		l.current.p = nil
		if neighbourWall.kind == NextSection {
			l.rooms = *neighbourWall.nextSection
		}
		l.current = neighbourWall.leadsTo
		l.previous = (l.previous + int(direction) + DOORS_PER_ROOM + NEW_NEIGHBOUR_OFFSET) % len(l.current.neighbours)
	}
	l.current.p = l.p
	if l.current.HasEnemies() {
		l.eventsChannel <- FightEvent
		return l.current.StartFight()
	} else {
		defer func() { l.eventsChannel <- NoEvent }()
		return 0, []Stack{}
	}
}

func (l *Labyrinth) GetValues() (int, []Stack) {
	return l.current.GetValues()
}

func (l *Labyrinth) Light() (int, []Stack) {
	if l.current.HasShadowEnemies() {
		l.eventsChannel <- FightEvent
	} else {
		defer func() { l.eventsChannel <- NoEvent }()
	}
	return l.current.Light()
}

func (l *Labyrinth) UnlockChest() (int, []Stack) {
	return l.current.UnlockChest()
}

func (l *Labyrinth) GetNeighbours() map[string]bool {
	res := make(map[string]bool)
	for i, v := range l.current.GetNeighbours() {
		res[DirToStr[getRelativeDirection(i, l.previous)]] = v.CanGoThrough()
	}
	return res
}

func (l *Labyrinth) GetCurrentRoom() *Room {
	return l.current
}

func (l *Labyrinth) GetEventsChan() chan Event {
	return l.eventsChannel
}

func (l *Labyrinth) GetPlayer() *Player {
	return l.p
}

func getNextDoorNum(direction, previous int) int {
	return (direction + previous + 1) % DOORS_PER_ROOM
}

func getRelativeDirection(newDirection, prevDirection int) Direction {
	return Direction((-prevDirection + newDirection + DOORS_PER_ROOM + NEW_NEIGHBOUR_OFFSET) % DOORS_PER_ROOM)
}
