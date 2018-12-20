package Interfaces

import "math/rand"

type EnemyType int

const (
	Human EnemyType = iota
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
	enemyType       EnemyType
	name            string
	skills          []EnemySkill
	stats           map[Stat]int
	effects         []Effect
	equipment       []Equippable
	loot            []Lootable
	provision       []Carriable
	aiLevel         AiLevel
	curHP           int
	maxHP           int
	dmgTakenTrigger *Trigger
}

func (e *Enemy) GetCurHP() int {
	return e.curHP
}

func (e *Enemy) GetMaxHP() int {
	return e.maxHP
}

func (e *Enemy) IsAlive() bool {
	return e.curHP > 0
}

func (e *Enemy) GetDamageTrigger() *Trigger {
	return e.dmgTakenTrigger
}

func (e *Enemy) AddDamageTriggerable(t Triggerable) {
	e.dmgTakenTrigger.AddEvent(t)
}

func (e *Enemy) GetHP() int {
	return e.curHP
}

func (e *Enemy) GetName() string {
	return e.name
}

func (e *Enemy) AddEffect(effect Effect) {
	e.effects = append(e.effects, effect)
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

func (e *Enemy) GetEffects() *[]Effect {
	return &e.effects
}

func (e *Enemy) ChooseSkill() EnemySkill {
	switch e.aiLevel {
	case Usual:
		return e.skills[rand.Intn(len(e.skills))]
	case Miniboss: //TODO write the minimap or another algorithm for their ai
	case Boss:
	default:
		return e.skills[rand.Intn(len(e.skills))]
	}
	return nil
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

func (e *Enemy) AddSkill(skill EnemySkill) {
	e.skills = append(e.skills, skill)
}
