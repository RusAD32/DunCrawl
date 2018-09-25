package PlayerSkills

import (
	. "../Interfaces"
	"fmt"
)

type Heal struct {
	HP       int
	Lvl      int
	MaxLvl   int
	CurExp   int
	LvlupExp []int
	Speed    int
	Name     string
}

func (h *Heal) GetSpeed() int {
	return h.Speed
}

func (h *Heal) GetName() string {
	return h.Name
}

func (h *Heal) Init() {
	h.HP = 8
	h.Lvl = 1
	h.CurExp = 0
	h.MaxLvl = 3
	h.LvlupExp = []int{1, 4}
	h.Speed = 3
	h.Name = "Heal"
}

func (h *Heal) LvlUp() {
	if h.Lvl < h.MaxLvl && h.CurExp >= h.LvlupExp[h.Lvl-1] {
		h.CurExp -= h.LvlupExp[h.Lvl+1]
		h.Lvl++
		h.HP = int(h.HP * 3 / 2) // Why won't you multiply by 1.5?
	} else {
		Inform(fmt.Sprintf("Error: Requirements for levelling up skill %s not met", h.Name))
	}
}

func (h *Heal) AddExp(amount int) {
	if h.Lvl < h.MaxLvl {
		h.CurExp += amount
		if h.CurExp > h.LvlupExp[h.Lvl-1] {
			h.LvlUp()
		}
	}
}

func (h *Heal) Apply(wielder *Player) {
	panic("implement me")
}
