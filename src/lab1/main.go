package main

import (
	"fmt"
	"strconv"
	"strings"
)

var err error

func isCorrect(str string) bool {
	var length int = len(str)
	var symbols_count int = 0

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

func converToArray(str string) []int {
	var arr []int
	var element int

	for i := 0; i < len(str); i++ {
		element, err = strconv.Atoi(string([]rune(str)[i]))
		arr = append(arr, element)
	}

	return arr
}

func converToString(arr []int) string {
	var str string

	for i := 0; i < len(arr); i++ {
		str = strconv.Itoa(arr[i]) + str
	}

	return str
}

func main() {
	var input, num1, num2, symbol string
	// var number1, number2 []int

	fmt.Println("Enter expression: ")
	fmt.Scanln(&input)

	if isCorrect(input) {
		num1, num2, symbol = getNumbers(input)
		number1 := converToArray(num1)
		number2 := converToArray(num2)
		fmt.Println(number1, number2, symbol)
	} else {
		fmt.Println("Expression is not correct.")
	}
}
