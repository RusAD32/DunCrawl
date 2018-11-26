package Interfaces

type Event int

const (
	NoEvent Event = iota
	FightEvent
)

//TODO давать плееру то, что он в комнатах нашел
//TODO написать Init для лабиринта

type Labyrinth struct {
	p                 *Player
	rooms             []*Room
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

func (l *Labyrinth) GoToRoom(direction Direction) (int, []Carriable) {
	if l.current == nil {
		l.current = l.rooms[l.startingRoomNum]
	} else if int(direction) >= 0 {
		l.current.p = nil
		l.current = l.current.GetNeighbours()[(int(direction)+l.previous+1)%4].leadsTo
		l.previous = (l.previous + int(direction) + 3) % len(l.current.neighbours)
	}
	l.current.p = l.p
	if l.current.HasEnemies() {
		l.eventsChannel <- FightEvent
		return l.current.StartFight()
	} else {
		defer func() { l.eventsChannel <- NoEvent }()
		return 0, []Carriable{}
	}
}

func (l *Labyrinth) GetValues() (int, []Carriable) {
	return l.current.GetValues()
}

func (l *Labyrinth) Light() (int, []Carriable) {
	if l.current.HasShadowEnemies() {
		l.eventsChannel <- FightEvent
	} else {
		defer func() { l.eventsChannel <- NoEvent }()
	}
	return l.current.Light()
}

func (l *Labyrinth) UnlockChest() (int, []Carriable) {
	return l.current.UnlockChest()
}

func (l *Labyrinth) GetNeighbours() map[string]int {
	res := make(map[string]int, 0)
	for i, v := range l.current.GetNeighbours() {
		if v.CanGoThrough() {
			res[DirToStr[Direction((-l.previous+i+3)%4)]] = v.leadsTo.Num
		}
	}
	return res
}

func (l *Labyrinth) GetCurrentRoom() Room {
	return *l.current
}

func (l *Labyrinth) GetEventsChan() chan Event {
	return l.eventsChannel
}
