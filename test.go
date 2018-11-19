package main

import (
	"./EnemySkills"
	. "./Equipment"
	. "./Interfaces"
	"./PlayerSkills"
	"./UI"
)

func main() {
	p := GetDefaultPlayer()
	h := Hatchet{}
	h.Init()
	p.Equip(Equippable(h), MainHand)
	heal := PlayerSkills.Heal{}
	heal.Init(&p)
	p.AddSelfSkill(&heal)
	cntr := PlayerSkills.Counter{}
	cntr.Init(&p)
	p.AddSelfSkill(&cntr)
	atk := PlayerSkills.SimpleAttack{}
	atk.Init(&p)
	p.AddDmgSkill(&atk)
	stn := PlayerSkills.StunningBlow{}
	stn.Init(&p)
	p.AddDmgSkill(&stn)

	//dog := GetDefaultEnemy(index)
	enemies := make([]Enemy, 4)
	for i := range enemies {
		enemies[i] = GetDefaultEnemy(i)
		//enemies[i].name += fmt.Sprintf(" %d", i)
		bite := EnemySkills.DogBite{}
		bite.Init(&enemies[i])
		enemies[i].AddSkill(&bite)
		//enemies[i].skills = append(enemies[i].skills, &bite)
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
	r.Init(ptrenemies, bgToUi, uiToBg, confirm, 0)
	r2 := Room{}
	r2.Init(make([]*Enemy, 0), bgToUi, uiToBg, confirm, 1)
	r2.AddLoot(GenerateLootable("Stuff", 10))
	ch := Chest{
		nil,
		[]Lootable{GenerateLootable("Different stuff", 40)},
		make([]Carriable, 0),
	}
	r2.SetChest(&ch)
	r.AddShadowLoot(GenerateLootable("Other stuff", 200))
	l := Labyrinth{}
	rooms := make([]*Room, 2)
	rooms[0] = &r
	rooms[1] = &r2
	ConnectRooms(&r, &r2, Left)
	l.Init(&p, rooms, confirm, bgToUi, uiToBg, events)
	//	UI.TextFight(&r)
	UI.EnterLabyrinth(&l)
}
