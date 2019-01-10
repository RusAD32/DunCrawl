package Enemies

import (
	. "DunCrawl/Interfaces"
	"DunCrawl/NPCSkills"
	"fmt"
)

func NewDefaultDog(index int) *Enemy {
	e := &Enemy{}
	skills := make([]NPCSkill, 0)
	sk1 := NPCSkills.NewDogBite(e)
	skills = append(skills, sk1)
	loot := []Lootable{NewLootable("Stuff", 15)}
	*e = *NewEnemy(
		Animal,
		skills,
		make([]*Equippable, 0),
		//make([]Lootable, 0),
		loot,
		make([]Stack, 0),
		Usual,
		fmt.Sprintf("Rabid dog %d", index),
		7,
		make(map[Stat]int))
	return e
}
