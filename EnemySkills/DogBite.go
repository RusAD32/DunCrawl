package EnemySkills

import . "DunCrawl/Interfaces"

type DogBite struct {
	baseDmg int
	CommonEnSkill
}

func (b *DogBite) Init(enemy Unit) Skill {
	b.baseDmg = 5
	b.speed = 6
	b.name = "Bite"
	b.wielder = enemy
	b.res = ""
	return b
}

func (b *DogBite) Apply(r *Room) string {
	if b.wielder.GetHP() > 0 {
		b.res = DealDamage(b.wielder, b.target, b.baseDmg)
	}
	return b.res
}
