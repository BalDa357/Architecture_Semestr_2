package main

import (
	"fmt"
	"strconv"
)

func isCorrect(str string) bool {
	var length int = len(str)
	var symbols_count int = 0
	var err error

	for i := 0; i < length; i++ {
		if string([]rune(str)[i]) == "+" || string([]rune(str)[i]) == "-" {
			symbols_count++
			if symbols_count > 1 || i == 0 || i == length {
				return false
			}
		} else {
			_, err = strconv.Atoi(string([]rune(str)[i]))
			if err != nil {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Print("Enter expression: ")
	var input string
	fmt.Scanln(&input)
	fmt.Println(isCorrect(input))
}
