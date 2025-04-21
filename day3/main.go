package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed puzzle.txt
var input string

func main() {
	fmt.Println(partOne())
}

func partOne() int {
	mulReg := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := mulReg.FindAllStringSubmatch(input, -1)
	var sum int = 0
	for _, match := range matches {
		num1, err1 := strconv.Atoi(match[1])
		num2, err2 := strconv.Atoi(match[2])

		if err1 != nil || err2 != nil {
			fmt.Println("Error parsing the string to numbers: ", err1, err2)
			break
		}

		sum += num1 * num2

	}

	return sum
}
