package Equipment

import . "DunCrawl/Interfaces"

func NewHatchet() *Equippable {
	return NewEquippable(
		[]Slot{MainHand, OffHand},
		"Hatchet",
		0,
		3,
		map[Stat]int{},
		[]Effect{},
		[]Triggerable{})
}
