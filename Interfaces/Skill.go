package Interfaces

type SkillType int

type Skill interface {
	GetName() string
	GetTarget() Unit
	GetWielder() Unit
	Init(wielder Unit)
	Apply(f *Fight) string
	ApplyVoid(res string)
	GetRes() string
}

type HasSpeed interface {
	GetSpeed() int
}

type PlayerDmgSkill interface {
	HasSpeed
	Skill
	LvlUp()
	AddExp(amount int)
	GetUses() int
	SetTarget(target Unit)
	Reset()
}

type PlayerSelfSkill interface {
	HasSpeed
	Skill
	LvlUp()
	AddExp(amount int)
}

type EnemySkill interface {
	HasSpeed
	Skill
	SetTarget(player Unit)
}
