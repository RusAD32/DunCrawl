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
	Wielder  Unit
	Res      string
}

func (h *Heal) GetRes() string {
	return h.Res
}

func (h *Heal) ApplyVoid(res string) {
	h.Res = res
}

func (h *Heal) GetTarget() Unit {
	return h.Wielder
}

func (h *Heal) Apply(r *Room) string {
	h.Res = HealthUp(h.Wielder, h.Wielder, h.HP)
	return h.Res
}

func (h *Heal) GetWielder() Unit {
	return h.Wielder
}

func (h *Heal) GetSpeed() int {
	return h.Speed
}

func (h *Heal) GetName() string {
	return h.Name
}

func (h *Heal) Init(player Unit) Skill {
	h.HP = 8
	h.Lvl = 1
	h.CurExp = 0
	h.MaxLvl = 3
	h.LvlupExp = []int{1, 4}
	h.Speed = 3
	h.Name = "Heal"
	h.Wielder = player
	return h
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
