package Interfaces

type Trigger struct {
	Events    []Triggerable
	Condition string
}

func TriggerInit(cond string) *Trigger {
	t := &Trigger{}
	t.Condition = cond
	t.Events = make([]Triggerable, 0)
	return t
}

func (t *Trigger) AddEvent(event Triggerable) {
	t.Events = append(t.Events, event)
}

func (t *Trigger) RemoveEvent(event Triggerable) {
	for i, v := range t.Events {
		if v == event {
			t.Events[i] = nil
			t.Events = append(t.Events[:i], t.Events[i+1:]...)
			return
		}
	}
}

func (t *Trigger) Call() {
	for _, v := range t.Events {
		v.Apply()
	}
}
