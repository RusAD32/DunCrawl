package Equipment

import . "DunCrawl/Interfaces"

type Hatchet Equippable

func (h *Hatchet) Init() {
	e := new(Equippable)
	e.Init([]Slot{MainHand, OffHand}, "Hatchet", 0, 3, map[Stat]int{}, []Effect{}, []Triggerable{})
	*h = Hatchet(*e)
}
