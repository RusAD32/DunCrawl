package Interfaces

type WallType int

const (
	Solid WallType = iota
	Door
	NextSection
)

type Direction int

const (
	Left Direction = iota
	Up
	Right
	Down
)

var DirToStr = map[Direction]string{
	Left:  "Left",
	Down:  "Down",
	Right: "Right",
	Up:    "Up",
}

type Wall struct {
	kind        WallType
	leadsTo     *Room
	nextSection *[]*Room
}

func (w *Wall) CanGoThrough() bool {
	return w.kind == Door || w.kind == NextSection
}

func (w *Wall) GetNextDoor() *Room {
	return w.leadsTo
}

func ConnectSection(sectionFrom, sectionTo *[]*Room, roomFrom, roomTo *Room, d Direction) {
	roomFrom.neighbours[int(d)].leadsTo = roomTo
	roomFrom.neighbours[int(d)].kind = NextSection
	roomFrom.neighbours[int(d)].nextSection = sectionTo
	roomTo.neighbours[(int(d)+2)%4].leadsTo = roomFrom
	roomTo.neighbours[(int(d)+2)%4].kind = NextSection
	roomTo.neighbours[(int(d)+2)%4].nextSection = sectionFrom
}

func ConnectRooms(r1, r2 *Room, d Direction, kind WallType) {
	r1.neighbours[int(d)].leadsTo = r2
	r1.neighbours[int(d)].kind = kind
	r2.neighbours[(int(d)+2)%4].leadsTo = r1
	r2.neighbours[(int(d)+2)%4].kind = kind
}

func LockRooms(r1, r2 *Room, d int) {
	//r1.neighbours[d].leadsTo = r2
	r1.neighbours[d].kind = Solid
	//r2.neighbours[(d+2)%4].leadsTo = r1
	r2.neighbours[(d+2)%4].kind = Solid
}

func UnockRooms(r1, r2 *Room, d int) {
	//r1.neighbours[d].leadsTo = r2
	r1.neighbours[d].kind = Door
	//r2.neighbours[(d+2)%4].leadsTo = r1
	r2.neighbours[(d+2)%4].kind = Door
}
