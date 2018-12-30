package PlayerSkills

import (
	"DunCrawl/Effects"
	. "DunCrawl/Interfaces"
	"fmt"
	"math"
)

type StunningBlow struct {
	baseDMG int
	CommonDmgSkill
}

func (s *StunningBlow) Apply(r *Room) string {
	equipDmg := 0
	for _, v := range s.wielder.(*Player).GetEquipment() {
		equipDmg += v.GetAttack()
	}
	s.lastTarget = s.targets[0]
	s.targets = s.targets[1:]
	res := DealDamage(s.wielder, s.lastTarget, s.baseDMG+equipDmg)
	s.res = append(s.res, res)
	effect := (&Effects.StunEffect{}).Init()
	AddEffect(s.lastTarget, effect)
	return res
}

func (s *StunningBlow) Init(player Unit) Skill {
	s.name = "Stunning Blow"
	s.baseDMG = 3
	s.lvl = 1
	s.maxLvl = 3
	s.curExp = 0
	s.speed = 5
	s.uses = 2
	s.lvlupExp = make([]int, 4)
	s.wielder = player
	s.res = make([]string, 0)
	for i := range s.lvlupExp {
		s.lvlupExp[i] = int(math.Pow(float64(i+2), 2.0) / 3.0)
	}
	return s
}

func (s *StunningBlow) LvlUp() {
	if s.lvl < s.maxLvl && s.curExp >= s.lvlupExp[s.lvl-1] {
		s.curExp -= s.lvlupExp[s.lvl+1]
		s.lvl++
		s.baseDMG = int(math.Pow(3.0, math.Sqrt(float64(s.lvl))))
	} else {
		fmt.Sprintln("Error: Requirements for levelling up skill %s not met", s.name)
	}
}
