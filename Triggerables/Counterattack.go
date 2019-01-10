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

func NewCounterTriggerable(atk, amnt int, user Interfaces.Unit) Interfaces.Triggerable {
	c := &Counterattack{}
	c.baseAtk = atk
	c.triggerNum = amnt
	c.user = user
	cntr := Effects.NewCounterEff(amnt) //TODO этот эффект будет уменьшаться в начале хода. Переделать на отображение самого триггера в ui?
	Interfaces.AddEffect(c.user, cntr)
	c.eff = cntr
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
