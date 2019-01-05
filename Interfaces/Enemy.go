package Interfaces

import "math/rand"

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
	equipment []Equippable
	loot      []Lootable
	provision []Carriable
	aiLevel   AiLevel
	BasicUnit
}

func (e *Enemy) GetMaxHP() int {
	return e.maxHP
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
		st := CarriableStack{}
		st.Init(v, 1)
		res = append(res, &st)
	}
	for _, v := range e.provision {
		st := CarriableStack{}
		st.Init(v, 1)
		res = append(res, &st)
	}
	return res
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
	case Enemies:
		{
			switch e.aiLevel {
			case Miniboss: //TODO write the minimap or another algorithm for their ai
				return nil
			case Boss:
				return nil
			default:
				{
					coin := rand.Intn(1)
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
