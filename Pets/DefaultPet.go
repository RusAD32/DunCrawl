package Pets

import (
	. "DunCrawl/Interfaces"
	. "DunCrawl/NPCSkills"
)

func NewDefaultPet() *Pet {
	d := &Pet{}
	skills := make([]NPCSkill, 0)
	sk1 := NewDogBite(d)
	skills = append(skills, sk1)
	*d = *NewPet(
		Animal,
		skills,
		Usual,
		"Doggy",
		25,
		make(map[Stat]int))
	return d
}
