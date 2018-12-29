package EnemySkills

import . "DunCrawl/Interfaces"

type commonSkill struct {
	speed   int
	name    string
	wielder Unit
	target  Unit
	res     string
}

func (b *commonSkill) GetRes() string {
	return b.res
}

func (b *commonSkill) ApplyVoid(res string) {
	b.res = res
}

func (b *commonSkill) SetTarget(player Unit) {
	b.target = player
}

func (b *commonSkill) GetTarget() Unit {
	return b.target
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
