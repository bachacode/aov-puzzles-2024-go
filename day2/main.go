package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed puzzle.txt
var input string

func main() {
	var reports, err = parseInput(input)

	if err != nil {
		fmt.Println("An error ocurred parsing the numbers: ", err)
	}
	fmt.Println(partOne(reports))
	fmt.Println(partTwo(reports))
}

func partOne(reports [][]int) int {
	var safe int = len(reports)
	var increasing bool

	for _, report := range reports {
		for i, number := range report {
			if i == 0 {
				increasing = (number < report[1])
			}

			if i == len(report)-1 {
				break
			}

			if increasing && number > report[i+1] {
				safe--
				break
			} else if !increasing && number < report[i+1] {
				safe--
				break
			}

			difference := intAbs(number, report[i+1])
			if difference < 1 || difference > 3 {
				safe--
				break
			}
		}
	}

	return safe
}

func partTwo(reports [][]int) int {
	var safeCount int = 0

	for _, report := range reports {
		for i := 0; i < len(report); i++ {
			isSafe := true
			newReport := make([]int, len(report))
			copy(newReport, report)
			if i == len(report)-1 {
				newReport = newReport[:len(newReport)-1]
			} else {
				newReport = append(newReport[:i], newReport[i+1:]...)
			}
			for j := 0; j < len(newReport)-1; j++ {

				increasing := (newReport[0] < newReport[1])
				level := newReport[j]
				nextLevel := newReport[j+1]

				if isUnsafe(increasing, level, nextLevel) {
					isSafe = false
				}
			}

			if isSafe {
				safeCount++
				break
			}
		}
	}

	return safeCount
}

func isUnsafe(increasing bool, x int, y int) bool {
	diff := intAbs(x, y)

	if increasing && x > y {
		return true
	}

	if !increasing && x < y {
		return true
	}

	if diff < 1 || diff > 3 {
		return true
	}
	return false

}

func intAbs(x int, y int) int {
	if x < y {
		return y - x
	} else {
		return x - y
	}
}

func parseInput(input string) ([][]int, error) {
	var reports [][]int
	for _, line := range strings.Split(input, "\n") {
		var report []int
		for _, number := range strings.Split(line, " ") {
			parseNum, err := strconv.Atoi(number)

			if err != nil {
				return reports, err
			}

			report = append(report, parseNum)
		}
		reports = append(reports, report)
	}

	return reports, nil
}
