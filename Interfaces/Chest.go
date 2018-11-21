package Interfaces

//TODO Generating chests
type Chest struct {
	trap       Trap
	loot       []Lootable
	usefulLoot []Carriable
}

func (c *Chest) TrapTrigger(p *Player) {
	if c.trap != nil {
		c.trap.Trigger(p)
	}
}

func (c *Chest) TrapDisarm() {
	if c.trap != nil {
		c.trap = nil
	}
}

func (c *Chest) GetMoney() int {
	total := 0
	for _, v := range c.loot {
		total += v.GetValue()
	}
	c.loot = make([]Lootable, 0)
	return total
}

func (c *Chest) GetValuables() []Carriable {
	res := c.usefulLoot
	c.usefulLoot = make([]Carriable, 0)
	return res
}

func GetDefaultChest() *Chest {
	return &Chest{
		nil,
		[]Lootable{GenerateLootable("Different stuff", 40)},
		make([]Carriable, 0),
	}
}
