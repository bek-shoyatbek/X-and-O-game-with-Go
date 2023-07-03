package main

import (
	"strings"
)

func findWinner(board [][]string) string {
	str := ""
	for i := range board {
		str = strings.Join(board[i], "")
		switch str {
		case "OOO":
			return "Winner is O"
		case "XXX":
			return "Winner is X"
		}
	}
	arr := make([]string, 5)
	c, y := 0, 0
	for x := 0; x < 3; x++ {
		makeStr := ""
		for i := range board {
			makeStr += board[i][y]
		}
		arr[c] = makeStr
		c++
		y++
	}
	ver, unver := board[0][0]+board[1][1]+board[2][2], board[0][2]+board[1][1]+board[2][0]
	arr[3] = ver
	arr[4] = unver
	return winner(arr)
}

func winner(arr []string) string {
	for i := range arr {
		if arr[i] == "OOO" {
			return "Winner is O"
		}
		if arr[i] == "XXX" {
			return "Winner is X"
		}
	}
	return ""
}
