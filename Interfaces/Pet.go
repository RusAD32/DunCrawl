package Interfaces

import "math/rand"

type Pet struct {
	creatureType CreatureType
	skills       []NPCSkill
	aiLevel      AiLevel
	BasicUnit
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
	case Enemies:
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
	default:
		return p
	}
}
