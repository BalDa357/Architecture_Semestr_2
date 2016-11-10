package main

import (
	"fmt"
	"strconv"
	"strings"
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
	if symbols_count == 0 {
		return false
	}
	return true
}

func get_numbers(str string) (string, string, string) {
	var num []string

	if strings.Contains(str, "+") {
		num = strings.Split(str, "+")
		fmt.Println(num)
		return num[0], num[1], "+"
	} else {
		num = strings.Split(str, "-")
		return num[0], num[1], "-"
	}

}

func main() {
	var input, num1, num2, symbol string

	fmt.Println("Enter expression: ")
	fmt.Scanln(&input)

	if isCorrect(input) {
		num1, num2, symbol = get_numbers(input)

	} else {
		fmt.Println("Expression is not correct.")
	}
}
