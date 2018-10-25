package main

import (
	"./EnemySkills"
	. "./Equipment"
	. "./Interfaces"
	"./PlayerSkills"
	"./UI"
	"fmt"
)

func main() {
	p := Player{
		Stats:           map[Stat]int{},
		Equipment:       map[Slot]Equippable{},
		Inventory:       []Carriable{},
		DmgSkills:       []PlayerDmgSkill{},
		SelfSkills:      []PlayerSelfSkill{},
		Effects:         []Effect{},
		DmgTakenTrigger: TriggerInit(),
		CurPhysHP:       100,
		MaxPhysHP:       100,
		Lvl:             1,
		Exp:             0,
		CurMentHP:       100,
		MaxMentHP:       100,
	}
	h := Hatchet{}
	h.Init()
	p.Equipment[MainHand] = Equippable(h)
	heal := PlayerSkills.Heal{}
	heal.Init(&p)
	p.SelfSkills = append(p.SelfSkills, &heal)
	cntr := PlayerSkills.Counter{}
	cntr.Init(&p)
	p.SelfSkills = append(p.SelfSkills, &cntr)
	atk := PlayerSkills.SimpleAttack{}
	atk.Init(&p)
	p.DmgSkills = append(p.DmgSkills, &atk)
	stn := PlayerSkills.StunningBlow{}
	stn.Init(&p)
	p.DmgSkills = append(p.DmgSkills, &stn)
	dog := Enemy{
		Type:            Animal,
		Name:            "Rabid dog",
		Skills:          []EnemySkill{},
		Effects:         []Effect{},
		Equipment:       []Equippable{},
		DmgTakenTrigger: TriggerInit(),
		AILevel:         Usual,
		CurHP:           15,
		MaxHP:           15,
	}
	enemies := make([]Enemy, 4)
	for i := range enemies {
		enemies[i] = dog
		enemies[i].Name += fmt.Sprintf(" %d", i)
		bite := EnemySkills.DogBite{}
		bite.Init(&enemies[i])
		enemies[i].Skills = append(enemies[i].Skills, &bite)
	}
	ptrenemies := make([]*Enemy, 4)
	for i := range enemies {
		ptrenemies[i] = &enemies[i]
	}

	//f := Room{&p, *[]*Enemy{&dog, &dog2, &dog3, &dog4}, PriorityQueue, 0{}}
	r := Room{}
	uiToBg := make(chan string)
	bgToUi := make(chan []SkillInfo)
	confirm := make(chan bool)
	events := make(chan Event)
	r.Init(ptrenemies, bgToUi, uiToBg, confirm)
	r2 := Room{}
	r2.Init(make([]*Enemy, 0), bgToUi, uiToBg, confirm)
	r2.Loot = make([]Lootable, 1)
	r2.Loot[0] = Lootable{
		Name:  "Stuff",
		Value: 10,
	}
	l := Labyrinth{}
	rooms := make([]*Room, 2)
	rooms[0] = &r
	rooms[1] = &r2
	l.Init(&p, rooms, confirm, bgToUi, uiToBg, events)
	//	UI.TextFight(&r)
	UI.EnterLabyrinth(&l)
}
