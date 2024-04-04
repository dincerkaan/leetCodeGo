func exist(board [][]byte, word string) bool {
	for i, row := range board {
		for j, char := range row {
			if char == word[0] && dfs(board, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i, j, index int, word string) bool {
	if index == len(word) {
		return true
	}

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[index] {
		return false
	}

	temp := board[i][j]
	board[i][j] = 0 

	found := dfs(board, i+1, j, index+1, word) ||
		dfs(board, i-1, j, index+1, word) ||
		dfs(board, i, j+1, index+1, word) ||
		dfs(board, i, j-1, index+1, word)

	board[i][j] = temp 

	return found
}