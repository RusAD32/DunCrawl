package Interfaces

type Triggerable interface {
	Init(values ...interface{}) Triggerable
	Apply(values ...interface{}) string
}
