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
	for i := len(r.enemies) - 1; i >= 0; i-- {
		if r.enemies[i].GetHP() == 0 {
			r.enemies[i].GetDeathTrigger().Call()
			r.defeated = append(r.defeated, r.enemies[i])
			r.enemies[i] = nil
			r.enemies = append(r.enemies[:i], r.enemies[i+1:]...)
		}
	}
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
