package Interfaces

type Lootable struct {
    Name string
    Value int
}

func (l *Lootable) GetName() string {
    return l.Name
}

func (l *Lootable) GetValue() int {
    return l.Value
}
