package Interfaces

import (
	"fmt"
	"strconv"
)

type Fight struct {
	p        *Player
	enemies  []*Enemy
	defeated []*Enemy
	pq       PriorityQueue
}

func (f *Fight) Turn() {
	for _, v := range f.p.GetEffects() {
		v.DecreaseCD()
	}
	for _, en := range f.enemies {
		for _, v := range en.GetEffects() {
			v.DecreaseCD()
		}
	}
	Inform(fmt.Sprintf("Your HP: %d/%d\n", f.p.CurPhysHP, f.p.MaxPhysHP))
	for _, v := range f.enemies {
		Inform(fmt.Sprintf("%s's HP: %d/%d\n", v.Name, v.CurHP, v.MaxHP))
	}
	prompt := "Choose a skill to use on yourself\n"
	for i, v := range f.p.SelfSkills {
		prompt += fmt.Sprintf("%d. %s\n", i, v.GetName())
	}
	res := Prompt(prompt, MakeStrRange(0, len(f.p.SelfSkills)-1))
	var chosenSelfSkill PlayerSelfSkill
	if res == "" {
		Inform("Prompt returned empty string, selfskill")
		return
	}
	skillNum, err := strconv.Atoi(res)
	if err != nil {
		Inform("Prompt returned bad value: " + res)
		return
	}
	chosenSelfSkill = f.p.SelfSkills[skillNum]

	f.pq.Push(chosenSelfSkill)
	info := "Your skills:\n"
	for i, v := range f.p.DmgSkills {
		info += fmt.Sprintf("%d. %s\n", i, v.GetName())
	}
	Inform(info)
	Inform("Select a skill to use on each enemy\n")
	for _, v := range f.enemies {
		dmgSkill := Prompt(v.Name+": ", MakeStrRange(0, len(f.p.DmgSkills)-1))
		if dmgSkill == "" {
			Inform("Prompt returned empty string, dmgskill")
		}
		dmgSkillNum, err := strconv.Atoi(dmgSkill)
		if err != nil {
			Inform("Prompt returned bad value: " + dmgSkill)
			return
		}
		chosenDmgSkill := f.p.DmgSkills[dmgSkillNum]
		chosenDmgSkill.SetTarget(v)
		f.pq.Push(chosenDmgSkill)
		ensk := v.ChooseSkill()
		ensk.SetTarget(f.p)
		f.pq.Push(v.ChooseSkill())
	}
	for f.pq.Len() > 0 {
		sk := f.pq.Pop().(Skill)
		if sk.GetWielder().GetHP() > 0 {
			res := sk.Apply(f)
			Inform(fmt.Sprintf(
				"%s used %s on %s, %s\n",
				sk.GetWielder().GetName(),
				sk.GetName(),
				sk.GetTarget().GetName(),
				res))
		}
	}
	RemoveExpiredEffects(&f.p.Effects)
	for _, en := range f.enemies {
		RemoveExpiredEffects(&en.Effects)
	}
	RemoveDeadEnemies(f)
}

func (f *Fight) StartFight(p *Player, enemies []*Enemy) {
	//heap.Init(&f.pq)
	f.p = p
	f.enemies = enemies
	f.defeated = make([]*Enemy, 0)
	for len(f.enemies) > 0 && p.CurMentHP > 0 && p.CurPhysHP > 0 {
		f.Turn()
	}
}
