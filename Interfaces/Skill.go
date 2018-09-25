package Interfaces

type HasSpeed interface {
	GetSpeed() int
}

type PlayerDmgSkill interface {
	HasSpeed
	Init()
	LvlUp()
	AddExp(amount int)
	Apply(wielder *Player, opp *Enemy)
	GetUses() int
	GetName() string
}

type PlayerSelfSkill interface {
	HasSpeed
	Init()
	LvlUp()
	AddExp(amount int)
	Apply(wielder *Player)
	GetName() string
}

type EnemySkill interface {
	HasSpeed
	Apply(wielder *Enemy, opp *Player)
	Init()
	GetName() string
}
