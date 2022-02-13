package main

func main() {
	strs := []string{"flower", "flow", "flight"}
	longestCommonPrefix(strs)
}

func prefix(result *string, str string, size int) {
	tmp := ""
	for j := 0; j < size; j++ {
		resultTmp := *result
		if resultTmp[j] == str[j] {
			tmp += string(str[j])
		} else {
			break
		}
	}
	*result = tmp
}

func longestCommonPrefix(strs []string) string {
	result := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(result) > len(strs[i]) {
			prefix(&result, strs[i], len(strs[i]))
		} else {
			prefix(&result, strs[i], len(result))
		}
	}
	return result
}
