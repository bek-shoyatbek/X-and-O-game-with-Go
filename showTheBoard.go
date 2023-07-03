package main

import (
	"fmt"
	"strings"
)

func showTheBoard(arr [][]string) {
	callCleaner()
	fmt.Println()
	for i := range arr {
		fmt.Println(strings.Join(arr[i], " "))
	}
}
