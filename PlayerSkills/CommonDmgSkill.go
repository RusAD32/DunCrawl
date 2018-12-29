package PlayerSkills

import . "DunCrawl/Interfaces"

type CommonDmgSkill struct {
	targets    []Unit
	res        []string
	lastTarget Unit
	uses       int
	CommonPlSkill
}

func (dsk *CommonDmgSkill) Reset() {
	dsk.uses = 4
}

func (dsk *CommonDmgSkill) SetTarget(enemy Unit) {
	dsk.uses--
	if dsk.lastTarget == nil {
		dsk.lastTarget = enemy
	}
	dsk.targets = append(dsk.targets, enemy)
}

func (dsk *CommonDmgSkill) GetUses() int {
	return dsk.uses
}

func (dsk *CommonDmgSkill) GetRes() string {
	res := dsk.res[0]
	dsk.res = dsk.res[1:]
	return res
}

func (dsk *CommonDmgSkill) ApplyVoid(res string) {
	dsk.lastTarget = dsk.targets[0]
	dsk.targets = dsk.targets[1:]
	dsk.res = append(dsk.res, res)
}

func (dsk *CommonDmgSkill) GetTarget() Unit {
	return dsk.lastTarget
}
