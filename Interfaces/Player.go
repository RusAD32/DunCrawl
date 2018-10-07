package Interfaces

type Stat int

const (
	Strength Stat = iota
	Agility
	Endurance
	Constitution
	Intelligence
	Wisdom
	Luck
)

type Player struct {
	Stats           map[Stat]int
	Equipment       map[Slot]Equippable
	Inventory       []Carriable
	DmgSkills       []PlayerDmgSkill
	SelfSkills      []PlayerSelfSkill
	Effects         []Effect
	Lvl             int
	Exp             int
	CurPhysHP       int
	MaxPhysHP       int
	CurMentHP       int
	MaxMentHP       int
	DmgTakenTrigger *Trigger
}

func (p *Player) GetDamageTrigger() *Trigger {
	return p.DmgTakenTrigger
}

func (p *Player) AddDamageTriggerable(t Triggerable) {
	p.DmgTakenTrigger.AddEvent(t)
}

func (p *Player) GetHP() int {
	return p.CurPhysHP
}

func (p *Player) GetName() string {
	return "you"
}

func (p *Player) AddEffect(effect Effect) {
	p.Effects = append(p.Effects, effect)
}

func (p *Player) ChangeHealth(damage int) int {
	if damage < 0 { // значит, это хил
		p.CurPhysHP -= damage
		if p.CurPhysHP > p.MaxPhysHP {
			p.CurPhysHP = p.MaxPhysHP
		}
		return damage
	}
	def := 0
	for _, v := range p.Equipment {
		def += v.Defence
	}
	if def > 80 {
		def = 80
	}
	damage -= damage * def / 100
	p.CurPhysHP -= damage

	if p.CurPhysHP < 0 {
		p.CurPhysHP = 0
	}
	return damage
}

func (p *Player) GetEffects() *[]Effect {
	return &p.Effects
}
