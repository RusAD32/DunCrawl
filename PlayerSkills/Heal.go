package PlayerSkills

import (
	. "../Interfaces"
	"fmt"
)

type Heal struct {
	hp       int
	lvl      int
	maxLvl   int
	curExp   int
	lvlupExp []int
	speed    int
	name     string
	wielder  Unit
	res      string
}

func (h *Heal) GetRes() string {
	return h.res
}

func (h *Heal) ApplyVoid(res string) {
	h.res = res
}

func (h *Heal) GetTarget() Unit {
	return h.wielder
}

func (h *Heal) Apply(r *Room) string {
	h.res = HealthUp(h.wielder, h.wielder, h.hp)
	return h.res
}

func (h *Heal) GetWielder() Unit {
	return h.wielder
}

func (h *Heal) GetSpeed() int {
	return h.speed
}

func (h *Heal) GetName() string {
	return h.name
}

func (h *Heal) Init(player Unit) Skill {
	h.hp = 8
	h.lvl = 1
	h.curExp = 0
	h.maxLvl = 3
	h.lvlupExp = []int{1, 4}
	h.speed = 3
	h.name = "Heal"
	h.wielder = player
	return h
}

func (h *Heal) LvlUp() {
	if h.lvl < h.maxLvl && h.curExp >= h.lvlupExp[h.lvl-1] {
		h.curExp -= h.lvlupExp[h.lvl+1]
		h.lvl++
		h.hp = int(h.hp * 3 / 2) // Why won't you multiply by 1.5?
	} else {
		fmt.Sprintln("Error: Requirements for levelling up skill %s not met", h.name)
	}
}

func (h *Heal) AddExp(amount int) {
	if h.lvl < h.maxLvl {
		h.curExp += amount
		if h.curExp > h.lvlupExp[h.lvl-1] {
			h.LvlUp()
		}
	}
}
