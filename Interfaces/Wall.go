package Interfaces

type WallType int

const (
	Solid WallType = iota
	Door
)

type Direction int

const (
	Left Direction = iota
	Down
	Right
	Up
)

var DirToStr = map[Direction]string{
	Left:  "Left",
	Down:  "Down",
	Right: "Right",
	Up:    "Up",
}

type Wall struct {
	kind    WallType
	LeadsTo *Room
}

func (w *Wall) CanGoThrough() bool {
	return w.kind == Door
}

func (w *Wall) GetNextDoor() *Room {
	return w.LeadsTo
}

func ConnectRooms(r1, r2 *Room, d Direction) {
	r1.Neighbours[int(d)].LeadsTo = r2
	r1.Neighbours[int(d)].kind = Door
	r2.Neighbours[(int(d)+2)%4].LeadsTo = r1
	r2.Neighbours[(int(d)+2)%4].kind = Door
}
