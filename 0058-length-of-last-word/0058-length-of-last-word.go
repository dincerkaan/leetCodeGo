func lengthOfLastWord(s string) int {
    str := strings.Split(s, " ")
    count := len(str[len(str)-1])

    for i := 0; i < len(str); i++ {
		if count == 0 {
			count = len(str[len(str)-i-1])
		}
	}

    return count
}