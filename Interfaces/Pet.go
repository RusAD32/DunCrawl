package Interfaces

import "math/rand"

type Pet struct {
	creatureType CreatureType
	skills       []NPCSkill
	aiLevel      AiLevel
	BasicUnit
}

func (e *Pet) Init(typ CreatureType, skills []NPCSkill, level AiLevel,
	name string, hp int, stats map[Stat]int) *Pet {
	e.creatureType = typ
	e.skills = skills
	e.aiLevel = level
	e.name = name
	e.maxHP = hp
	e.curHP = hp
	e.stats = stats
	e.dmgTakenTrigger = new(Trigger).Init()
	e.onDeathTrigger = new(Trigger).Init()
	e.effects = make([]Effect, 0)
	return e
}

func (p *Pet) ChooseSkill() NPCSkill {
	switch p.aiLevel {
	case Usual:
		return p.skills[rand.Intn(len(p.skills))]
	case Miniboss: //TODO write the minimap or another algorithm for their ai
	case Boss:
	default:
		return p.skills[rand.Intn(len(p.skills))]
	}
	return nil
}

func (p *Pet) AddSkill(skill NPCSkill) {
	p.skills = append(p.skills, skill)
}

func (p *Pet) ChooseTarget(r *Room, skillType SkillType) Unit {
	switch skillType {
	case OppositeSide:
		switch p.aiLevel {
		case Miniboss: //TODO write the minimap or another algorithm for their ai
			return nil
		case Boss:
			return nil
		default:
			return r.enemies[rand.Intn(len(r.enemies))]
		}
	case Allies:
		{
			switch p.aiLevel {
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
	default:
		return p
	}
}
