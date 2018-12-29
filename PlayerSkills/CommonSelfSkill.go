package PlayerSkills

import . "DunCrawl/Interfaces"

type CommonSelfSkill struct {
	res string
	CommonPlSkill
}

func (s *CommonSelfSkill) GetRes() string {
	return s.res
}

func (s *CommonSelfSkill) GetTarget() Unit {
	return s.wielder
}

func (s *CommonSelfSkill) ApplyVoid(res string) {
	s.res = res
}
