package PlayerSkills

import . "DunCrawl/Interfaces"

type CommonDmgSkill struct {
	targets    []Unit
	res        []string
	lastTarget Unit
	uses       int
	CommonPlSkill
}

func (s *CommonDmgSkill) Reset() {
	s.uses = 4
}

func (s *CommonDmgSkill) SetTarget(enemy Unit) {
	s.uses--
	if s.lastTarget == nil {
		s.lastTarget = enemy
	}
	s.targets = append(s.targets, enemy)
}

func (s *CommonDmgSkill) GetUses() int {
	return s.uses
}

func (s *CommonDmgSkill) GetRes() string {
	res := s.res[0]
	s.res = s.res[1:]
	return res
}

func (s *CommonDmgSkill) ApplyVoid(res string) {
	s.lastTarget = s.targets[0]
	s.targets = s.targets[1:]
	s.res = append(s.res, res)
}

func (b *CommonDmgSkill) GetTarget() Unit {
	return b.lastTarget
}
