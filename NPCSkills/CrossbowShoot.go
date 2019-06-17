package NPCSkills

import . "DunCrawl/Interfaces"

type CrossbowShoot struct {
	CommonEnSkill
	baseDmg int
}

func NewCrossbowShoot(enemy Unit) NPCSkill {
	b := &CrossbowShoot{}
	b.baseDmg = 30
	b.speed = 7
	b.name = "Bite"
	b.wielder = enemy
	b.iconPath = "resources/NPCSkillsIcons/DogBite.PNG"
	b.res = ""
	return b
}

func (c *CrossbowShoot) Apply(r *Room) string {
	if FindEffect(c.wielder, Winded3) {
		RemoveEffect(c.wielder, Winded3)
		return DealDamage(c.wielder, c.target, c.baseDmg)
	}
	return "Thwonk"
}
