package PlayerSkills

import . "DunCrawl/Interfaces"

type commonSkill struct {
	speed      int
	name       string
	wielder    Unit
	targets    []Unit
	res        []string
	lastTarget Unit
}

func (s *commonSkill) GetRes() string {
	res := s.res[0]
	s.res = s.res[1:]
	return res
}

func (s *commonSkill) ApplyVoid(res string) {
	s.lastTarget = s.targets[0]
	s.targets = s.targets[1:]
	s.res = append(s.res, res)
}

func (b *commonSkill) GetTarget() Unit {
	return b.lastTarget
}

func (b *commonSkill) GetWielder() Unit {
	return b.wielder
}

func (b *commonSkill) GetSpeed() int {
	return b.speed
}

func (b *commonSkill) GetName() string {
	return b.name
}
