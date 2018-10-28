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
	AvailableSlots []Slot
	Name           string
	Defence        int
	Attack         int
	StatsBoost     map[Stat]int
	Effects        []Effect
	Triggerables   []Triggerable
}

func (e Equippable) GetName() string {
	return e.Name
}

func (e Equippable) Use(p *Player, values ...interface{}) {
	length := len(e.AvailableSlots)
	if len(values) > 0 {
		slotNum := values[0].(int)
		if length > 1 {
			prompt := "Choose where to equip:\n"
			for i, v := range e.AvailableSlots {
				prompt += fmt.Sprintf("%d: %s\n", i+1, SlotNames[v])
			}
			p.Equipment[e.AvailableSlots[slotNum]] = e

		} else if length == 1 {
			p.Equipment[e.AvailableSlots[0]] = e
		}
	}
}
