package Interfaces

//TODO Generating chests
type Chest struct {
	trap       Trap
	loot       []*Lootable
	usefulLoot []Stack
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

func (c *Chest) GetLoot() []*Lootable {
	return c.loot
}

func (c *Chest) GetValuables() []Stack {
	res := c.usefulLoot
	c.usefulLoot = make([]Stack, 0)
	return res
}

func NewChest() *Chest {
	return &Chest{
		nil,
		[]*Lootable{NewLootable("Different stuff", 40)},
		make([]Stack, 0),
	}
}
