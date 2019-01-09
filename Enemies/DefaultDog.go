package Enemies

import (
	. "DunCrawl/Interfaces"
	"DunCrawl/NPCSkills"
	"fmt"
)

type DefaultDog Enemy

func (d *DefaultDog) Init(index int) *Enemy {
	skills := make([]NPCSkill, 0)
	sk1 := new(NPCSkills.DogBite).Init(d)
	skills = append(skills, sk1)
	loot := []Lootable{GenerateLootable("Stuff", 15)}
	e := new(Enemy).Initialize(
		Animal,
		skills,
		make([]Equippable, 0),
		//make([]Lootable, 0),
		loot,
		make([]Stack, 0),
		Usual,
		fmt.Sprintf("Rabid dog %d", index),
		7,
		make(map[Stat]int))
	*d = *(*DefaultDog)(e)
	return (*Enemy)(d)
}
