package Interfaces

import "errors"

type Stat int

const (
	Strength Stat = iota
	Agility
	Endurance
	Constitution
	Intelligence
	Wisdom
	Luck
)

type Player struct {
	equipment  map[Slot]Equippable
	inventory  []Carriable
	dmgSkills  []PlayerDmgSkill
	selfSkills []PlayerSelfSkill
	lvl        int
	exp        int
	curMentHP  int
	maxMentHP  int
	inv        *Inventory
	money      int
	pet        *Pet
	BasicUnit
}

func (p *Player) GetPet() *Pet {
	return p.pet
}

func (p *Player) SetPet(pet *Pet) {
	p.pet = pet
}

func (p *Player) GetDmgSkillList() []PlayerDmgSkill {
	return p.dmgSkills
}

func (p *Player) GetSelfSkillList() []PlayerSelfSkill {
	return p.selfSkills
}

func (p *Player) ChangeHealth(damage int) int {
	if damage < 0 { // значит, это хил
		p.curHP -= damage
		if p.curHP > p.maxHP {
			p.curHP = p.maxHP
		}
		return -damage
	}
	def := 0
	for _, v := range p.equipment {
		def += v.defence
	}
	if def > 80 {
		def = 80
	}
	damage -= damage * def / 100
	p.curHP -= damage

	if p.curHP < 0 {
		p.curHP = 0
	}
	return damage
}

func (p *Player) GetEquipment() map[Slot]Equippable {
	return p.equipment
}

func (p *Player) Equip(e Equippable, slot Slot) {
	p.equipment[slot] = e
}

func (p *Player) AddSelfSkill(skill PlayerSelfSkill) {
	p.selfSkills = append(p.selfSkills, skill)
}

func (p *Player) AddDmgSkill(skill PlayerDmgSkill) {
	p.dmgSkills = append(p.dmgSkills, skill)
}

func (p *Player) AddToInventory(item Carriable, amount int) error {
	return p.inv.Add(item, amount)
}

func (p *Player) GetInventory() []Stack {
	return p.inv.slots
}

func (p *Player) RemoveFromInv(slot, amount int) {
	p.inv.Remove(slot, amount)
}

func (p *Player) ModifyMoney(amount int) error {
	if p.money < -amount {
		return errors.New("not enough money")
	}
	p.money += amount
	return nil
}

func (p *Player) GetMoney() int {
	return p.money
}

func (p *Player) InventoryFull() bool {
	for _, v := range p.inv.slots {
		if v == nil {
			return true
		}
	}
	return false
}
