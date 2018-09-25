package Interfaces

type EnemyType int

const (
	Human EnemyType = iota
	Undead
	Vampire
	Spider
)

type Enemy struct {
	Type      EnemyType
	Skills    []EnemySkill
	Stats     map[Stat]int
	Effects   []Effect
	Equipment []Equippable
	CurHP     int
	MaxHP     int
}

func (e *Enemy) TakeDamage(damage int) {
	def := 0
	for _, v := range e.Equipment {
		def += v.Defence
	}
	if def > 80 {
		def = 80
	}
	damage -= damage * def / 100
	e.CurHP -= damage
}

func (e *Enemy) GetEffects() []Effect {
	return e.Effects
}
