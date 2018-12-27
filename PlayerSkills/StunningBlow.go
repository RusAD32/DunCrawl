package PlayerSkills

import (
	"DunCrawl/Effects"
	. "DunCrawl/Interfaces"
	"fmt"
	"math"
)

type StunningBlow struct {
	name       string
	baseDMG    int
	lvl        int
	maxLvl     int
	curExp     int
	speed      int
	uses       int
	lvlupExps  []int
	wielder    Unit
	targets    []Unit
	lastTarget Unit
	res        []string
}

func (s *StunningBlow) GetRes() string {
	res := s.res[0]
	s.res = s.res[1:]
	return res
}

func (s *StunningBlow) ApplyVoid(res string) {
	s.lastTarget = s.targets[0]
	s.targets = s.targets[1:]
	s.res = append(s.res, res)
}

func (s *StunningBlow) Reset() {
	s.uses = 2
}

func (s *StunningBlow) GetTarget() Unit {
	return s.lastTarget
}

func (s *StunningBlow) GetWielder() Unit {
	return s.wielder
}

func (s *StunningBlow) SetTarget(enemy Unit) {
	s.uses--
	if s.lastTarget == nil {
		s.lastTarget = enemy
	}
	s.targets = append(s.targets, enemy)
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

func (s *StunningBlow) GetSpeed() int {
	return s.speed
}

func (s *StunningBlow) GetName() string {
	return s.name
}

func (s *StunningBlow) GetUses() int {
	return s.uses
}

func (s *StunningBlow) Init(player Unit) Skill {
	s.name = "Stunning Blow"
	s.baseDMG = 3
	s.lvl = 1
	s.maxLvl = 3
	s.curExp = 0
	s.speed = 5
	s.uses = 2
	s.lvlupExps = make([]int, 4)
	s.wielder = player
	s.res = make([]string, 0)
	for i := range s.lvlupExps {
		s.lvlupExps[i] = int(math.Pow(float64(i+2), 2.0) / 3.0)
	}
	return s
}

func (s *StunningBlow) LvlUp() {
	if s.lvl < s.maxLvl && s.curExp >= s.lvlupExps[s.lvl-1] {
		s.curExp -= s.lvlupExps[s.lvl+1]
		s.lvl++
		s.baseDMG = int(math.Pow(3.0, math.Sqrt(float64(s.lvl))))
	} else {
		fmt.Sprintln("Error: Requirements for levelling up skill %s not met", s.name)
	}
}

func (s *StunningBlow) AddExp(amount int) {
	if s.lvl < s.maxLvl {
		s.curExp += amount
		if s.curExp >= s.lvlupExps[s.lvl-1] {
			s.LvlUp()
		}
	}
}
