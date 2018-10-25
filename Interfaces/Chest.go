package Interfaces

type Chest struct {
	Trap       Trap
	Loot       []Lootable
	UsefulLoot []Carriable
}

func (c *Chest) TrapTrigger(p *Player) {
	if c.Trap != nil {
		c.Trap.Trigger(p)
	}
}

func (c *Chest) TrapDisarm() {
	if c.Trap != nil {
		c.Trap = nil
	}
}

func (c *Chest) GetMoney() int {
	total := 0
	for _, v := range c.Loot {
		total += v.GetValue()
	}
	c.Loot = make([]Lootable, 0)
	return total
}

func (c *Chest) GetValuables() []Carriable {
	res := c.UsefulLoot
	c.UsefulLoot = make([]Carriable, 0)
	return res
}
