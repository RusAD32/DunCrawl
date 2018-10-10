package EnemySkills

import . "../Interfaces"

type DogBite struct {
	BaseDmg int
	Speed   int
	Name    string
	Wielder Unit
	Target  Unit
	Res     string
}

func (b *DogBite) GetRes() string {
	return b.Res
}

func (b *DogBite) ApplyVoid(res string) {
	b.Res = res
}

func (b *DogBite) SetTarget(player Unit) {
	b.Target = player
}

func (b *DogBite) GetTarget() Unit {
	return b.Target
}

func (b *DogBite) GetWielder() Unit {
	return b.Wielder
}

func (b *DogBite) GetSpeed() int {
	return b.Speed
}

func (b *DogBite) GetName() string {
	return b.Name
}

func (b *DogBite) Init(enemy Unit) Skill {
	b.BaseDmg = 5
	b.Speed = 6
	b.Name = "Bite"
	b.Wielder = enemy
	return b
}

func (b *DogBite) Apply(f *Fight) string {
	if b.Wielder.GetHP() > 0 {
		return DealDamage(b.Wielder, b.Target, b.BaseDmg)
	}
	return ""
}
