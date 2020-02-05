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
		fmt.Println("Please enter an expression in the following format:\n(num1) (operator) (num2)\nType 'exit' to stop the calculator")
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("goCalculator>")
		var lastResult float64 = 0.0
		for scanner.Scan() {
			var expression string = scanner.Text()
			if expression == "exit" {
				os.Exit(0)
			}
			res, err := processExpression(expression, lastResult)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(expression + " =", res)
			}
			lastResult = res
			fmt.Print("goCalculator>")
		}
	}
}

func processExpression(e string, last float64) (float64, error) {
	result := 0.0
	c := strings.Split(e, " ")
	if last != 0.0 && (c[0] == "*" || c[0] == "/" || c[0] == "+" || c[0] == "-") {
		if len(c)-1 < 1 {
			return 0.0, errors.New("error: missing an argument.")
		}
		num2, err := strconv.ParseFloat(c[1], 64)
		if err != nil {
			return 0.0, errors.New("error: incorrect Argument type.")
		}
		switch c[0] {
		case "*":
			result = last * num2
		case "/":
			if num2 == 0.0 {
				return 0.0, errors.New("error: cannot divide by zero.")
			}
			result = last / num2
		case "+":
			result = last + num2
		case "-":
			result = last - num2
		}
	} else {
		if len(c)-1 < 2 {
			return 0.0, errors.New("error: missing some arguments.")
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
				return 0.0, errors.New("error: cannot divide by zero.")
			}
			result = num1 / num2
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		}
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
