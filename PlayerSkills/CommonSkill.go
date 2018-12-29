package PlayerSkills

import . "DunCrawl/Interfaces"

type CommonPlSkill struct {
	lvl      int
	maxLvl   int
	curExp   int
	lvlupExp []int
	speed    int
	name     string
	wielder  Unit
}

func (psk *CommonPlSkill) GetWielder() Unit {
	return psk.wielder
}

func (psk *CommonPlSkill) GetSpeed() int {
	return psk.speed
}

func (psk *CommonPlSkill) GetName() string {
	return psk.name
}

func (psk *CommonPlSkill) LvlUp() {
	psk.lvl++
}

func (psk *CommonPlSkill) AddExp(amount int) {
	if psk.lvl < psk.maxLvl {
		psk.curExp += amount
		if psk.curExp >= psk.lvlupExp[psk.lvl-1] {
			psk.LvlUp()
		}
	}
}

func (psk *CommonPlSkill) GetTarget() Unit         { return nil }
func (psk *CommonPlSkill) GetRes() string          { return "" }
func (psk *CommonPlSkill) ApplyVoid(string)        {}
func (psk *CommonPlSkill) Apply(*Room) string      { return "" }
func (psk *CommonPlSkill) Init(wielder Unit) Skill { return nil }
