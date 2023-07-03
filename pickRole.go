package main

import (
	"fmt"
	"os"
	"strings"
)

func pickRole() string {
	var role string
	fmt.Fprintln(os.Stderr, "Choose X or O")
	fmt.Scan(&role)
	role = strings.Trim(role, "")
	if role == "O" || role == "X" {
		return role
	} else {
		return ""
	}
}
