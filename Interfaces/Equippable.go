package Interfaces

import (
	"fmt"
	"strconv"
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
	Defence        int
	Attack         int
	StatsBoost     map[Stat]int
	Effects        []Effect
	Triggerables   []Triggerable
}

func (e Equippable) Use(p *Player) {
	length := len(e.AvailableSlots)
	if length > 1 {
		prompt := "Choose where to equip:\n"
		for i, v := range e.AvailableSlots {
			prompt += fmt.Sprintf("%d: %s\n", i+1, SlotNames[v])
		}
		reqInput := MakeStrRange(1, length)
		res, _ := Prompt(prompt, reqInput)
		if res != "" {
			slotNum, err := strconv.Atoi(res)
			if err != nil {
				Errors.Write([]byte(err.Error()))
				return
			}
			p.Equipment[e.AvailableSlots[slotNum]] = e
		}
	} else if length == 1 {
		p.Equipment[e.AvailableSlots[0]] = e
	}
}
