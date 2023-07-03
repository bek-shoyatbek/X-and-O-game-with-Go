package main

import (
	"fmt"
	"os"
)

func restartGame() {
	var restart int
	fmt.Println("To restart , Enter 1 else 0")
	fmt.Scan(&restart)
	if restart == 0 {
		os.Exit(0)
	}
}
