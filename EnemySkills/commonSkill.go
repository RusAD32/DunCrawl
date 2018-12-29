package EnemySkills

import . "DunCrawl/Interfaces"

type commonSkill struct {
	speed   int
	name    string
	wielder Unit
	target  Unit
	res     string
}

func (esk *commonSkill) GetRes() string {
	return esk.res
}

func (esk *commonSkill) ApplyVoid(res string) {
	esk.res = res
}

func (esk *commonSkill) SetTarget(player Unit) {
	esk.target = player
}

func (esk *commonSkill) GetTarget() Unit {
	return esk.target
}

func (esk *commonSkill) GetWielder() Unit {
	return esk.wielder
}

func (esk *commonSkill) GetSpeed() int {
	return esk.speed
}

func (esk *commonSkill) GetName() string {
	return esk.name
}

func (esk *commonSkill) Init(wielder Unit) Skill { return esk }
func (esk *commonSkill) Apply(r *Room) string    { return "" }
