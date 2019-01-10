package PlayerSkills

import (
	. "DunCrawl/Interfaces"
	"fmt"
)

type Heal struct {
	hp int
	CommonSelfSkill
}

func (h *Heal) Apply(r *Room) string {
	h.res = HealthUp(h.wielder, h.wielder, h.hp)
	return h.res
}

func NewHeal(player Unit) *Heal {
	h := &Heal{}
	h.hp = 8
	h.lvl = 1
	h.curExp = 0
	h.maxLvl = 3
	h.lvlupExp = []int{1, 4}
	h.speed = 3
	h.name = "Heal"
	h.iconPath = "resources/PlayerSkillsIcons/Heal.PNG"
	h.wielder = player
	return h
}

func (h *Heal) LvlUp() {
	if h.lvl < h.maxLvl && h.curExp >= h.lvlupExp[h.lvl-1] {
		h.curExp -= h.lvlupExp[h.lvl+1]
		h.lvl++
		h.hp = int(h.hp * 3 / 2) // Why won't you multiply by 1.5?
	} else {
		fmt.Printf("Error: Requirements for levelling up skill %s not met", h.name)
	}
}
