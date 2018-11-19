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
	stats           map[Stat]int
	equipment       map[Slot]Equippable
	inventory       []Carriable
	dmgSkills       []PlayerDmgSkill
	selfSkills      []PlayerSelfSkill
	effects         []Effect
	lvl             int
	exp             int
	curPhysHP       int
	maxPhysHP       int
	curMentHP       int
	maxMentHP       int
	dmgTakenTrigger *Trigger
}

func (p *Player) IsAlive() bool {
	return p.curPhysHP > 0 && p.curMentHP > 0
}

func (p *Player) GetDamageTrigger() *Trigger {
	return p.dmgTakenTrigger
}

func (p *Player) AddDamageTriggerable(t Triggerable) {
	p.dmgTakenTrigger.AddEvent(t)
}

func (p *Player) GetHP() int {
	return p.curPhysHP
}

func (p *Player) GetName() string {
	return "you"
}

func (p *Player) AddEffect(effect Effect) {
	p.effects = append(p.effects, effect)
}

func (p *Player) ChangeHealth(damage int) int {
	if damage < 0 { // значит, это хил
		p.curPhysHP -= damage
		if p.curPhysHP > p.maxPhysHP {
			p.curPhysHP = p.maxPhysHP
		}
		return -damage
	}
	def := 0
	for _, v := range p.equipment {
		def += v.defence
	}
	if def > 80 {
		def = 80
	}
	damage -= damage * def / 100
	p.curPhysHP -= damage

	if p.curPhysHP < 0 {
		p.curPhysHP = 0
	}
	return damage
}

func (p *Player) GetEffects() *[]Effect {
	return &p.effects
}

func (p *Player) GetCurHP() int {
	return p.curPhysHP
}

func (p *Player) GetMaxHP() int {
	return p.maxPhysHP
}

func (p *Player) GetEquipment() map[Slot]Equippable {
	return p.equipment
}

func (p *Player) Equip(e Equippable, slot Slot) {
	p.equipment[slot] = e
}

func (p *Player) AddSelfSkill(skill PlayerSelfSkill) {
	p.selfSkills = append(p.selfSkills, skill)
}

func (p *Player) AddDmgSkill(skill PlayerDmgSkill) {
	p.dmgSkills = append(p.dmgSkills, skill)
}
