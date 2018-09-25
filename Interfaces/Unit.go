package Interfaces

type Unit interface {
	TakeDamage(int)
	GetEffects() []Effect
}
