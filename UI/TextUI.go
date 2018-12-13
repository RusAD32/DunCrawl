package UI

import (
	. "../Interfaces"
	"fmt"
	"strings"
)

const (
	LightCmd = "light"
	ChestCmd = "chest"
	GotoCmd  = "goto"
	LeftCmd  = "l"
	FwdCmd   = "f"
	RightCmd = "r"
	BackCmd  = "b"
)

var directionMap = map[string]Direction{
	LeftCmd:  Left,
	FwdCmd:   Forward,
	RightCmd: Right,
	BackCmd:  Back,
}

var commands = []string{
	LightCmd,
	ChestCmd,
	GotoCmd,
	LeftCmd,
	RightCmd,
	FwdCmd,
	BackCmd,
}

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
}

func EnterLabyrinth(l *Labyrinth) { //TODO выводить номера комнат, в которые можно перейти, протестить
	next := Direction(-1)
	events := l.GetEventsChan()
	for {
		var money int
		var loot []Carriable
		go func() { money, loot = l.GoToRoom(next) }()
		f := <-events
		if f == FightEvent {
			TextFight(*l.GetCurrentRoom())
			InformLoot(money, loot)
		}
		PrintLabyrinth(l)
		for _, v := range l.GetCurrentRoom().GetNeighbours() {
			fmt.Print(v.CanGoThrough(), " ")
		}
		fmt.Println(l.GetCurrentRoom().Num)
		money, loot = l.GetValues()
		InformLoot(money, loot)
		stayHere := true
		for stayHere {
			cmd, _ := Prompt("Write a command...", commands)
			switch cmd {
			case LightCmd:
				{
					go func() { money, loot = l.Light() }()
					f = <-events
					if f == FightEvent {
						TextFight(*l.GetCurrentRoom())
					}
					InformLoot(money, loot)
				}
			case ChestCmd:
				{
					money, loot = l.UnlockChest()
					InformLoot(money, loot)
				}
			case GotoCmd:
				{
					neighboringRooms := l.GetNeighbours()
					Inform("Which room?\n")
					directions := make([]string, 0)
					for k, v := range neighboringRooms {
						if v {
							Inform(k + "\n")
							directions = append(directions, k)
						}
					}
					v, _ := Prompt("", directions)
					_next, ok := directionMap[strings.ToLower(v)]
					if !ok {
						fmt.Println(v)
					} else {
						next = _next
						stayHere = false
					}
				}
			case LeftCmd, FwdCmd, RightCmd, BackCmd:
				{
					_next, ok := directionMap[strings.ToLower(cmd)]
					if !ok {
						fmt.Println(cmd, " is an unknown direction somehow")
					}
					passable, ok := l.GetNeighbours()[cmd]
					if !ok {
						fmt.Println(cmd, " not a neighbour?")
					} else if !passable {
						Inform("There's a wall in that direction\n")
					} else {
						next = _next
						stayHere = false
					}

				}
			default:
				fmt.Println("Unknown command", cmd)
				break
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
