package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("goCalculator>")
		for scanner.Scan() {
			var expression string = scanner.Text()
			if expression == "exit" {
				os.Exit(0)
			}
			res, err := processExpression(expression)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(expression + " =", res)
			}
			fmt.Print("goCalculator>")
		}
	}
}

func processExpression(e string) (float64, error) {
	result := 0.0
	c := strings.Split(e, " ")
	if len(c)-1 < 2 {
		return 0.0, errors.New("error: some arguments are not supplied")
	}
	num1, num2, err := parseArgs(c)
	if err != nil {
		return 0.0, err
	}
	switch c[1] {
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0.0 {
			return 0.0, errors.New("error: you tried to divide by zero.")
		}
		result = num1 / num2
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	}
	return result, nil
}

func parseArgs(c []string) (float64, float64, error) {
	num1, err := strconv.ParseFloat(c[0], 64)
	if err != nil {
		return 0.0, 0.0, err
	}
	num2, err := strconv.ParseFloat(c[2], 64)
	if err != nil {
		return 0.0, 0.0, err
	}
	return num1, num2, nil
}
