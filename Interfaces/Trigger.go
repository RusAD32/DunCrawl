package Interfaces

type Triggerable interface {
	Init(values ...interface{}) Triggerable
	Apply(values ...interface{}) string
	Finished() bool
	Dispose()
}

type Trigger struct {
	events []Triggerable
}

func (t *Trigger) Init() *Trigger {
	t.events = make([]Triggerable, 0)
	return t
}

func (t *Trigger) AddEvent(event Triggerable) {
	t.events = append(t.events, event)
}

func (t *Trigger) RemoveEvent(event Triggerable) {
	for i := len(t.events); i >= 0; i-- {
		if t.events[i] == event {
			t.events[i] = nil
			t.events = append(t.events[:i], t.events[i+1:]...)
			return
		}
	}
}

func (t *Trigger) Call(values ...interface{}) string {
	/*
		For now used in counterattack to specify targets
		should probably be changed to ...Unit or something.
		But maybe TurnStart/TurnEnd would be dependant on turn number or a number of enemies left, so I'll leave it for now
	*/
	res := ""
	for _, v := range t.events {
		res += "\n" + v.Apply(values...)
		if v.Finished() {
			t.RemoveEvent(v)
		}
	}
	return res
}
