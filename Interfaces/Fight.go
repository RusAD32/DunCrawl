package Interfaces

import (
	"container/heap"
	"fmt"
	"strconv"
)

type Fight struct {
	p       *Player
	enemies []*Enemy
	pq      PriorityQueue
}

func (f *Fight) Turn(p *Player, enemies []*Enemy) {
	prompt := "Choose a skill to use on yourself\n"
	for i, v := range p.SelfSkills {
		prompt += fmt.Sprintf("%d. %s\n", i, v.GetName())
	}
	res := Prompt(prompt, MakeStrRange(1, len(p.SelfSkills)-1))
	var chosenSelfSkill PlayerSelfSkill
	if res != "" {
		Inform("Prompt returned empty string, selfskill")
		return
	}
	skillNum, err := strconv.Atoi(res)
	if err != nil {
		Inform("Prompt returned bad value: " + res)
		return
	}
	chosenSelfSkill = p.SelfSkills[skillNum]

	f.pq.Push(chosenSelfSkill)
	info := "Your skills:"
	for i, v := range p.DmgSkills {
		info += fmt.Sprintf("%d. %s\n", i, v.GetName())
	}
	Inform(info)
	Inform("Select a skill to use each enemy")
	for _, v := range enemies {
		dmgSkill := Prompt(v.Name, MakeStrRange(1, len(p.DmgSkills)-1))
		if dmgSkill == "" {
			Inform("Prompt returned empty string, dmgskill")
		}
		dmgSkillNum, err := strconv.Atoi(dmgSkill)
		if err != nil {
			Inform("Prompt returned bad value: " + dmgSkill)
			return
		}
		chosenDmgSkill := p.DmgSkills[dmgSkillNum]
		chosenDmgSkill.SetTarget(v)
		f.pq.Push(chosenDmgSkill)
		f.pq.Push(v.ChooseSkill())
	}

	for sk := f.pq.Pop().(Skill); f.pq.Len() > 0; sk = f.pq.Pop().(Skill) {
		res := sk.Apply(f)
		Inform(fmt.Sprintf(
			"%s used %s on %s, %s",
			sk.GetWielder().GetName(),
			sk.GetName(),
			sk.GetTarget().GetName(),
			res))
	}
	//TODO timeout the effects
}

func (f *Fight) StartFight(p *Player, enemies []*Enemy) {
	heap.Init(&f.pq)
	for len(enemies) > 0 && p.CurMentHP > 0 && p.CurPhysHP > 0 {

	}
}
