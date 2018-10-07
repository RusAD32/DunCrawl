package Interfaces

import (
	"bufio"
	"os"
	"strconv"
)

var Input = bufio.NewReader(os.Stdin)
var Output = bufio.NewWriter(os.Stdout)
var Errors = bufio.NewWriter(os.Stderr)

func Prompt(message string, reqInput []string) string {
	Output.Write([]byte(message))
	Output.Flush()
	for true {
		text, err := Input.ReadString('\n')
		if err != nil {
			Errors.Write([]byte(err.Error()))
			return ""
		}
		for _, v := range reqInput {
			length := len(v)
			//fmt.Println(text, text[:length], v, text[:length] == v, len(text) >= length)
			if len(text) >= length && text[:length] == v {
				return v
			}
		}
	}
	return ""
}

func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func MakeStrRange(min, max int) []string {
	a := make([]string, max-min+1)
	for i := range a {
		a[i] = strconv.Itoa(min + i)
	}
	return a
}

func Inform(message string) {
	Output.Write([]byte(message))
	Output.Flush()
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

func RemoveDeadEnemies(f *Fight) {
	numToRemove := make([]int, 0)
	for i, x := range f.enemies {
		if x.GetHP() == 0 {
			numToRemove = append(numToRemove, i-len(numToRemove))
		}
	}
	for _, i := range numToRemove {
		f.defeated = append(f.defeated, f.enemies[i])
		f.enemies[i] = nil
		f.enemies = append(f.enemies[:i], f.enemies[i+1:]...)
	}
}
