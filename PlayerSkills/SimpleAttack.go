package PlayerSkills

import (
	. "../Interfaces"
	"fmt"
	"math"
)

type SimpleAttack struct {
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
	Res        string
}

func (s *SimpleAttack) GetRes() string {
	return s.Res
}

func (s *SimpleAttack) Reset() {
	s.Uses = 4
}

func (s *SimpleAttack) ApplyVoid(res string) {
	s.LastTarget = s.Targets[0]
	s.Targets = s.Targets[1:]
	s.Res = res
}

func (s *SimpleAttack) GetTarget() Unit {
	return s.LastTarget
}

func (s *SimpleAttack) GetWielder() Unit {
	return s.Wielder
}

func (s *SimpleAttack) SetTarget(enemy Unit) {
	s.Uses--
	if s.LastTarget == nil {
		s.LastTarget = enemy
	}
	s.Targets = append(s.Targets, enemy)
}

func (s *SimpleAttack) Apply(f *Fight) string {
	equipDmg := 0
	for _, v := range s.Wielder.(*Player).Equipment {
		equipDmg += v.Attack
	}
	s.LastTarget = s.Targets[0]
	s.Targets = s.Targets[1:]
	s.Res = DealDamage(s.Wielder, s.LastTarget, s.BaseDMG+equipDmg)
	return s.Res
}

func (s *SimpleAttack) GetSpeed() int {
	return s.Speed
}

func (s *SimpleAttack) GetName() string {
	return s.Name
}

func (s *SimpleAttack) GetUses() int {
	return s.Uses
}

func (s *SimpleAttack) Init(player Unit) {
	s.Name = "Simple Attack"
	s.BaseDMG = 5
	s.Lvl = 1
	s.MaxLvl = 5
	s.CurExp = 0
	s.Speed = 7
	s.Uses = 4
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
