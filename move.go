package main

func Move(board [][]string, n int, xOrO string) [][]string {
	if n <= 3 {
		board[0][n-1] = xOrO
	}
	if n > 3 && n <= 6 {
		board[1][n-(3+1)] = xOrO
	}
	if n > 6 && n <= 9 {
		board[2][n-(6+1)] = xOrO
	}
	if n > 9 {
		return board
	}
	return board
}
