package Interfaces

import "errors"

type Inventory struct {
	slots    []Stack
	maxSlots int
}

func (i *Inventory) Init(size int) {
	i.maxSlots = size
	i.slots = make([]Stack, size)
}

func (i *Inventory) Add(item Carriable, amount int) error {
	if amount > item.StacksBy() {
		return errors.New("trying to add more than one stack of items")
	}
	hasFreeSlots := false
	for _, v := range i.slots {
		hasFreeSlots = hasFreeSlots || v == nil
		if v == nil {
			continue
		}
		if v.GetName() == item.GetName() {
			amount = v.Add(amount)
			if amount == 0 {
				return nil
			}
		}
	}
	if hasFreeSlots {
		stack := new(CarriableStack)
		stack.Init(item, amount)
		for ind := range i.slots {
			if i.slots[ind] == nil {
				i.slots[ind] = stack
				return nil
			}
		}
	}
	return errors.New("no room in inventory")
}

func (i *Inventory) Remove(slot, amount int) {
	res := i.slots[slot].Remove(amount)
	if res == 0 {
		i.slots[slot] = nil
	}
}

func (i *Inventory) Use(slot int, player *Player, values ...interface{}) {
	i.slots[slot].Use(player, values...)
}