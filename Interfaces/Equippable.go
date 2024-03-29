package Interfaces

type Slot int

const (
	Head Slot = iota
	Body
	Hands
	Legs
	Feet
	MainHand
	OffHand
)

var SlotNames = map[Slot]string{
	Head:     "Head",
	Body:     "Body",
	Hands:    "Hands",
	Legs:     "Legs",
	Feet:     "Feet",
	MainHand: "Main hand",
	OffHand:  "Offhand",
}

type Equippable struct {
	availableSlots []Slot
	name           string
	defence        int
	attack         int
	statsBoosts    map[Stat]int
	effects        []Effect
	triggerables   []Triggerable
}

func NewEquippable(availableSlots []Slot, name string, defence int, attack int,
	statsBoosts map[Stat]int, effects []Effect, triggerables []Triggerable) *Equippable {
	return &Equippable{
		availableSlots: availableSlots,
		name:           name,
		defence:        defence,
		attack:         attack,
		statsBoosts:    statsBoosts,
		effects:        effects,
		triggerables:   triggerables,
	}
}

func (e *Equippable) GetAttack() int {
	return e.attack
}

func (e *Equippable) GetName() string {
	return e.name
}

func (e *Equippable) Use(p *Player, values ...interface{}) {
	length := len(e.availableSlots)
	var slotNum int
	if len(values) > 0 {
		slotNum = values[0].(int)
	} else {
		slotNum = 0
	}
	if length > 1 {
		p.equipment[e.availableSlots[slotNum]] = e
	} else if length == 1 {
		p.equipment[e.availableSlots[0]] = e
	}

}

func (e *Equippable) StacksBy() int {
	return 1
}
