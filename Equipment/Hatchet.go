package Equipment

import . "../Interfaces"

type Hatchet Equippable

func (h *Hatchet) Init() {
	h.Defence = 0
	h.Attack = 3
	h.AvailableSlots = []Slot{MainHand, OffHand}
	h.Effects = []Effect{}
	h.StatsBoost = map[Stat]int{}
	h.Triggerables = []Triggerable{}
}
