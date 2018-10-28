package UI

import (
	"bufio"
	"os"
	"strings"
)

var Input = bufio.NewReader(os.Stdin)
var Output = bufio.NewWriter(os.Stdout)
var Errors = bufio.NewWriter(os.Stderr)

func Inform(message string) {
	Output.Write([]byte(message))
	Output.Flush()
}

func Prompt(message string, reqInput []string) (string, string) {
	Output.Write([]byte(message))
	Output.Flush()
	for true {
		text, err := Input.ReadString('\n')
		if err != nil {
			Errors.Write([]byte(err.Error()))
			return "", ""
		}
		for _, v := range reqInput {
			length := len(v)
			//fmt.Println(text, text[:length], v, text[:length] == v, len(text) >= length)
			if len(text) >= length && text[:length] == v {
				return v, strings.TrimSpace(text)
			}
		}
	}
	return "", ""
}
