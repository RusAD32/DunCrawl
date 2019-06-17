package NPCSkills

import (
	. "DunCrawl/Interfaces"
)

type DogBite struct {
	baseDmg int
	CommonEnSkill
}

func NewDogBite(enemy Unit) NPCSkill {
	b := &DogBite{}
	b.baseDmg = 10
	b.speed = 6
	b.name = "Bite"
	b.wielder = enemy
	b.iconPath = "resources/NPCSkillsIcons/DogBite.PNG"
	b.res = ""
	return b
}

func (b *DogBite) Apply(r *Room) string {
	if b.wielder.GetHP() > 0 {
		b.res = DealDamage(b.wielder, b.target, b.baseDmg)
	}
	return b.res
}
