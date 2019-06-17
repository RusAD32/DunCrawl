package NPCSkills

import (
	. "DunCrawl/Effects"
	. "DunCrawl/Interfaces"
)

type CrossbowWind1 struct {
	CommonEnSkill
}

type CrossbowWind2 struct {
	CommonEnSkill
}

type CrossbowWind3 struct {
	CommonEnSkill
}

func (c *CrossbowWind1) Apply(r *Room) string {
	eff := NewCrossbowWind1Effect()
	c.wielder.AddEffect(eff)
	return eff.GetInfo()
}

func (c *CrossbowWind2) Apply(r *Room) string {
	if !FindEffect(c.wielder, Winded1) {
		return "Nothing happened"
	}
	RemoveEffect(c.wielder, Winded1)
	eff := NewCrossbowWind2Effect()
	c.wielder.AddEffect(eff)
	return eff.GetInfo()
}

func (c *CrossbowWind3) Apply(r *Room) string {
	if !FindEffect(c.wielder, Winded2) {
		return "Nothing happened"
	}
	RemoveEffect(c.wielder, Winded2)
	eff := NewCrossbowWind3Effect()
	c.wielder.AddEffect(eff)
	return eff.GetInfo()
}

func NewCrossbowWind1(enemy Unit) NPCSkill {
	b := &CrossbowWind1{}
	b.speed = 3
	b.name = "Crossbow winding start"
	b.wielder = enemy
	b.iconPath = "resources/NPCSkillsIcons/DogBite.PNG"
	b.res = ""
	return b
}

func NewCrossbowWind2(enemy Unit) NPCSkill {
	b := &CrossbowWind2{}
	b.speed = 3
	b.name = "Crossbow winding continuation"
	b.wielder = enemy
	b.iconPath = "resources/NPCSkillsIcons/DogBite.PNG"
	b.res = ""
	return b
}

func NewCrossbowWind3(enemy Unit) NPCSkill {
	b := &CrossbowWind3{}
	b.speed = 3
	b.name = "Crossbow winding continuation"
	b.wielder = enemy
	b.iconPath = "resources/NPCSkillsIcons/DogBite.PNG"
	b.res = ""
	return b
}
