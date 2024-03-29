package Interfaces

type Unit interface {
	ChangeHealth(damgage int) int
	GetEffects() *[]Effect
	AddEffect(effect Effect)
	GetName() string
	GetHP() int
	GetMaxHP() int
	IsAlive() bool
	GetDamageTrigger() *Trigger
	GetDeathTrigger() *Trigger
	AddDamageTriggerable(t Triggerable)
}

type BasicUnit struct {
	name            string
	stats           map[Stat]int
	effects         []Effect
	curHP           int
	maxHP           int
	dmgTakenTrigger *Trigger
	onDeathTrigger  *Trigger
}

func (bu *BasicUnit) IsAlive() bool {
	return bu.curHP > 0
}

func (bu *BasicUnit) GetDamageTrigger() *Trigger {
	return bu.dmgTakenTrigger
}

func (bu *BasicUnit) GetDeathTrigger() *Trigger {
	return bu.onDeathTrigger
}

func (bu *BasicUnit) AddDamageTriggerable(t Triggerable) {
	bu.dmgTakenTrigger.AddEvent(t)
}

func (bu *BasicUnit) GetHP() int {
	return bu.curHP
}

func (bu *BasicUnit) GetMaxHP() int {
	return bu.maxHP
}

func (bu *BasicUnit) GetName() string {
	return bu.name
}

func (bu *BasicUnit) AddEffect(effect Effect) {
	bu.effects = append(bu.effects, effect)
}

func (bu *BasicUnit) ChangeHealth(damage int) int {
	if damage < 0 { // значит, это хил
		bu.curHP -= damage
		if bu.curHP > bu.maxHP {
			bu.curHP = bu.maxHP
		}
		return -damage
	}
	def := 0
	damage -= damage * def / 100
	bu.curHP -= damage
	if bu.curHP < 0 {
		bu.curHP = 0
	}
	return damage
}

func (bu *BasicUnit) GetEffects() *[]Effect {
	return &bu.effects
}
