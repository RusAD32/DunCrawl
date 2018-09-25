package Interfaces

type EffectID int

const (
	Stun EffectID = iota
	Confusion
)

type Effect interface {
	Init(values ...interface{})
	GetID() EffectID
	GetAmount() int
	GetInfo() string
	DecreaseCDs()
	GetCDs()
}
