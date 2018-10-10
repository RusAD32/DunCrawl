package Interfaces

import "math/rand"

type EnemyType int

const (
	Human EnemyType = iota
	Undead
	Vampire
	Animal
)

type AILevel int

const (
	Usual AILevel = iota
	Miniboss
	Boss
)

type Enemy struct {
	Type      EnemyType
	Name      string
	Skills    []EnemySkill
	Stats     map[Stat]int
	Effects   []Effect
	Equipment []Equippable
	AILevel
	CurHP           int
	MaxHP           int
	DmgTakenTrigger *Trigger
}

func (e *Enemy) IsAlive() bool {
	return e.CurHP > 0
}

func (e *Enemy) GetDamageTrigger() *Trigger {
	return e.DmgTakenTrigger
}

func (e *Enemy) AddDamageTriggerable(t Triggerable) {
	e.DmgTakenTrigger.AddEvent(t)
}

func (e *Enemy) GetHP() int {
	return e.CurHP
}

func (e *Enemy) GetName() string {
	return e.Name
}

func (e *Enemy) AddEffect(effect Effect) {
	e.Effects = append(e.Effects, effect)
}

func (e *Enemy) ChangeHealth(damage int) int {
	if damage < 0 { // значит, это хил
		e.CurHP -= damage
		if e.CurHP > e.MaxHP {
			e.CurHP = e.MaxHP
		}
		return -damage
	}
	def := 0
	for _, v := range e.Equipment {
		def += v.Defence
	}
	if def > 80 { // ограничиваем максимальную броню 80 процентами поглощения урона
		def = 80
	}
	damage -= damage * def / 100
	e.CurHP -= damage
	if e.CurHP < 0 {
		e.CurHP = 0
	}
	return damage
}

func (e *Enemy) GetEffects() *[]Effect {
	return &e.Effects
}

func (e *Enemy) ChooseSkill() EnemySkill {
	switch e.AILevel {
	case Usual:
		return e.Skills[rand.Intn(len(e.Skills))]
	case Miniboss: //TODO write the minimap or another algorithm for their ai
	case Boss:
	default:
		return e.Skills[rand.Intn(len(e.Skills))]
	}
	return nil
}
