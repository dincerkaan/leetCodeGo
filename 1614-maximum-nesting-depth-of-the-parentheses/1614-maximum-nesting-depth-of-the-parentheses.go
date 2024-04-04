func maxDepth(s string) int {
	depth := 0
	maxDepth := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			depth++
		} else if s[i] == ')' {
			depth--
		}

		if depth > maxDepth {
			maxDepth = depth
		}
	}

	return maxDepth
}