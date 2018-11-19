package EnemySkills

import . "../Interfaces"

type DogBite struct {
	baseDmg int
	speed   int
	name    string
	wielder Unit
	target  Unit
	res     string
}

func (b *DogBite) GetRes() string {
	return b.res
}

func (b *DogBite) ApplyVoid(res string) {
	b.res = res
}

func (b *DogBite) SetTarget(player Unit) {
	b.target = player
}

func (b *DogBite) GetTarget() Unit {
	return b.target
}

func (b *DogBite) GetWielder() Unit {
	return b.wielder
}

func (b *DogBite) GetSpeed() int {
	return b.speed
}

func (b *DogBite) GetName() string {
	return b.name
}

func (b *DogBite) Init(enemy Unit) Skill {
	b.baseDmg = 5
	b.speed = 6
	b.name = "Bite"
	b.wielder = enemy
	b.res = ""
	return b
}

func (b *DogBite) Apply(r *Room) string {
	if b.wielder.GetHP() > 0 {
		b.res = DealDamage(b.wielder, b.target, b.baseDmg)
	}
	return b.res
}
