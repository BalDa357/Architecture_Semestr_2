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

func getNumbers(str string) (string, string, string) {
	var num []string
	var symbol string

	if strings.Contains(str, "+") {
		symbol = "+"
	} else {
		symbol = "-"
	}
	num = strings.Split(str, symbol)
	for len(num[0]) > len(num[1]) {
		num[1] = "0" + num[1]
	}
	for len(num[0]) < len(num[1]) {
		num[0] = "0" + num[0]
	}
	return num[0], num[1], "-"
}

func main() {
	var input, num1, num2, symbol string

	fmt.Println("Enter expression: ")
	fmt.Scanln(&input)

	if isCorrect(input) {
		num1, num2, symbol = getNumbers(input)
		fmt.Println(num1, num2, symbol)
	} else {
		fmt.Println("Expression is not correct.")
	}
}
