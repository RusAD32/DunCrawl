package Interfaces

type Event int

const (
	NoEvent Event = iota
	FightEvent
)

//TODO давать плееру то, что он в комнатах нашел
//TODO написать Init для лабиринта

type Labyrinth struct {
	Rooms            []*Room
	Current          *Room
	FightConfirmChan chan bool
	FightBgToUi      chan []SkillInfo
	FightUiToBg      chan string
	EventsChannel    chan Event
}

func (l *Labyrinth) GoToRoom(roomNum int) (int, []Carriable) {
	l.Current = l.Rooms[roomNum]
	if l.Current.HasEnemies() {
		l.EventsChannel <- FightEvent
		return l.Current.StartFight()
	} else {
		defer func() { l.EventsChannel <- NoEvent }()
		return 0, []Carriable{}
	}
}

func (l *Labyrinth) GetValues() (int, []Carriable) {
	return l.Current.GetValues()
}

func (l *Labyrinth) Light() (int, []Carriable) {
	if l.Current.HasShadowEnemies() {
		l.EventsChannel <- FightEvent
	} else {
		defer func() { l.EventsChannel <- NoEvent }()
	}
	return l.Current.Light()
}

func (l *Labyrinth) UnlockChest() (int, []Carriable) {
	return l.Current.UnlockChest()
}

func (l *Labyrinth) GetNeighbours() map[string]int {
	res := make(map[string]int, 0)
	for i, v := range l.Current.GetNeighbours() {
		if v.CanGoThrough() {
			res[DirToStr[Direction(i)]] = v.GetNextDoor().Num
		}
	}
	return res
}
