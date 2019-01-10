package Interfaces

import (
	"math/rand"
)

type CreatureType int

const (
	Human CreatureType = iota
	Undead
	Vampire
	Animal
)

type AiLevel int

const (
	Usual AiLevel = iota
	Miniboss
	Boss
)

type Enemy struct {
	enemyType CreatureType
	skills    []NPCSkill
	equipment []*Equippable
	loot      []Lootable
	provision []Stack
	aiLevel   AiLevel
	BasicUnit
}

func NewEnemy(typ CreatureType, skills []NPCSkill, eqiup []*Equippable,
	loot []Lootable, provision []Stack, level AiLevel,
	name string, hp int, stats map[Stat]int) *Enemy {
	e := &Enemy{}
	e.enemyType = typ
	e.skills = skills
	e.equipment = eqiup
	e.loot = loot
	e.provision = provision
	e.aiLevel = level
	e.name = name
	e.maxHP = hp
	e.curHP = hp
	e.stats = stats
	e.dmgTakenTrigger = NewTrigger()
	e.onDeathTrigger = NewTrigger()
	e.effects = make([]Effect, 0)
	return e

}

func (e *Enemy) ChangeHealth(damage int) int {
	if damage < 0 { // значит, это хил
		e.curHP -= damage
		if e.curHP > e.maxHP {
			e.curHP = e.maxHP
		}
		return -damage
	}
	def := 0
	for _, v := range e.equipment {
		def += v.defence
	}
	if def > 80 { // ограничиваем максимальную броню 80 процентами поглощения урона
		def = 80
	}
	damage -= damage * def / 100
	e.curHP -= damage
	if e.curHP < 0 {
		e.curHP = 0
	}
	return damage
}

func (e *Enemy) ChooseSkill() NPCSkill {
	switch e.aiLevel {
	case Usual:
		return e.skills[rand.Intn(len(e.skills))]
	case Miniboss: //TODO write the minimap or another algorithm for their ai
		return nil
	case Boss:
		return nil
	default:
		return e.skills[rand.Intn(len(e.skills))]
	}
}

func (e *Enemy) GetMoney() int {
	total := 0
	for _, v := range e.loot {
		total += v.GetValue()
	}
	return total
}

func (e *Enemy) GetProvision() []Stack {
	res := make([]Stack, 0)
	for _, v := range e.equipment {
		res = append(res, NewStack(v, 1))
	}
	return append(e.provision, res...)
}

func (e *Enemy) AddSkill(skill NPCSkill) {
	e.skills = append(e.skills, skill)
}

func (e *Enemy) ChooseTarget(r *Room, skillType SkillType) Unit {
	switch skillType {
	case Allies:
		switch e.aiLevel {
		case Miniboss: //TODO write the minimap or another algorithm for their ai
			return nil
		case Boss:
			return nil
		default:
			return r.enemies[rand.Intn(len(r.enemies))]
		}
	case OppositeSide:
		{
			switch e.aiLevel {
			case Miniboss: //TODO write the minimap or another algorithm for their ai
				return nil
			case Boss:
				return nil
			default:
				{
					coin := rand.Intn(2)
					if r.p.pet == nil || coin == 0 {
						return r.p
					}
					return r.p.pet
				}
			}
		}
	case OnlyPlayer:
		return r.p
	case OnlyPet:
		return r.p.pet
	default:
		return e
	}
}
