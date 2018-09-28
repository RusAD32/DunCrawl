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
	for true {
		text, err := Input.ReadString('\n')
		if err != nil {
			Errors.Write([]byte(err.Error()))
			return ""
		}
		for _, v := range reqInput {
			length := len(v)
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
}

func DealDamage(from, to Unit, dmg int) string {
	//TODO: триггеры
	//TODO: эффекты
	return strconv.Itoa(to.ChangeHealth(dmg))
}

func HealthUp(from, to Unit, amount int) string {
	return strconv.Itoa(to.ChangeHealth(-amount))
}
