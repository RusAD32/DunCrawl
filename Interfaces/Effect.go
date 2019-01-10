package Interfaces

type EffectID int

const (
	Stun EffectID = iota
	Confusion
	CounterAtk
)

type Effect interface {
	GetID() EffectID
	GetAmount() int // this is for effects that have extra description. Like stat modifier
	GetInfo() string
	DecreaseCD()
	GetCD() int
}
