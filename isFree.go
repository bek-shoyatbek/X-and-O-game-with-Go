package main

func isFree(n int, board [][]string) bool {
	if n <= 3 {
		if board[0][n-1] != "_" {
			return false
		}
	}
	if n > 3 && n <= 6 {
		if board[1][n-(3+1)] != "_" {
			return false
		}
	}
	if n > 6 && n <= 9 {
		if board[2][n-(6+1)] != "_" {
			return false
		}
	}
	return true
}
