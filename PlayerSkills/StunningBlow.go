package PlayerSkills

import (
	"../Effects"
	. "../Interfaces"
	"fmt"
	"math"
)

type StunningBlow struct {
	Name       string
	BaseDMG    int
	Lvl        int
	MaxLvl     int
	CurExp     int
	Speed      int
	Uses       int
	LvlupExp   []int
	Wielder    Unit
	Targets    []Unit
	LastTarget Unit
	Res        []string
}

func (s *StunningBlow) GetRes() string {
	res := s.Res[0]
	s.Res = s.Res[1:]
	return res
}

func (s *StunningBlow) ApplyVoid(res string) {
	s.LastTarget = s.Targets[0]
	s.Targets = s.Targets[1:]
	s.Res = append(s.Res, res)
}

func (s *StunningBlow) Reset() {
	s.Uses = 2
}

func (s *StunningBlow) GetTarget() Unit {
	return s.LastTarget
}

func (s *StunningBlow) GetWielder() Unit {
	return s.Wielder
}

func (s *StunningBlow) SetTarget(enemy Unit) {
	s.Uses--
	if s.LastTarget == nil {
		s.LastTarget = enemy
	}
	s.Targets = append(s.Targets, enemy)
}

func (s *StunningBlow) Apply(f *Fight) string {
	equipDmg := 0
	for _, v := range s.Wielder.(*Player).Equipment {
		equipDmg += v.Attack
	}
	s.LastTarget = s.Targets[0]
	s.Targets = s.Targets[1:]
	res := DealDamage(s.Wielder, s.LastTarget, s.BaseDMG+equipDmg)
	s.Res = append(s.Res, res)
	effect := (&Effects.StunEffect{}).Init()
	AddEffect(s.LastTarget, effect)
	return res
}

func (s *StunningBlow) GetSpeed() int {
	return s.Speed
}

func (s *StunningBlow) GetName() string {
	return s.Name
}

func (s *StunningBlow) GetUses() int {
	return s.Uses
}

func (s *StunningBlow) Init(player Unit) Skill {
	s.Name = "Stunning Blow"
	s.BaseDMG = 3
	s.Lvl = 1
	s.MaxLvl = 3
	s.CurExp = 0
	s.Speed = 5
	s.Uses = 2
	s.LvlupExp = make([]int, 4)
	s.Wielder = player
	s.Res = make([]string, 0)
	for i := range s.LvlupExp {
		s.LvlupExp[i] = int(math.Pow(float64(i+2), 2.0) / 3.0)
	}
	return s
}

func (s *StunningBlow) LvlUp() {
	if s.Lvl < s.MaxLvl && s.CurExp >= s.LvlupExp[s.Lvl-1] {
		s.CurExp -= s.LvlupExp[s.Lvl+1]
		s.Lvl++
		s.BaseDMG = int(math.Pow(3.0, math.Sqrt(float64(s.Lvl))))
	} else {
		Inform(fmt.Sprintf("Error: Requirements for levelling up skill %s not met", s.Name))
	}
}

func (s *StunningBlow) AddExp(amount int) {
	if s.Lvl < s.MaxLvl {
		s.CurExp += amount
		if s.CurExp >= s.LvlupExp[s.Lvl-1] {
			s.LvlUp()
		}
	}
}
