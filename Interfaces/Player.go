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
	Stats      map[Stat]int
	Equipment  map[Slot]Equippable
	Inventory  []int
	DmgSkills  []PlayerDmgSkill
	SelfSkills []PlayerSelfSkill
	Effects    []Effect
	Lvl        int
	Exp        int
	CurPhysHP  int
	MaxPhysHP  int
	CurMentHP  int
	MaxMentHP  int
}

func (p *Player) TakeDamage(damage int) {
	def := 0
	for _, v := range p.Equipment {
		def += v.Defence
	}
	if def > 80 {
		def = 80
	}
	damage -= damage * def / 100
	p.CurPhysHP -= damage
}

func (p *Player) GetEffects() []Effect {
	return p.Effects
}
