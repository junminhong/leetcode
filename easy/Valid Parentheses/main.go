package main

import "fmt"

var chars = map[string]string{
	"[": "]",
	"{": "}",
	"(": ")",
}

func isValid(s string) bool {
	for i := 0; i < len(s)-1; i += 2 {
		if string(s[i+1]) != chars[string(s[i])] {
			return false
		}
	}
	return true
}

func main() {
	str := "()"
	fmt.Println(isValid(str))
}
