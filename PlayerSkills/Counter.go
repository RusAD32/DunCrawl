package PlayerSkills

import (
	. "DunCrawl/Interfaces"
	. "DunCrawl/Triggerables"
	"fmt"
)

type Counter struct {
	hp int
	CommonSelfSkill
}

func (c *Counter) Apply(r *Room) string {
	c.wielder.AddDamageTriggerable(NewCounterTriggerable(3, 4, c.wielder))
	c.res = "Counter"
	return c.res
}

func (c *Counter) Init(player Unit) Skill {
	c.lvl = 1
	c.curExp = 0
	c.maxLvl = 3
	c.lvlupExp = []int{1, 4}
	c.speed = 9
	c.name = "Counter"
	c.wielder = player
	return c
}

func (c *Counter) LvlUp() {
	if c.lvl < c.maxLvl && c.curExp >= c.lvlupExp[c.lvl-1] {
		c.curExp -= c.lvlupExp[c.lvl+1]
		c.lvl++
		c.hp = int(c.hp * 3 / 2) // Why won't you multiply by 1.5?
	} else {
		fmt.Printf("Error: Requirements for levelling up skill %s not met", c.name)
	}
}
