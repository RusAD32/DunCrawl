package Interfaces

type WallType int

const (
	Solid WallType = iota
	Door
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
	kind    WallType
	leadsTo *Room
}

func (w *Wall) CanGoThrough() bool {
	return w.kind == Door
}

func (w *Wall) GetNextDoor() *Room {
	return w.leadsTo
}

func ConnectRooms(r1, r2 *Room, d Direction) {
	r1.neighbours[int(d)].leadsTo = r2
	r1.neighbours[int(d)].kind = Door
	r2.neighbours[(int(d)+2)%4].leadsTo = r1
	r2.neighbours[(int(d)+2)%4].kind = Door
}

func LockRooms(r1, r2 *Room, d int) {
	r1.neighbours[d].leadsTo = r2
	r1.neighbours[d].kind = Solid
	r2.neighbours[(d+2)%4].leadsTo = r1
	r2.neighbours[(d+2)%4].kind = Solid
}
