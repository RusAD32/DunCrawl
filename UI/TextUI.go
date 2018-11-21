package UI

import (
	. "../Interfaces"
	"fmt"
)

const (
	LIGHT_CMD = "light"
	CHEST_CMD = "chest"
	GOTO_CMD  = "goto"
	LEFT_CMD  = "left"
	UP_CMD    = "up"
	RIGHT_CMD = "right"
	BACK_CMD  = "back"
)

var directionMap = map[string]Direction{
	LEFT_CMD:  Left,
	UP_CMD:    Up,
	RIGHT_CMD: Right,
	BACK_CMD:  Down,
}

var commands = []string{LIGHT_CMD, CHEST_CMD, GOTO_CMD}

func TextFight( /*p *Player, enemies []*Enemy*/ r Room) {
	/*r := Room{}
	uiToBg := make(chan string)
	bgToUi := make(chan []SkillInfo)
	confirm := make(chan bool)
	r.Init(p, enemies, bgToUi, uiToBg, confirm)*/
	uiToBg, bgToUi, confirm := r.GetChannels()
	/*go func() {
		money, loot = r.StartFight()
	}()*/
	for {
		select {
		case alive := <-confirm:
			{
				if !alive {
					Inform("You're dead")
				}
				return
			}
		case selfSkills, ok := <-bgToUi:
			{
				if !ok {
					Inform("Something wrong while getting selfskills")
					return
				}
				if !r.GetPlayer().IsAlive() { // should work, but only theoretically. Maybe handling player death should be different?
					Inform("You died\n")
					return
				}
				Inform(fmt.Sprintf("Your hp: %d/%d\n", r.GetPlayer().GetCurHP(), r.GetPlayer().GetMaxHP()))
				for _, v := range r.GetEnemies() {
					Inform(fmt.Sprintf("%s's hp: %d/%d\n", v.GetName(), v.GetCurHP(), v.GetMaxHP()))
				}
				prompt := "Choose a skill to use on yourself\n"
				for i, v := range selfSkills {
					prompt += fmt.Sprintf("%d. %s\n", i+1, v.GetName())
				}
				res, _ := Prompt(prompt, MakeStrRange(1, len(selfSkills)))
				if res == "" {
					Inform("Prompt returned empty string, selfskill")
					return
				}
				uiToBg <- res[:1]
				Inform("Select a skill to use on each enemy\n")
				for _, v := range r.GetEnemies() {
					dmgSkills, ok := <-bgToUi
					if !ok {
						Inform("The turn ended in the middle!!")
						return // Не должно!!!
					}
					Inform(fmt.Sprintf("%s. hp: %d/%d\n", v.GetName(), v.GetCurHP(), v.GetMaxHP()))
					info := "Your skills:\n"
					for i, sk := range dmgSkills {
						// this is not quite safe, but __should__ work
						info += fmt.Sprintf("%d. %s (uses left: %d)\n", i+1, sk.GetName(), sk.(PlayerDmgSkill).GetUses())
					}
					Inform(info)
					dmgSkill, _ := Prompt(v.GetName()+": ", MakeStrRange(1, len(dmgSkills)))
					if dmgSkill == "" {
						Inform("Prompt returned empty string, dmgskill")
					}
					uiToBg <- dmgSkill[:1]
				}
				skillsUsed, ok := <-bgToUi
				if !ok {
					Inform("Something went wrong applying skills")
					return
				}
				for _, sk := range skillsUsed {
					switch res := sk.GetRes(); res {
					case "stun":
						{
							Inform(fmt.Sprintf(
								"%s tried to use %s on %s, but was stunned\n",
								sk.GetWielder().GetName(),
								sk.GetName(),
								sk.GetTarget().GetName()))
						}
					default:
						{
							Inform(fmt.Sprintf(
								"%s used %s on %s, %s\n",
								sk.GetWielder().GetName(),
								sk.GetName(),
								sk.GetTarget().GetName(),
								res))
						}
					}
				}
			}

		}
	}
	return
}

func EnterLabyrinth(l *Labyrinth) { //TODO выводить номера комнат, в которые можно перейти, протестить
	next := Up
	events := l.GetEventsChan()
	for next >= 0 {
		var money int
		var loot []Carriable
		go func() { money, loot = l.GoToRoom(next) }()
		f := <-events
		if f == FightEvent {
			TextFight(l.GetCurrentRoom())
			InformLoot(money, loot)
		}
		money, loot = l.GetValues()
		InformLoot(money, loot)
		stayHere := true
		for stayHere {
			cmd, _ := Prompt("Write a command...", commands)
			switch cmd {
			case LIGHT_CMD:
				{
					go func() { money, loot = l.Light() }()
					f = <-events
					if f == FightEvent {
						TextFight(l.GetCurrentRoom())
					}
					InformLoot(money, loot)
				}
			case CHEST_CMD:
				{
					money, loot = l.UnlockChest()
					InformLoot(money, loot)
				}
			case GOTO_CMD:
				{
					rooms := l.GetNeighbours()
					Inform("Which room?\n")
					directions := make([]string, 0)
					for k, v := range rooms {
						if v {
							Inform(k + "\n")
							directions = append(directions, k)
						}
					}
					v, _ := Prompt("", directions)
					next = directionMap[v]
					stayHere = false
				}
			default:
				fmt.Println("Unknown command", cmd)
			}

		}
	}
}

func InformLoot(money int, loot []Carriable) {
	Inform(fmt.Sprintf("You found %d money", money))
	if len(loot) > 0 {
		Inform(" and some loot!\n")
	}
	for _, v := range loot {
		Inform(v.GetName() + "\n")
	}
	Inform("\n")
}
