package Effects

import . "DunCrawl/Interfaces"

func NewCounterEff(cd int) Effect {
	return newBasicEffect(CounterAtk, cd, "Counter", 0)
}
