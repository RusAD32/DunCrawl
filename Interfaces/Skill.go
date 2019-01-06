package Interfaces

type SkillType int

const (
	Self SkillType = iota
	OppositeSide
	Allies
	OnlyPlayer
	OnlyPet
)

type SkillInfo interface {
	GetName() string
	GetTarget() Unit
	GetWielder() Unit
	GetRes() string
}

type Skill interface {
	SkillInfo
	Apply(r *Room) string
	ApplyVoid(res string)
	GetSkillType() SkillType
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
	Copy() PlayerDmgSkill
}

type PlayerSelfSkill interface {
	// No, playerDmgSkill can't be a playerSelfSkill.
	HasSpeed
	Skill
	LvlUp()
	AddExp(amount int)
}

type NPCSkill interface {
	HasSpeed
	Skill
	SetTarget(player Unit)
}
