package Triggerables

import (
	"../Interfaces"
	"fmt"
)

type Counterattack struct {
	triggerNum int
	baseAtk    int
	user       Interfaces.Unit
}

func (c *Counterattack) Init(values ...interface{}) Interfaces.Triggerable {
	if len(values) < 3 {
		panic("Counterattack should get its attack, number of triggers and user as Init argument")
	}
	c.baseAtk = values[0].(int)
	c.triggerNum = values[1].(int)
	c.user = values[2].(Interfaces.Unit)
	return c
}

func (c *Counterattack) Apply(values ...interface{}) string {
	if len(values) < 1 {
		panic("First argument of Apply should be its target")
	}
	target := values[0].(Interfaces.Unit)
	return fmt.Sprintf("%s dealt %s damage back", c.user.GetName(), Interfaces.DealRawDamage(target, c.baseAtk))
}
