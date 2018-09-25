package EnemySkills

import . "../Interfaces"

type DogBite struct {
	BaseDmg int
	Speed   int
	Name    string
}

func (b *DogBite) GetSpeed() int {
	return b.Speed
}

func (b *DogBite) GetName() string {
	return b.Name
}

func (b *DogBite) Init() {
	b.BaseDmg = 5
	b.Speed = 6
	b.Name = "Bite"
}

func (b *DogBite) Apply(wielder *Enemy, opp *Player) {
	DealDamage(wielder, opp, b.BaseDmg)
}
