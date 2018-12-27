package Triggerables

import (
	"DunCrawl/Effects"
	"DunCrawl/Interfaces"
	"fmt"
)

type Counterattack struct {
	triggerNum int
	baseAtk    int
	user       Interfaces.Unit
	eff        Interfaces.Effect
}

func (c *Counterattack) Dispose() {
	Interfaces.RemoveEffect(c.user, Interfaces.CounterAtk)
}

func (c *Counterattack) Init(values ...interface{}) Interfaces.Triggerable {
	if len(values) < 3 {
		panic("Counterattack should get its attack, number of triggers and user as Init argument")
	}
	c.baseAtk = values[0].(int)
	c.triggerNum = values[1].(int)
	c.user = values[2].(Interfaces.Unit)
	cntr := Effects.CounterEff{}
	cntr.Init(4)
	Interfaces.AddEffect(c.user, &cntr)
	c.eff = &cntr
	return c
}

func (c *Counterattack) Apply(values ...interface{}) string {
	if !c.Finished() {
		c.eff.DecreaseCD()
		if len(values) < 1 {
			panic("First argument of Apply should be its target")
		}
		target := values[0].(Interfaces.Unit)
		c.triggerNum--
		res := fmt.Sprintf("(%s counterdamage)", Interfaces.DealRawDamage(target, c.baseAtk))
		return res
	}
	return ""
}

func (c *Counterattack) Finished() bool {
	return c.triggerNum <= 0
}
