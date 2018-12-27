package PlayerSkills

import (
	. "DunCrawl/Interfaces"
	. "DunCrawl/Triggerables"
	"fmt"
)

type Counter struct {
	hp       int
	lvl      int
	maxLvl   int
	curExp   int
	lvlupExp []int
	speed    int
	name     string
	wielder  Unit
	res      string
}

func (c *Counter) GetRes() string {
	return c.res
}

func (c *Counter) ApplyVoid(res string) {
	c.res = res
}

func (c *Counter) GetTarget() Unit {
	return c.wielder
}

func (c *Counter) Apply(r *Room) string {
	cntr := Counterattack{}
	c.wielder.AddDamageTriggerable(cntr.Init(3, 4, c.wielder))
	c.res = "Counter"
	return c.res
}

func (c *Counter) GetWielder() Unit {
	return c.wielder
}

func (c *Counter) GetSpeed() int {
	return c.speed
}

func (c *Counter) GetName() string {
	return c.name
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
		fmt.Sprintln("Error: Requirements for levelling up skill %s not met", c.name)
	}
}

func (c *Counter) AddExp(amount int) {
	if c.lvl < c.maxLvl {
		c.curExp += amount
		if c.curExp > c.lvlupExp[c.lvl-1] {
			c.LvlUp()
		}
	}
}
