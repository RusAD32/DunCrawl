package Interfaces

type Unit interface {
	ChangeHealth(damgage int) int
	GetEffects() []Effect
	AddEffect(effect Effect)
	GetName() string
}
