package PlayerSkills

import . "DunCrawl/Interfaces"

type CommonDmgSkill struct {
	targets Unit
	res     string
	uses    int
	maxUses int
	CommonPlSkill
}

func (dsk *CommonDmgSkill) Reset() {
	dsk.uses = dsk.maxUses
}

func (dsk *CommonDmgSkill) SetTarget(enemy Unit) {
	dsk.uses--
	dsk.targets = enemy
}

func (dsk *CommonDmgSkill) GetUses() int {
	return dsk.uses
}

func (dsk *CommonDmgSkill) GetRes() string {
	return dsk.res
}

func (dsk *CommonDmgSkill) ApplyVoid(res string) {
	dsk.res = res
}

func (dsk *CommonDmgSkill) GetTarget() Unit {
	return dsk.targets
}

func (dsk *CommonDmgSkill) Copy() PlayerDmgSkill {
	sk := *dsk
	return &sk
}
