package Interfaces

type Lootable struct {
	name  string
	value int
}

func (l *Lootable) GetName() string {
	return l.name
}

func (l *Lootable) GetValue() int {
	return l.value
}

func NewLootable(name string, value int) Lootable {
	return Lootable{
		name,
		value,
	}
}
