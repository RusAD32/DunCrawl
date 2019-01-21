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
	loot := []*Lootable{NewLootable("Stuff", 15)}
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
		make(map[Stat]int),
		[]string{
			"resources/EnemySprites/DefaultDog/Idle/Idle_000.png",
			"resources/EnemySprites/DefaultDog/Idle/Idle_001.png",
			"resources/EnemySprites/DefaultDog/Idle/Idle_002.png",
			"resources/EnemySprites/DefaultDog/Idle/Idle_003.png",
			"resources/EnemySprites/DefaultDog/Idle/Idle_004.png",
			"resources/EnemySprites/DefaultDog/Idle/Idle_005.png",
			"resources/EnemySprites/DefaultDog/Idle/Idle_006.png",
			"resources/EnemySprites/DefaultDog/Idle/Idle_007.png",
			"resources/EnemySprites/DefaultDog/Idle/Idle_008.png",
			"resources/EnemySprites/DefaultDog/Idle/Idle_009.png",
			"resources/EnemySprites/DefaultDog/Idle/Idle_010.png",
			"resources/EnemySprites/DefaultDog/Idle/Idle_011.png",
		},
		[]string{
			"resources/EnemySprites/DefaultDog/Biting/Biting_000.png",
			"resources/EnemySprites/DefaultDog/Biting/Biting_001.png",
			"resources/EnemySprites/DefaultDog/Biting/Biting_002.png",
			"resources/EnemySprites/DefaultDog/Biting/Biting_003.png",
			"resources/EnemySprites/DefaultDog/Biting/Biting_004.png",
			"resources/EnemySprites/DefaultDog/Biting/Biting_005.png",
			"resources/EnemySprites/DefaultDog/Biting/Biting_006.png",
			"resources/EnemySprites/DefaultDog/Biting/Biting_007.png",
			"resources/EnemySprites/DefaultDog/Biting/Biting_008.png",
			"resources/EnemySprites/DefaultDog/Biting/Biting_009.png",
			"resources/EnemySprites/DefaultDog/Biting/Biting_010.png",
			"resources/EnemySprites/DefaultDog/Biting/Biting_011.png",
		},
		[]string{
			"resources/EnemySprites/DefaultDog/Hurt/Hurt_000.png",
			"resources/EnemySprites/DefaultDog/Hurt/Hurt_001.png",
			"resources/EnemySprites/DefaultDog/Hurt/Hurt_002.png",
			"resources/EnemySprites/DefaultDog/Hurt/Hurt_003.png",
			"resources/EnemySprites/DefaultDog/Hurt/Hurt_004.png",
			"resources/EnemySprites/DefaultDog/Hurt/Hurt_005.png",
			"resources/EnemySprites/DefaultDog/Hurt/Hurt_006.png",
			"resources/EnemySprites/DefaultDog/Hurt/Hurt_007.png",
			"resources/EnemySprites/DefaultDog/Hurt/Hurt_008.png",
			"resources/EnemySprites/DefaultDog/Hurt/Hurt_009.png",
			"resources/EnemySprites/DefaultDog/Hurt/Hurt_010.png",
			"resources/EnemySprites/DefaultDog/Hurt/Hurt_011.png",
		},
		[]string{
			"resources/EnemySprites/DefaultDog/Dying/Dying_000.png",
			"resources/EnemySprites/DefaultDog/Dying/Dying_001.png",
			"resources/EnemySprites/DefaultDog/Dying/Dying_002.png",
			"resources/EnemySprites/DefaultDog/Dying/Dying_003.png",
			"resources/EnemySprites/DefaultDog/Dying/Dying_004.png",
			"resources/EnemySprites/DefaultDog/Dying/Dying_005.png",
			"resources/EnemySprites/DefaultDog/Dying/Dying_006.png",
			"resources/EnemySprites/DefaultDog/Dying/Dying_007.png",
			"resources/EnemySprites/DefaultDog/Dying/Dying_008.png",
			"resources/EnemySprites/DefaultDog/Dying/Dying_009.png",
			"resources/EnemySprites/DefaultDog/Dying/Dying_010.png",
			"resources/EnemySprites/DefaultDog/Dying/Dying_011.png",
			"resources/EnemySprites/DefaultDog/Dying/Dying_012.png",
			"resources/EnemySprites/DefaultDog/Dying/Dying_013.png",
			"resources/EnemySprites/DefaultDog/Dying/Dying_014.png",
		})
	return e
}
