package PlayerSkills

import (
	. "../Interfaces"
	. "../Triggerables"
	"fmt"
)

type Counter struct {
	HP       int
	Lvl      int
	MaxLvl   int
	CurExp   int
	LvlupExp []int
	Speed    int
	Name     string
	Wielder  Unit
	Res      string
}

func (c *Counter) GetRes() string {
	return c.Res
}

func (c *Counter) ApplyVoid(res string) {
	c.Res = res
}

func (c *Counter) GetTarget() Unit {
	return c.Wielder
}

func (c *Counter) Apply(f *Fight) string {
	cntr := Counterattack{}
	c.Wielder.AddDamageTriggerable(cntr.Init(3, 4, c.Wielder))
	c.Res = "Counter"
	return c.Res
}

func (c *Counter) GetWielder() Unit {
	return c.Wielder
}

func (c *Counter) GetSpeed() int {
	return c.Speed
}

func (c *Counter) GetName() string {
	return c.Name
}

func (c *Counter) Init(player Unit) {
	c.Lvl = 1
	c.CurExp = 0
	c.MaxLvl = 3
	c.LvlupExp = []int{1, 4}
	c.Speed = 9
	c.Name = "Counter"
	c.Wielder = player
}

func (c *Counter) LvlUp() {
	if c.Lvl < c.MaxLvl && c.CurExp >= c.LvlupExp[c.Lvl-1] {
		c.CurExp -= c.LvlupExp[c.Lvl+1]
		c.Lvl++
		c.HP = int(c.HP * 3 / 2) // Why won't you multiply by 1.5?
	} else {
		Inform(fmt.Sprintf("Error: Requirements for levelling up skill %s not met", c.Name))
	}
}

func (c *Counter) AddExp(amount int) {
	if c.Lvl < c.MaxLvl {
		c.CurExp += amount
		if c.CurExp > c.LvlupExp[c.Lvl-1] {
			c.LvlUp()
		}
	}
}
