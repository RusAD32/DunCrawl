package Interfaces

import (
	"fmt"
	"strconv"
)

type Fight struct {
	p                *Player
	enemies          []*Enemy
	defeated         []*Enemy
	pq               PriorityQueue
	TurnStartTrigger Trigger
	TurnEndTrigger   Trigger
	EnemyDeadTrigger Trigger
	uiToBg           chan string
	bgToUi           chan []Skill
}

func (f *Fight) Turn() {
	for _, v := range *f.p.GetEffects() {
		v.DecreaseCD()
	}
	for _, en := range f.enemies {
		for _, v := range *en.GetEffects() {
			v.DecreaseCD()
		}
	}
	skills := make([]Skill, 0)
	for _, v := range f.p.SelfSkills {
		skills = append(skills, v)
	}
	f.bgToUi <- skills
	res := <-f.uiToBg
	//var chosenSelfSkill PlayerSelfSkill
	if res == "" {
		Inform("Prompt returned empty string, selfskill")
		return
	}
	skillNum, err := strconv.Atoi(res)
	if err != nil {
		Inform("Prompt returned bad value: " + res)
		return
	}
	chosenSelfSkill := f.p.SelfSkills[skillNum-1]

	f.pq.Push(chosenSelfSkill)

	for _, v := range f.enemies {

		dmgSkills := make([]Skill, 0)
		for _, v := range f.p.DmgSkills {
			if v.GetUses() > 0 {
				dmgSkills = append(dmgSkills, v)
			}
		}
		f.bgToUi <- dmgSkills
		dmgSkill := <-f.uiToBg
		if dmgSkill == "" {
			Inform("Prompt returned empty string, dmgskill")
		}
		dmgSkillNum, err := strconv.Atoi(dmgSkill)
		if err != nil {
			Inform("Prompt returned bad value: " + dmgSkill)
			return
		}
		chosenDmgSkill := f.p.DmgSkills[dmgSkillNum-1]
		chosenDmgSkill.SetTarget(v)
		f.pq.Push(chosenDmgSkill)
		ensk := v.ChooseSkill()
		ensk.SetTarget(f.p)
		f.pq.Push(v.ChooseSkill())
	}
	for f.pq.Len() > 0 {
		sk := f.pq.Pop().(Skill)
		// what if the target died? Just miss that use? Redirect to random?
		// if player is dead, then skip 100%. For consistency, let's for now skip all the time
		if sk.GetWielder().IsAlive() && !FindEffect(sk.GetWielder(), Stun) && sk.GetTarget().IsAlive() {
			res := sk.Apply(f)
			f.bgToUi <- []Skill{sk}
			Inform(fmt.Sprintf(
				"%s used %s on %s, %s\n",
				sk.GetWielder().GetName(),
				sk.GetName(),
				sk.GetTarget().GetName(),
				res))
		} else if FindEffect(sk.GetWielder(), Stun) {
			sk.ApplyVoid("stun")
			f.bgToUi <- []Skill{sk}
			Inform(fmt.Sprintf(
				"%s tried to use %s on %s, but was stunned\n",
				sk.GetWielder().GetName(),
				sk.GetName(),
				sk.GetTarget().GetName()))
			RemoveEffect(sk.GetWielder(), Stun)
		}
	}
	f.bgToUi <- nil
	for _, v := range f.p.DmgSkills {
		v.Reset()
	}
	RemoveExpiredEffects(f.p)
	for _, en := range f.enemies {
		RemoveExpiredEffects(en)
	}
	RemoveDeadEnemies(f)
}

func (f *Fight) StartFight(p *Player, enemies []*Enemy, bgToUi chan []Skill, uiToBg chan string) {
	//heap.Init(&f.pq)
	f.p = p
	f.enemies = enemies
	f.defeated = make([]*Enemy, 0)
	f.uiToBg = uiToBg
	f.bgToUi = bgToUi
	for len(f.enemies) > 0 && p.CurMentHP > 0 && p.CurPhysHP > 0 {
		f.Turn()
	}
	Inform(fmt.Sprintf("Your HP: %d", f.p.CurPhysHP))
}
