package PlayerSkills

import (
	. "DunCrawl/Interfaces"
	"fmt"
	"math"
)

type SimpleAttack struct {
	baseDMG  int
	lvl      int
	maxLvl   int
	curExp   int
	lvlupExp []int
	CommonDmgSkill
}

func (s *SimpleAttack) Apply(r *Room) string {
	equipDmg := 0
	for _, v := range s.wielder.(*Player).GetEquipment() {
		equipDmg += v.GetAttack()
	}
	/*s.lastTarget = s.targets[0]
	s.targets = s.targets[1:]*/
	res := DealDamage(s.wielder, s.targets, s.baseDMG+equipDmg)
	s.res = res
	return res
}

func NewSimpleAttack(player Unit) *SimpleAttack {
	s := &SimpleAttack{}
	s.name = "Simple attack"
	s.iconPath = "resources/PlayerSkillsIcons/SimpleAttack.PNG"
	s.baseDMG = 5
	s.lvl = 1
	s.maxLvl = 5
	s.curExp = 0
	s.speed = 7
	s.uses = 4
	s.lvlupExp = make([]int, 4)
	s.wielder = player
	s.maxUses = 4
	for i := range s.lvlupExp {
		s.lvlupExp[i] = int(math.Pow(float64(i+2), 2.0) / 4.0)
	}
	return s
}

func (s *SimpleAttack) LvlUp() {
	if s.lvl < s.maxLvl && s.curExp >= s.lvlupExp[s.lvl-1] {
		s.curExp -= s.lvlupExp[s.lvl+1]
		s.lvl++
		s.baseDMG = int(math.Pow(5.0, math.Sqrt(float64(s.lvl))))
	} else {
		fmt.Printf("Error: Requirements for levelling up skill %s not met", s.name)
	}
}

func (dsk *SimpleAttack) Copy() PlayerDmgSkill {
	sk := *dsk
	return &sk
}
