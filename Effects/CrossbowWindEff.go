package Effects

import (
	. "DunCrawl/Interfaces"
)

func NewCrossbowWind1Effect() Effect {
	return newBasicEffect(Winded1, 1, "Started wounding the crossbow", 0)
}

func NewCrossbowWind2Effect() Effect {
	return newBasicEffect(Winded2, 1, "Midway through wounding the crossbow", 0)
}

func NewCrossbowWind3Effect() Effect {
	return newBasicEffect(Winded3, 1, "Finished wounding the crossbow", 0)
}
