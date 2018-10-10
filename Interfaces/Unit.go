package Interfaces

type Unit interface {
	ChangeHealth(damgage int) int
	GetEffects() *[]Effect
	AddEffect(effect Effect)
	GetName() string
	GetHP() int
	IsAlive() bool
	GetDamageTrigger() *Trigger
	AddDamageTriggerable(t Triggerable)
}
