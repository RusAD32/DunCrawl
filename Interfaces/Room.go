package Interfaces

import (
	"strconv"
)

type Room struct {
	P                *Player
	Enemies          []*Enemy
	Defeated         []*Enemy
	pq               PriorityQueue
	TurnStartTrigger Trigger
	TurnEndTrigger   Trigger
	EnemyDeadTrigger Trigger
	uiToBg           chan string
	bgToUi           chan []SkillInfo
	confirm          chan bool
	Loot             []Lootable
	ShadowLoot       []Lootable
	ShadowEnemies    []*Enemy
	ShadowProvision  []Carriable
	Provision        []Carriable
	Chest            *Chest
	Neighbours       []*Wall
	Num              int
}

/**
FightTurn writes skills to you this way:
Skills, that you can use on yourself
Skills to use on each of the Enemies (so you need to keep the same Enemies array on the UI side)
Skills that were used, in the order of them being used
*/
func (r *Room) FightTurn() {
	for _, v := range *r.P.GetEffects() {
		v.DecreaseCD()
	}
	for _, en := range r.Enemies {
		for _, v := range *en.GetEffects() {
			v.DecreaseCD()
		}
	}
	skills := make([]SkillInfo, 0)
	for _, v := range r.P.SelfSkills {
		skills = append(skills, v)
	}
	r.bgToUi <- skills
	res := <-r.uiToBg
	//var chosenSelfSkill PlayerSelfSkill
	skillNum, err := strconv.Atoi(res)
	if err != nil {
		Inform("Prompt returned bad value: " + res)
		return
	}
	chosenSelfSkill := r.P.SelfSkills[skillNum-1]

	r.pq.Push(chosenSelfSkill)

	for _, v := range r.Enemies {

		dmgSkills := make([]SkillInfo, 0)
		for _, v := range r.P.DmgSkills {
			if v.GetUses() > 0 {
				dmgSkills = append(dmgSkills, v)
			}
		}
		r.bgToUi <- dmgSkills
		dmgSkill := <-r.uiToBg
		dmgSkillNum, err := strconv.Atoi(dmgSkill)
		if err != nil {
			Inform("Prompt returned bad value: " + dmgSkill)
			return
		}
		chosenDmgSkill := r.P.DmgSkills[dmgSkillNum-1]
		chosenDmgSkill.SetTarget(v)
		r.pq.Push(chosenDmgSkill)
		ensk := v.ChooseSkill()
		ensk.SetTarget(r.P)
		r.pq.Push(ensk)
	}
	skillsUsed := make([]SkillInfo, 0)
	for r.pq.Len() > 0 {
		sk := r.pq.Pop().(Skill)
		// what if the target died? Just miss that use? Redirect to random?
		// if player is dead, then skip 100%. For consistency, let's for now skip all the time
		if sk.GetWielder().IsAlive() && !FindEffect(sk.GetWielder(), Stun) {
			sk.Apply(r)
			skillsUsed = append(skillsUsed, sk)
		} else if FindEffect(sk.GetWielder(), Stun) {
			sk.ApplyVoid("stun")
			skillsUsed = append(skillsUsed, sk)
			RemoveEffect(sk.GetWielder(), Stun)
		}
	}
	r.bgToUi <- skillsUsed
	for _, v := range r.P.DmgSkills {
		v.Reset()
	}
	RemoveExpiredEffects(r.P)
	for _, en := range r.Enemies {
		RemoveExpiredEffects(en)
	}
	RemoveDeadEnemies(r)
}

func (r *Room) Init(p *Player, enemies []*Enemy, bgToUi chan []SkillInfo, uiToBg chan string, confirm chan bool, num int) {
	r.P = p
	r.Enemies = enemies
	r.Defeated = make([]*Enemy, 0)
	r.ShadowEnemies = make([]*Enemy, 0)
	r.uiToBg = uiToBg
	r.bgToUi = bgToUi
	r.confirm = confirm
	r.Neighbours = make([]*Wall, 4)
	r.Num = num
}

func (r *Room) StartFight() (int, []Carriable) {
	for len(r.Enemies) > 0 && r.P.CurMentHP > 0 && r.P.CurPhysHP > 0 {
		r.FightTurn()
	}
	if r.P.IsAlive() {
		totalMoney := 0
		totalProvision := make([]Carriable, 0)
		for _, v := range r.Defeated {
			totalMoney += v.GetMoney()
			totalProvision = append(totalProvision, v.GetProvision()...)
		}
		defer func() { r.confirm <- true }()
		return totalMoney, totalProvision
	}
	defer func() { r.confirm <- false }()
	//Inform(fmt.Sprintf("Your HP: %d", r.P.CurPhysHP))
	return 0, make([]Carriable, 0)
}

func (r *Room) GetValues() (int, []Carriable) {
	return r.GetMoney(), r.GetLoot()
}

func (r *Room) GetMoney() int {
	total := 0
	for _, v := range r.Loot {
		total += v.GetValue()
	}
	return total
}

func (r *Room) GetLoot() []Carriable {
	return r.Provision
}

func (r *Room) HasChest() bool {
	return r.Chest != nil
}

func (r *Room) HasEnemies() bool {
	return len(r.Enemies) > 0
}

func (r *Room) HasShadowEnemies() bool {
	return len(r.ShadowEnemies) > 0
}

func (r *Room) UnlockChest() (int, []Carriable) {
	if r.Chest != nil {
		return r.Chest.GetMoney(), r.Chest.GetValuables()
	}
	return 0, make([]Carriable, 0)
}

func (r *Room) Light() (int, []Carriable) {
	totalMoney := 0
	totalProvision := make([]Carriable, 0)
	if len(r.ShadowEnemies) > 0 {
		r.Enemies = r.ShadowEnemies
		money, prov := r.StartFight()
		totalMoney += money
		totalProvision = append(totalProvision, prov...)
	}
	for _, v := range r.ShadowLoot {
		totalMoney += v.GetValue()
	}
	totalProvision = append(totalProvision, r.ShadowProvision...)
	return totalMoney, totalProvision
}

func (r *Room) GetNeighbours() []*Wall {
	return r.Neighbours
}

func (r *Room) GetChannels() (chan string, chan []SkillInfo, chan bool) {
	return r.uiToBg, r.bgToUi, r.confirm
}
