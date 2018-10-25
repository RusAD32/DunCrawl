package Interfaces

import (
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
	numToRemove := make([]int, 0)
	for i, x := range *s {
		if x.GetCD() == 0 {
			numToRemove = append(numToRemove, i-len(numToRemove))
		}
	}
	for _, i := range numToRemove {
		(*s)[i] = nil
		*s = append((*s)[:i], (*s)[i+1:]...)
	}
}

func RemoveEffect(unit Unit, id EffectID) {
	s := unit.GetEffects()
	numToRemove := make([]int, 0)
	for i, x := range *s {
		if x.GetID() == id {
			numToRemove = append(numToRemove, i-len(numToRemove))
		}
	}
	for _, i := range numToRemove {
		(*s)[i] = nil
		*s = append((*s)[:i], (*s)[i+1:]...)
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
	for i, x := range r.Enemies {
		if x.GetHP() == 0 {
			numToRemove = append(numToRemove, i-len(numToRemove))
		}
	}
	for _, i := range numToRemove {
		r.Defeated = append(r.Defeated, r.Enemies[i])
		r.Enemies[i] = nil
		r.Enemies = append(r.Enemies[:i], r.Enemies[i+1:]...)
	}
}
