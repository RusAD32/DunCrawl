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

func (b *CommonPlSkill) GetWielder() Unit {
	return b.wielder
}

func (b *CommonPlSkill) GetSpeed() int {
	return b.speed
}

func (b *CommonPlSkill) GetName() string {
	return b.name
}

func (c *CommonPlSkill) LvlUp() {
	c.lvl++
}

func (s *CommonPlSkill) AddExp(amount int) {
	if s.lvl < s.maxLvl {
		s.curExp += amount
		if s.curExp >= s.lvlupExp[s.lvl-1] {
			s.LvlUp()
		}
	}
}

func (s *CommonPlSkill) GetTarget() Unit         { return nil }
func (s *CommonPlSkill) GetRes() string          { return "" }
func (s *CommonPlSkill) ApplyVoid(string)        {}
func (s *CommonPlSkill) Apply(*Room) string      { return "" }
func (b *CommonPlSkill) Init(wielder Unit) Skill { return nil }
