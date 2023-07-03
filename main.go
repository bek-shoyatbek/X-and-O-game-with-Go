package main

import (
	"fmt"
	"time"
)

func main() {
	callCleaner()
Restart:
	var board = [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	start("Game Started")
	role := pickRole()
	if role != "" {
		showTheBoard(board)
		maxMove := 9
		for maxMove > 0 {
			var inp int
		getInput:
			fmt.Scan(&inp)
			isFree := isFree(inp, board)
			if isFree {
				board = Move(board, inp, role)
				DoWeHaveWinner := findWinner(board)
				if DoWeHaveWinner != "" {
					fmt.Println(DoWeHaveWinner)
					restartGame()
					goto Restart
				}
				showTheBoard(board)
				maxMove--
			} else {
				fmt.Println("It's not empty")
				goto getInput
			}
			if role == "X" {
				role = "O"
			} else {
				role = "X"
			}

		}
		fmt.Println("Draw")

	} else {
		fmt.Println("Wrong input")
		time.Sleep(1 * time.Second)
		goto Restart
	}
}
