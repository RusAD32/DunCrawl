package Interfaces

import (
	"fmt"
)

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

func (e *Equippable) Init(availableSlots []Slot, name string, defence, attack int, statsBoosts map[Stat]int, effects []Effect, triggerables []Triggerable) *Equippable {
	//e := new(Equippable)
	e.availableSlots = availableSlots
	e.name = name
	e.defence = defence
	e.attack = attack
	e.statsBoosts = statsBoosts
	e.effects = effects
	e.triggerables = triggerables
	return e
}

func (e *Equippable) GetAttack() int {
	return e.attack
}

func (e Equippable) GetName() string {
	return e.name
}

func (e Equippable) Use(p *Player, values ...interface{}) {
	length := len(e.availableSlots)
	if len(values) > 0 {
		slotNum := values[0].(int)
		if length > 1 {
			prompt := "Choose where to equip:\n"
			for i, v := range e.availableSlots {
				prompt += fmt.Sprintf("%d: %s\n", i+1, SlotNames[v])
			}
			p.equipment[e.availableSlots[slotNum]] = e

		} else if length == 1 {
			p.equipment[e.availableSlots[0]] = e
		}
	}
}
