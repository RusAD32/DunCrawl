package UI

import (
	. "../Interfaces"
	"fmt"
	"strconv"
	"strings"
)

const (
	LIGHT_CMD = "light"
	CHEST_CMD = "chest"
	GOTO_CMD  = "goto"
)

var commands = []string{LIGHT_CMD, CHEST_CMD, GOTO_CMD}

func TextFight( /*p *Player, enemies []*Enemy*/ r *Room) {
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
				if r.P.CurPhysHP == 0 { // should work, but only theoretically. Maybe handling player death should be different?
					Inform("You died")
					return
				}
				Inform(fmt.Sprintf("Your HP: %d/%d\n", r.P.CurPhysHP, r.P.MaxPhysHP))
				for _, v := range r.Enemies {
					Inform(fmt.Sprintf("%s's HP: %d/%d\n", v.Name, v.CurHP, v.MaxHP))
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
				for _, v := range r.Enemies {
					dmgSkills, ok := <-bgToUi
					if !ok {
						Inform("The turn ended in the middle!!")
						return // Не должно!!!
					}
					Inform(fmt.Sprintf("%s. HP: %d/%d\n", v.Name, v.CurHP, v.MaxHP))
					info := "Your skills:\n"
					for i, sk := range dmgSkills {
						// this is not quite safe, but __should__ work
						info += fmt.Sprintf("%d. %s (uses left: %d)\n", i+1, sk.GetName(), sk.(PlayerDmgSkill).GetUses())
					}
					Inform(info)
					dmgSkill, _ := Prompt(v.Name+": ", MakeStrRange(1, len(dmgSkills)))
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
	next := 0
	events := l.EventsChannel
	for next >= 0 {
		var money int
		var loot []Carriable
		go func() { money, loot = l.GoToRoom(next) }()
		f := <-events
		if f == FightEvent {
			TextFight(l.Current)
			Inform(fmt.Sprintf("You got %d gold and some loot from the fight\n", money))
			fmt.Println(loot) //TODO убрать заглушку, выписывать все их имена через Inform
		}
		money, loot = l.GetValues()
		Inform(fmt.Sprintf("You got %d gold and some loot from the fight\n", money))
		fmt.Println(loot) //TODO убрать заглушку, выписывать все их имена через Inform
		stayhere := true
		for stayhere {
			cmd, cmd_ext := Prompt("Write a command... ", commands)
			switch cmd {
			case LIGHT_CMD:
				{
					go func() { money, loot = l.Light() }()
					f = <-events
					if f == FightEvent {
						TextFight(l.Current)
					}
					Inform(fmt.Sprintf("You got %d gold and some loot from the fight\n", money))
					fmt.Println(loot) //TODO убрать заглушку, выписывать все их имена через Inform
				}
			case CHEST_CMD:
				{
					money, loot = l.UnlockChest()
					Inform(fmt.Sprintf("You got %d gold and some loot from the fight\n", money))
					fmt.Println(loot) //TODO убрать заглушку, выписывать все их имена через Inform
				}
			case GOTO_CMD:
				{
					next, _ = strconv.Atoi(strings.Split(cmd_ext, " ")[1])
					stayhere = false
				}
			default:
				fmt.Println("Unknown command", cmd, cmd_ext)
			}

		}
	}
}
