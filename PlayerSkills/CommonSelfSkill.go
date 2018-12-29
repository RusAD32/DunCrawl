package PlayerSkills

import . "DunCrawl/Interfaces"

type CommonSelfSkill struct {
	res string
	CommonPlSkill
}

func (ssk *CommonSelfSkill) GetRes() string {
	return ssk.res
}

func (ssk *CommonSelfSkill) GetTarget() Unit {
	return ssk.wielder
}

func (ssk *CommonSelfSkill) ApplyVoid(res string) {
	ssk.res = res
}
