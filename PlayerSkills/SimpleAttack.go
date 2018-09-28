package PlayerSkills

import (
	. "../Interfaces"
	"fmt"
	"math"
)

type SimpleAttack struct {
	Name     string
	BaseDMG  int
	Lvl      int
	MaxLvl   int
	CurExp   int
	Speed    int
	LvlupExp []int
	Wielder  Unit
	Target   Unit
}

func (s *SimpleAttack) GetTarget() Unit {
	return s.Target
}

func (s *SimpleAttack) GetEnemy() Unit {
	return s.Target
}

func (s *SimpleAttack) GetWielder() Unit {
	return s.Wielder
}

func (s *SimpleAttack) SetTarget(enemy Unit) {
	s.Target = enemy
}

func (s *SimpleAttack) Apply(f *Fight) string {
	equipDmg := 0
	for _, v := range s.Wielder.(*Player).Equipment {
		equipDmg += v.Attack
	}
	return DealDamage(s.Wielder, s.Target, s.BaseDMG+equipDmg)
}

func (s *SimpleAttack) GetSpeed() int {
	return s.Speed
}

func (s *SimpleAttack) GetName() string {
	return s.Name
}

func (s *SimpleAttack) GetUses() int {
	return -1
}

func (s *SimpleAttack) Init(player Unit) {
	s.Name = "Simple Attack"
	s.BaseDMG = 5
	s.Lvl = 1
	s.MaxLvl = 5
	s.CurExp = 0
	s.Speed = 7
	s.LvlupExp = make([]int, 4)
	s.Wielder = player
	for i := range s.LvlupExp {
		s.LvlupExp[i] = int(math.Pow(float64(i+2), 2.0) / 4.0)
	}
}

func (s *SimpleAttack) LvlUp() {
	if s.Lvl < s.MaxLvl && s.CurExp >= s.LvlupExp[s.Lvl-1] {
		s.CurExp -= s.LvlupExp[s.Lvl+1]
		s.Lvl++
		s.BaseDMG = int(math.Pow(5.0, math.Sqrt(float64(s.Lvl))))
	} else {
		Inform(fmt.Sprintf("Error: Requirements for levelling up skill %s not met", s.Name))
	}
}

func (s *SimpleAttack) AddExp(amount int) {
	if s.Lvl < s.MaxLvl {
		s.CurExp += amount
		if s.CurExp >= s.LvlupExp[s.Lvl-1] {
			s.LvlUp()
		}
	}
}
