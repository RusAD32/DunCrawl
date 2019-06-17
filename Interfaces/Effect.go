package Interfaces

type EffectID int

const (
	Stun EffectID = iota
	Confusion
	CounterAtk
	Winded1
	Winded2
	Winded3
)

type Effect interface {
	GetID() EffectID
	GetAmount() int // this is for effects that have extra description. Like stat modifier
	GetInfo() string
	DecreaseCD()
	GetCD() int
}
