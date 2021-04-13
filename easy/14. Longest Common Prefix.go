package main

import "fmt"

func main() {
	strAry := []string{"flower", "flow", "flight"}
	longestCommonPrefix(strAry)
}

func longestCommonPrefix(strAry []string) string {
	prefix := ""
	for _, value := range strAry {
		if prefix == "" {
			prefix = value
		} else {
			for _, prefixValue := range prefix {
				for _, valueValue := range value {
					fmt.Println(prefixValue)
					fmt.Println(valueValue)
				}
			}
		}
	}
	return ""
}
