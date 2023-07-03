package main

func isValidInput(input int) bool {
	if input < 1 || input > 9 {
		return false
	}
	return true
}
