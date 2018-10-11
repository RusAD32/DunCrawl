package UI

import (
	. "../Interfaces"
	"fmt"
)

func TextFight(p *Player, enemies []*Enemy) {
	f := Fight{}
	uiToBg := make(chan string)
	bgToUi := make(chan []SkillInfo)
	f.Init(p, enemies, bgToUi, uiToBg)
	go f.StartFight()
	for {
		if f.P.CurPhysHP == 0 { // should work, but only theoretically. Maybe handling player death should be different?
			Inform("You died")
			break
		}
		Inform(fmt.Sprintf("Your HP: %d/%d\n", f.P.CurPhysHP, f.P.MaxPhysHP))
		for _, v := range f.Enemies {
			Inform(fmt.Sprintf("%s's HP: %d/%d\n", v.Name, v.CurHP, v.MaxHP))
		}
		selfSkills, ok := <-bgToUi
		if !ok {
			Inform("Fight is over") // here we are supposed to break
			break
		}
		prompt := "Choose a skill to use on yourself\n"
		for i, v := range selfSkills {
			prompt += fmt.Sprintf("%d. %s\n", i+1, v.GetName())
		}
		res := Prompt(prompt, MakeStrRange(1, len(selfSkills)))
		if res == "" {
			Inform("Prompt returned empty string, selfskill")
			return
		}
		uiToBg <- res
		Inform("Select a skill to use on each enemy\n")
		for _, v := range f.Enemies {
			dmgSkills, ok := <-bgToUi
			if !ok {
				Inform("The turn ended in the middle!!")
				break // Не должно!!!
			}
			Inform(fmt.Sprintf("%s. HP: %d/%d\n", v.Name, v.CurHP, v.MaxHP))
			info := "Your skills:\n"
			for i, sk := range dmgSkills {
				// this is not quite safe, but __should__ work
				info += fmt.Sprintf("%d. %s (uses left: %d)\n", i+1, sk.GetName(), sk.(PlayerDmgSkill).GetUses())
			}
			Inform(info)
			dmgSkill := Prompt(v.Name+": ", MakeStrRange(1, len(dmgSkills)))
			if dmgSkill == "" {
				Inform("Prompt returned empty string, dmgskill")
			}
			uiToBg <- dmgSkill
		}
		skillsUsed, ok := <-bgToUi
		if !ok {
			Inform("Something went wrong applying skills")
			break
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
