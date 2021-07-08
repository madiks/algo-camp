package week03

// https://leetcode-cn.com/problems/surrounded-regions/
// https://leetcode-cn.com/submissions/detail/193587903/

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	maxX, maxY := len(board[0]), len(board)
	keepO := make([][]int, maxY)
	for i := 0; i < len(board); i++ {
		keepO[i] = make([]int, maxX)
	}
	// do mark
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			if (x == 0 || x == maxX-1 || y == 0 || y == maxY-1) && board[y][x] == 'O' {
				traversal(board, keepO, x, y, maxX, maxY)
			}
		}
	}
	// update board
	for k := range board {
		for kk := range board[k] {
			if keepO[k][kk] == 0 {
				board[k][kk] = 'X'
			}
		}
	}
}

func traversal(board [][]byte, keepO [][]int, x, y int, maxX, maxY int) {
	if x >= maxX || x < 0 {
		return
	}
	if y >= maxY || y < 0 {
		return
	}
	if board[y][x] == 'X' || keepO[y][x] == 1 {
		return
	}
	keepO[y][x] = 1
	// go left
	traversal(board, keepO, x-1, y, maxX, maxY)
	// go right
	traversal(board, keepO, x+1, y, maxX, maxY)
	// go up
	traversal(board, keepO, x, y-1, maxX, maxY)
	// go down
	traversal(board, keepO, x, y+1, maxX, maxY)
}
