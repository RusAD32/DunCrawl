package Interfaces

type Triggerable interface {
	Init(values ...interface{})
	Apply()
	GetCondition() string
}
