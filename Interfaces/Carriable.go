package Interfaces

type Carriable interface {
	Use(player *Player, values ...interface{})
	GetName() string
	StacksBy() int
}

type Stack interface {
	Use(player *Player, values ...interface{})
	Init(item Carriable, amount int)
	Add(amount int) int
	Remove(amount int) int
	GetName() string
}

type CarriableStack struct {
	item     Carriable
	stacksBy int
	amount   int
}

func (c *CarriableStack) Init(item Carriable, amount int) {
	if amount > item.StacksBy() {
		panic("Stack overflow! (in inventory)")
	}
	c.item = item
	c.stacksBy = item.StacksBy()
	c.amount = amount
}

func (c *CarriableStack) Add(amount int) int {
	if c.amount+amount <= c.stacksBy {
		c.amount += amount
		return 0
	} else {
		adding := c.stacksBy - c.amount
		c.amount = c.stacksBy
		return amount - adding
	}
}

func (c *CarriableStack) Remove(amount int) int {
	if amount <= c.amount {
		c.amount -= amount
		return 0
	} else {
		removing := c.amount
		c.amount = 0
		return amount - removing
	}
}

func (c *CarriableStack) GetName() string { // assuming the name is unique
	return c.item.GetName()
}

func (c *CarriableStack) Use(player *Player, values ...interface{}) {
	c.item.Use(player, values...)
}
