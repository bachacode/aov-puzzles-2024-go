package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	var file, err = os.Open("day1/puzzle.txt")

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	var scanner = bufio.NewScanner(file)
	var leftList []int
	var rightList []int

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "   ")
		leftNum, err := strconv.Atoi(split[0])

		if err != nil {
			fmt.Println("Error parsing left character into int: ", err)
			return
		}

		rightNum, err := strconv.Atoi(split[1])

		if err != nil {
			fmt.Println("Error parsing right character into int: ", err)
			return
		}

		leftList = append(leftList, int(leftNum))
		rightList = append(rightList, int(rightNum))
	}

	slices.SortFunc(leftList, func(a, b int) int {
		return cmp.Compare(a, b)
	})

	slices.SortFunc(rightList, func(a, b int) int {
		return cmp.Compare(a, b)
	})

	totalDistance := partOne(leftList, rightList)

	println(totalDistance)

	similarityScore := partTwo(leftList, rightList)

	println(similarityScore)

}

func partOne(leftList []int, rightList []int) int {
	var totalDistance int = 0

	for i := range len(leftList) {
		if leftList[i] < rightList[i] {
			totalDistance += rightList[i] - leftList[i]

		} else {
			totalDistance += leftList[i] - rightList[i]
		}

	}

	return totalDistance
}

func partTwo(leftList []int, rightList []int) int {
	var similarityScore int = 0

	for i := range len(leftList) {
		var multiplier = 0
		for j := range len(rightList) {
			if leftList[i] == rightList[j] {
				multiplier++
			}
		}
		similarityScore += leftList[i] * multiplier

	}
	return similarityScore
}
