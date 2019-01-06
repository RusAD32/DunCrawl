package Interfaces

import (
	"bytes"
	"fmt"
	"strconv"
)

func MakeStrRange(min, max int) []string {
	a := make([]string, max-min+1)
	for i := range a {
		a[i] = strconv.Itoa(min + i)
	}
	return a
}

func DealDamage(from, to Unit, dmg int) string {
	//TODO: триггеры
	//TODO: эффекты
	res := to.GetDamageTrigger().Call(from)
	return strconv.Itoa(to.ChangeHealth(dmg)) + res
}

func DealRawDamage(to Unit, dmg int) string {
	return strconv.Itoa(to.ChangeHealth(dmg))
}

func HealthUp(from, to Unit, amount int) string {
	return strconv.Itoa(to.ChangeHealth(-amount))
}

func AddEffect(unit Unit, effect Effect) {
	effects := unit.GetEffects()
	*effects = append(*effects, effect)
}

func RemoveExpiredEffects(unit Unit) {
	s := unit.GetEffects()
	for i := len(*s) - 1; i >= 0; i-- {
		if (*s)[i].GetCD() == 0 {
			(*s)[i] = nil
			*s = append((*s)[:i], (*s)[i+1:]...)
		}
	}
}

func RemoveSkillsOfDeadUnits(pq *PriorityQueue) {
	for i := len(*pq) - 1; i >= 0; i-- {
		if !(*pq)[i].(Skill).GetWielder().IsAlive() {
			(*pq)[i] = nil
			*pq = append((*pq)[:i], (*pq)[i+1:]...)
		}
	}
}

func RemoveEffect(unit Unit, id EffectID) {
	s := unit.GetEffects()
	for i := len(*s) - 1; i >= 0; i-- {
		if (*s)[i].GetID() == id {
			(*s)[i] = nil
			*s = append((*s)[:i], (*s)[i+1:]...)
		}
	}
}

func FindEffect(unit Unit, id EffectID) bool {
	for _, v := range *unit.GetEffects() {
		if v.GetID() == id {
			return true
		}
	}
	return false
}

func RemoveDeadEnemies(r *Room) {
	numToRemove := make([]int, 0)
	for i, x := range r.enemies {
		if x.GetHP() == 0 {
			numToRemove = append(numToRemove, i-len(numToRemove))
		}
	}
	for _, i := range numToRemove {
		r.defeated = append(r.defeated, r.enemies[i])
		r.enemies[i] = nil
		r.enemies = append(r.enemies[:i], r.enemies[i+1:]...)
	}
}

func GetDefaultPlayer() *Player {
	inv := Inventory{}
	inv.Init(8)
	du := BasicUnit{
		name:            "you",
		stats:           map[Stat]int{},
		effects:         []Effect{},
		dmgTakenTrigger: new(Trigger).Init(),
		curHP:           100,
		maxHP:           100,
	}
	return &Player{
		equipment:  map[Slot]Equippable{},
		inventory:  []Carriable{},
		dmgSkills:  []PlayerDmgSkill{},
		selfSkills: []PlayerSelfSkill{},
		lvl:        1,
		exp:        0,
		curMentHP:  100,
		maxMentHP:  100,
		money:      0,
		inv:        &inv,
		BasicUnit:  du,
	}
}

func GetDefaultLabyrinth(width, length, FirstDirection int, corners []int) *Labyrinth {
	// Move this to Init?
	l := Labyrinth{
		nil,
		make([]*Room, 0),
		make([]*[]*Room, 0),
		width/2*length + length/2, // center
		nil,
		FirstDirection,
		make(chan bool),
		make(chan []SkillInfo),
		make(chan string),
		make(chan Event),
		length,
		width,
		corners,
		Initializing,
	}
	return &l
}

var PlayerDir = []string{">", "v", "<", "^"}

func PrintLabyrinth(l *Labyrinth) {
	labMap := bytes.Buffer{}
	labMap.WriteString("꜒")
	for j := 0; j < l.length; j++ {
		if l.rooms[j].GetNeighbours()[int(Forward)].CanGoThrough() {
			labMap.WriteString(" ")
		} else {
			labMap.WriteString("-")
		}
		if j != l.length-1 {
			labMap.WriteString("-")
		} else {
			labMap.WriteString("˥")
		}
	}
	labMap.WriteString("\n")
	for i := 0; i < l.width; i++ {
		if l.rooms[i*l.length].GetNeighbours()[int(Left)].CanGoThrough() {
			labMap.WriteString(" ")
		} else {
			labMap.WriteString("|")
		}
		for j := 0; j < l.length; j++ {
			curRoom := l.rooms[i*l.length+j]
			if curRoom == l.current {
				labMap.WriteString(PlayerDir[l.previous])
			} else {
				labMap.WriteString(" ")
			}
			if curRoom.GetNeighbours()[int(Right)].CanGoThrough() {
				labMap.WriteString(" ")
			} else {
				labMap.WriteString("|")
			}
		}
		labMap.WriteString("\n")
		if i == l.width-1 {
			labMap.WriteString("꜖")
		} else {
			labMap.WriteString("꜔")
		}
		for j := 0; j < l.length; j++ {
			curRoom := l.rooms[i*l.length+j]
			if curRoom.GetNeighbours()[int(Back)].CanGoThrough() {
				labMap.WriteString(" ")
			} else {
				labMap.WriteString("-")
			}
			if j == l.width-1 {
				labMap.WriteString("˧")
			} else {
				labMap.WriteString("÷")
			}
		}
		labMap.WriteString("\n")
	}
	fmt.Println(labMap.String())
}
