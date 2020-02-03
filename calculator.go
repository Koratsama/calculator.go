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
	expressions := make([]string, 1)
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("gocalc>")
		for scanner.Scan() {
			expressions = append(expressions, scanner.Text())
			//res, err := processStack(expressions)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(res)
			}
			fmt.Print("gocalc>")
		}
	}
}
