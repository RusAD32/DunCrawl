package Pets

import (
	"DunCrawl/Effects"
	. "DunCrawl/EnemySkills"
	. "DunCrawl/Interfaces"
)

type DefaultPet Pet

func (d *DefaultPet) Init() *Pet {
	skills := make([]NPCSkill, 0)
	sk1 := new(DogBite).Init(d)
	skills = append(skills, sk1)
	e := new(Pet).Initialize(
		Animal,
		skills,
		Usual,
		"Doggy",
		25,
		make(map[Stat]int))
	*d = *(*DefaultPet)(e)
	d.AddEffect(new(Effects.CounterEff).Init())
	return (*Pet)(d)
}
