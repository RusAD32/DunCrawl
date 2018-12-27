package Interfaces

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Room struct {
	p                *Player
	enemies          []*Enemy
	defeated         []*Enemy
	pq               PriorityQueue
	turnStartTrigger Trigger
	turnEndTrigger   Trigger
	enemyDeadTrigger Trigger
	uiToBg           chan string
	bgToUi           chan []SkillInfo
	confirm          chan bool
	loot             []Lootable
	shadowLoot       []Lootable
	shadowEnemies    []*Enemy
	shadowProvision  []Stack
	provision        []Stack
	chest            *Chest
	neighbours       []*Wall
	FightState
	dmgSkillsPushed int

	DistFromCenter int
	pathNum        int
	Num            int
	seenInDfs      bool
}

type FightState int

const (
	TurnStart FightState = iota
	AwaitingSelfSkill
	AwaitingDmgSkill
	ResolvingSkills
	FightEnd
)

func (r *Room) AtTurnStart() {
	if r.FightState == TurnStart {
		if len(r.enemies) == 0 || r.p.curMentHP == 0 || r.p.curPhysHP == 0 {
			r.FightState = FightEnd
			return
		}
		for _, v := range *r.p.GetEffects() {
			v.DecreaseCD()
		}
		for _, en := range r.enemies {
			for _, v := range *en.GetEffects() {
				v.DecreaseCD()
			}
		}
		r.FightState = AwaitingSelfSkill
	} else {
		fmt.Println("Error at start!")
	}
}

func (r *Room) SubmitSelfSkill(s PlayerSelfSkill) {
	if r.FightState == AwaitingSelfSkill {
		r.pq.Push(s)
		r.FightState = AwaitingDmgSkill
		r.dmgSkillsPushed = 0
	} else {
		fmt.Println("Error submitting self!")
	}
}

func (r *Room) SubmitDmgSkill(s PlayerDmgSkill) {
	if r.FightState == AwaitingDmgSkill && s.GetUses() >= 0 {
		r.pq.Push(s)
		ensk := r.enemies[r.dmgSkillsPushed].ChooseSkill()
		ensk.SetTarget(r.p)
		r.pq.Push(ensk)
		r.dmgSkillsPushed++
		if r.dmgSkillsPushed == len(r.enemies) {
			r.FightState = ResolvingSkills
		}
	} else {
		fmt.Println("Error at submitting dmg!", s.GetUses(), s.GetName())
	}
}

func (r *Room) GetNextSkillUsed() SkillInfo {
	if r.FightState == ResolvingSkills {
		for {
			sk := r.pq.Pop().(Skill)
			if sk.GetWielder().IsAlive() && !FindEffect(sk.GetWielder(), Stun) {
				sk.Apply(r)
			} else if FindEffect(sk.GetWielder(), Stun) {
				sk.ApplyVoid("stun")
				RemoveEffect(sk.GetWielder(), Stun)
			} else {
				continue
			}
			if len(r.pq) == 0 {
				for _, v := range r.p.dmgSkills {
					v.Reset()
				}
				RemoveExpiredEffects(r.p)
				for _, en := range r.enemies {
					RemoveExpiredEffects(en)
				}
				RemoveDeadEnemies(r)
				r.FightState = TurnStart
				defer r.AtTurnStart()
			}
			return sk
		}
	} else {
		fmt.Println("Error at resolving!")
		return nil
	}
}

func (r *Room) GetPlayer() *Player {
	return r.p
}

func (r *Room) GetEnemies() []*Enemy {
	return r.enemies
}

/**
FightTurn writes skills to you this way:
skills, that you can use on yourself
skills to use on each of the enemies (so you need to keep the same enemies array on the UI side)
skills that were used, in the order of them being used
*/
func (r *Room) FightTurn() {
	for _, v := range *r.p.GetEffects() {
		v.DecreaseCD()
	}
	for _, en := range r.enemies {
		for _, v := range *en.GetEffects() {
			v.DecreaseCD()
		}
	}
	skills := make([]SkillInfo, 0)
	for _, v := range r.p.selfSkills {
		skills = append(skills, v)
	}
	r.bgToUi <- skills
	res := <-r.uiToBg
	//var chosenSelfSkill PlayerSelfSkill
	skillNum, err := strconv.Atoi(res)
	if err != nil {
		fmt.Println("Prompt returned bad value: " + res)
		return
	}
	chosenSelfSkill := r.p.selfSkills[skillNum-1]

	r.pq.Push(chosenSelfSkill)

	for _, v := range r.enemies {

		dmgSkills := make([]SkillInfo, 0)
		for _, v := range r.p.dmgSkills {
			if v.GetUses() > 0 {
				dmgSkills = append(dmgSkills, v)
			}
		}
		r.bgToUi <- dmgSkills
		dmgSkill := <-r.uiToBg
		dmgSkillNum, err := strconv.Atoi(dmgSkill)
		if err != nil {
			fmt.Println("Prompt returned bad value: " + dmgSkill)
			return
		}
		chosenDmgSkill := r.p.dmgSkills[dmgSkillNum-1]
		chosenDmgSkill.SetTarget(v)
		r.pq.Push(chosenDmgSkill)
		ensk := v.ChooseSkill()
		ensk.SetTarget(r.p)
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
	for _, v := range r.p.dmgSkills {
		v.Reset()
	}
	RemoveExpiredEffects(r.p)
	for _, en := range r.enemies {
		RemoveExpiredEffects(en)
	}
	RemoveDeadEnemies(r)
}

func (r *Room) Init(enemies []*Enemy, l *Labyrinth) {
	bgToUi := l.fightBgToUi
	uiToBg := l.fightUiToBg
	confirm := l.fightConfirmChan
	r.enemies = enemies
	r.defeated = make([]*Enemy, 0)
	r.shadowEnemies = make([]*Enemy, 0)
	r.uiToBg = uiToBg
	r.bgToUi = bgToUi
	r.confirm = confirm
	r.neighbours = make([]*Wall, 0)
	r.neighbours = make([]*Wall, 0)
	r.FightState = TurnStart
	for i := 0; i < 4; i++ {
		newWall := Wall{
			Solid,
			nil,
			nil,
		}
		r.neighbours = append(r.neighbours, &newWall)
	}
}

func (r *Room) StartFight() (int, []Stack) {
	for len(r.enemies) > 0 && r.p.curMentHP > 0 && r.p.curPhysHP > 0 {
		r.FightTurn()
	}
	if r.p.IsAlive() {
		totalMoney := 0
		totalProvision := make([]Stack, 0)
		for _, v := range r.defeated {
			totalMoney += v.GetMoney()
			totalProvision = append(totalProvision, v.GetProvision()...)
		}
		r.defeated = make([]*Enemy, 0)
		defer func() {
			r.confirm <- true
		}()
		return totalMoney, totalProvision
	}
	defer func() { r.confirm <- false }()
	//Inform(fmt.Sprintf("Your hp: %d", r.p.curPhysHP))
	return 0, make([]Stack, 0)
}

func (r *Room) GetValues() (int, []Stack) {
	return r.GetMoney(), r.GetLoot()
}

func (r *Room) GetMoney() int {
	total := 0
	for _, v := range r.loot {
		total += v.GetValue()
	}
	r.loot = make([]Lootable, 0)
	return total
}

func (r *Room) GetLoot() []Stack {
	res := r.provision
	r.provision = make([]Stack, 0)
	return res
}

func (r *Room) HasChest() bool {
	return r.chest != nil
}

func (r *Room) HasEnemies() bool {
	return len(r.enemies) > 0
}

func (r *Room) HasShadowEnemies() bool {
	return len(r.shadowEnemies) > 0
}

func (r *Room) UnlockChest() (int, []Stack) {
	if r.chest != nil {
		return r.chest.GetMoney(), r.chest.GetValuables()
	}
	return 0, make([]Stack, 0)
}

func (r *Room) Light() (int, []Stack) {
	totalMoney := 0
	totalProvision := make([]Stack, 0)
	if len(r.shadowEnemies) > 0 {
		r.enemies = r.shadowEnemies
		money, prov := r.StartFight()
		totalMoney += money
		totalProvision = append(totalProvision, prov...)
	}
	for _, v := range r.shadowLoot {
		totalMoney += v.GetValue()
	}
	totalProvision = append(totalProvision, r.shadowProvision...)
	return totalMoney, totalProvision
}

func (r *Room) GetNeighbours() []*Wall {
	return r.neighbours
}

func (r *Room) GetChannels() (chan string, chan []SkillInfo, chan bool) {
	return r.uiToBg, r.bgToUi, r.confirm
}

func (r *Room) AddLoot(lootable Lootable) {
	r.loot = append(r.loot, lootable)
}

func (r *Room) SetChest(chest *Chest) {
	r.chest = chest
}

func (r *Room) AddShadowLoot(lootable Lootable) {
	r.shadowLoot = append(r.shadowLoot, lootable)
}

func (r *Room) GetLocks() int {
	locks := 0
	for _, v := range r.GetNeighbours() {
		if !v.CanGoThrough() {
			locks++
		}
	}
	return locks
}

func BackTrackerLabGen(room *Room, distFromStart int) {
	if room.seenInDfs {
		return
	}
	room.seenInDfs = true
	room.DistFromCenter = distFromStart
	availNeighbourNums := make([]int, 0)
	for i, v := range room.neighbours {
		if v.GetNextDoor() != nil {
			availNeighbourNums = append(availNeighbourNums, i)
		}
	}
	for _, v := range rand.Perm(len(availNeighbourNums)) {
		num := availNeighbourNums[v]
		if !room.neighbours[num].leadsTo.seenInDfs {
			UnockRooms(room, room.neighbours[num].leadsTo, num)
			BackTrackerLabGen(room.neighbours[num].leadsTo, distFromStart+1)
		}
	}
}
