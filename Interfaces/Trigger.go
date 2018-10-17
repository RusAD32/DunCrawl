package Interfaces

type Trigger struct {
	Events []Triggerable
}

func TriggerInit() *Trigger {
	t := &Trigger{}
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

func (t *Trigger) Call(values ...interface{}) string {
	res := ""
	for _, v := range t.Events {
		res += "\n" + v.Apply(values...)
		if v.Finished() {
			t.RemoveEvent(v)
		}
	}
	return res
}
