package Interfaces

import "fmt"

type Event int

const (
	NoEvent Event = iota
	FightEvent
)

//TODO давать плееру то, что он в комнатах нашел
//TODO написать Init для лабиринта

type Labyrinth struct {
	p                *Player
	rooms            []*Room
	startingRoomNum  int
	current          *Room
	previous         *Room
	fightConfirmChan chan bool
	fightBgToUi      chan []SkillInfo
	fightUiToBg      chan string
	eventsChannel    chan Event
	length           int
	width            int
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
	} else {
		/*prevNeighbourNum := 0
		if l.previous != nil {
			for i, v := range l.current.GetNeighbours() {
				if v.leadsTo == l.previous {
					prevNeighbourNum = i
				}
			}
		}*/
		l.previous = l.current
		l.current.p = nil
		l.current = l.current.GetNeighbours()[int(direction)].leadsTo
	}
	for i, v := range l.current.neighbours {
		if v.CanGoThrough() {
			fmt.Println(i, v.leadsTo.Num)
		}
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
			res[DirToStr[Direction(i)]] = v.leadsTo.Num
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

func rotate(walls []*Wall, prev *Room) []*Wall {
	return walls
	startNum := 0
	for i, v := range walls {
		if v.leadsTo == prev {
			startNum = (i + 1) % len(walls)
		}
	}
	return append(walls[startNum:], walls[:startNum]...)
}
