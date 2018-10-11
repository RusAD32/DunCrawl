package Interfaces

import (
	"strconv"
)

type Fight struct {
	P                *Player
	Enemies          []*Enemy
	Defeated         []*Enemy
	pq               PriorityQueue
	TurnStartTrigger Trigger
	TurnEndTrigger   Trigger
	EnemyDeadTrigger Trigger
	uiToBg           chan string
	bgToUi           chan []SkillInfo
}

/**
Turn writes skills to you this way:
Skills, that you can use on yourself
Skills to use on each of the Enemies (so you need to keep the same Enemies array on the UI side)
Skills that were used, in the order of them being used
*/
func (f *Fight) Turn() {
	for _, v := range *f.P.GetEffects() {
		v.DecreaseCD()
	}
	for _, en := range f.Enemies {
		for _, v := range *en.GetEffects() {
			v.DecreaseCD()
		}
	}
	skills := make([]SkillInfo, 0)
	for _, v := range f.P.SelfSkills {
		skills = append(skills, v)
	}
	f.bgToUi <- skills
	res := <-f.uiToBg
	//var chosenSelfSkill PlayerSelfSkill
	skillNum, err := strconv.Atoi(res)
	if err != nil {
		Inform("Prompt returned bad value: " + res)
		return
	}
	chosenSelfSkill := f.P.SelfSkills[skillNum-1]

	f.pq.Push(chosenSelfSkill)

	for _, v := range f.Enemies {

		dmgSkills := make([]SkillInfo, 0)
		for _, v := range f.P.DmgSkills {
			if v.GetUses() > 0 {
				dmgSkills = append(dmgSkills, v)
			}
		}
		f.bgToUi <- dmgSkills
		dmgSkill := <-f.uiToBg
		dmgSkillNum, err := strconv.Atoi(dmgSkill)
		if err != nil {
			Inform("Prompt returned bad value: " + dmgSkill)
			return
		}
		chosenDmgSkill := f.P.DmgSkills[dmgSkillNum-1]
		chosenDmgSkill.SetTarget(v)
		f.pq.Push(chosenDmgSkill)
		ensk := v.ChooseSkill()
		ensk.SetTarget(f.P)
		f.pq.Push(ensk)
	}
	skillsUsed := make([]SkillInfo, 0)
	for f.pq.Len() > 0 {
		sk := f.pq.Pop().(Skill)
		// what if the target died? Just miss that use? Redirect to random?
		// if player is dead, then skip 100%. For consistency, let's for now skip all the time
		if sk.GetWielder().IsAlive() && !FindEffect(sk.GetWielder(), Stun) {
			sk.Apply(f)
			skillsUsed = append(skillsUsed, sk)
		} else if FindEffect(sk.GetWielder(), Stun) {
			sk.ApplyVoid("stun")
			skillsUsed = append(skillsUsed, sk)
			RemoveEffect(sk.GetWielder(), Stun)
		}
	}
	f.bgToUi <- skillsUsed
	for _, v := range f.P.DmgSkills {
		v.Reset()
	}
	RemoveExpiredEffects(f.P)
	for _, en := range f.Enemies {
		RemoveExpiredEffects(en)
	}
	RemoveDeadEnemies(f)
}

func (f *Fight) Init(p *Player, enemies []*Enemy, bgToUi chan []SkillInfo, uiToBg chan string) {
	f.P = p
	f.Enemies = enemies
	f.Defeated = make([]*Enemy, 0)
	f.uiToBg = uiToBg
	f.bgToUi = bgToUi
}

func (f *Fight) StartFight() {
	for len(f.Enemies) > 0 && f.P.CurMentHP > 0 && f.P.CurPhysHP > 0 {
		f.Turn()
	}
	close(f.uiToBg)
	close(f.bgToUi)
	//Inform(fmt.Sprintf("Your HP: %d", f.P.CurPhysHP))
}
