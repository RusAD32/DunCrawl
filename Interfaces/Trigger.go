package Interfaces

type Trigger struct {
	events []Triggerable
}

func TriggerInit() *Trigger {
	t := &Trigger{}
	t.events = make([]Triggerable, 0)
	return t
}

func (t *Trigger) AddEvent(event Triggerable) {
	t.events = append(t.events, event)
}

func (t *Trigger) RemoveEvent(event Triggerable) {
	for i, v := range t.events {
		if v == event {

			t.events[i] = nil
			t.events = append(t.events[:i], t.events[i+1:]...)
			return
		}
	}
}

func (t *Trigger) Call(values ...interface{}) string {
	res := ""
	for _, v := range t.events {
		res += "\n" + v.Apply(values...)
		if v.Finished() {
			t.RemoveEvent(v)
		}
	}
	return res
}
