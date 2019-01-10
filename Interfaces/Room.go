package Interfaces

import (
	"fmt"
	"math/rand"
)

type Room struct {
	p                *Player
	enemies          []*Enemy
	defeated         []*Enemy
	pq               PriorityQueue
	turnStartTrigger Trigger
	turnEndTrigger   Trigger
	loot             []*Lootable
	shadowLoot       []*Lootable
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

func NewRoom(enemies, shEnemies []*Enemy, loot, shLoot []*Lootable, provision, shProvision []Stack, chest *Chest) *Room {
	r := &Room{
		enemies:         enemies,
		defeated:        make([]*Enemy, 0),
		shadowEnemies:   shEnemies,
		neighbours:      make([]*Wall, 0),
		chest:           chest,
		loot:            loot,
		shadowLoot:      shLoot,
		provision:       provision,
		shadowProvision: shProvision,
	}
	if enemies != nil && len(enemies) > 0 {
		r.FightState = TurnStart
	} else {
		r.FightState = FightEnd
	}
	for i := 0; i < 4; i++ {
		newWall := Wall{
			kind: Solid,
		}
		r.neighbours = append(r.neighbours, &newWall)
	}
	return r
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
		r.turnStartTrigger.Call()
		if len(r.enemies) == 0 || !r.p.IsAlive() {
			for _, v := range r.defeated {
				r.loot = append(r.loot, v.loot...)
				r.provision = append(r.provision, v.GetProvision()...)
			}
			r.defeated = make([]*Enemy, 0)
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
		if r.p.pet != nil {
			petsk := r.p.pet.ChooseSkill()
			petsk.SetTarget(r.p.pet.ChooseTarget(r, petsk.GetSkillType()))
			r.pq.Push(petsk)
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
		sk := s.Copy()
		r.pq.Push(sk)
		r.dmgSkillsPushed++
		if r.dmgSkillsPushed == len(r.enemies) {
			for _, v := range r.enemies {
				ensk := v.ChooseSkill()
				ensk.SetTarget(v.ChooseTarget(r, ensk.GetSkillType()))
				r.pq.Push(ensk)
			}
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
				RemoveSkillsOfDeadUnits(&r.pq)
			} else if sk.GetWielder().IsAlive() {
				sk.ApplyVoid("stun")
				RemoveEffect(sk.GetWielder(), Stun)
			} else {
				continue
			}
			RemoveDeadEnemies(r)
			if len(r.pq) == 0 {
				r.turnEndTrigger.Call()
				for _, v := range r.p.dmgSkills {
					v.Reset()
				}
				RemoveExpiredEffects(r.p)
				for _, en := range r.enemies {
					RemoveExpiredEffects(en)
				}
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

func (r *Room) GetValues() ([]*Lootable, []Stack) {
	return r.GetLoot(), r.GetGoodies()
}

func (r *Room) GetLoot() []*Lootable {
	//TODO решить, надо ли их забирать, и привести к единому стандарту
	return r.loot
}

func (r *Room) GetGoodies() []Stack {
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

func (r *Room) UnlockChest() ([]*Lootable, []Stack) {
	if r.chest != nil {
		return r.chest.GetLoot(), r.chest.GetValuables()
	}
	return make([]*Lootable, 0), make([]Stack, 0)
}

func (r *Room) Light() {
	r.enemies = append(r.enemies, r.shadowEnemies...)
	r.provision = append(r.provision, r.shadowProvision...)
	r.loot = append(r.loot, r.shadowLoot...)
	if len(r.enemies) > 0 {
		r.FightState = TurnStart
		r.AtTurnStart()
	}
}

func (r *Room) GetNeighbours() []*Wall {
	return r.neighbours
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

func (r *Room) GetSkQueue() []SkillInfo {
	res := make([]SkillInfo, 0)
	for _, v := range r.pq {
		res = append(res, v.(SkillInfo))
	}
	return res
}
