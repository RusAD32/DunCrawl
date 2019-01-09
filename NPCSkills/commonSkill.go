package NPCSkills

import . "DunCrawl/Interfaces"

type CommonEnSkill struct {
	speed   int
	name    string
	wielder Unit
	target  Unit
	res     string
}

func (esk *CommonEnSkill) GetRes() string {
	return esk.res
}

func (esk *CommonEnSkill) ApplyVoid(res string) {
	esk.res = res
}

func (esk *CommonEnSkill) SetTarget(player Unit) {
	esk.target = player
}

func (esk *CommonEnSkill) GetTarget() Unit {
	return esk.target
}

func (esk *CommonEnSkill) GetWielder() Unit {
	return esk.wielder
}

func (esk *CommonEnSkill) GetSpeed() int {
	return esk.speed
}

func (esk *CommonEnSkill) GetName() string {
	return esk.name
}

func (esk *CommonEnSkill) Init(wielder Unit) Skill {
	panic("implement me")
}

func (esk *CommonEnSkill) Apply(r *Room) string {
	panic("implement me")
}

func (esk *CommonEnSkill) GetSkillType() SkillType {
	return OppositeSide
}
