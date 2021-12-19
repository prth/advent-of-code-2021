package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const inputFilePath = "input.txt"

func main() {
	input, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	sort.Ints(input)

	medianPosition, medianTargetCost := getMedianPositionWithCost(input)

	answer1 := medianTargetCost
	answer2 := calculateCostForPartTwo(input, medianPosition)

	log.Printf("Answer #1 :: %d", answer1)
	log.Printf("Answer #2 :: %d", answer2)
}

func calculateCostForPartTwo(positions []int, startingPosition int) int {
	cost := 0
	costMap := make(map[int]int)

	i := positions[startingPosition]

	for true {
		currCost, ok := costMap[i]
		if !ok {
			currCost = computeCostPartTwo(positions, i)
		}

		var leftCost int
		var rightCost int

		if i > 0 {
			var ok bool

			leftCost, ok = costMap[i-1]
			if !ok {
				leftCost = computeCostPartTwo(positions, i-1)
			}
		}

		if i < len(positions)-1 {
			var ok bool

			rightCost, ok = costMap[i+1]
			if !ok {
				rightCost = computeCostPartTwo(positions, i+1)
			}
		}

		if currCost < rightCost && currCost < leftCost {
			cost = currCost
			break
		} else {
			if leftCost < rightCost {
				i = i - 1
			} else {
				i = i + 1
			}
		}
	}

	return cost
}

func getMedianPositionWithCost(positions []int) (int, int) {
	var medianPosition int
	var targetCost int

	if len(positions)%2 == 1 {
		midIndex := (len(positions) + 1) / 2

		targetCost = computeCostPartOne(positions, positions[midIndex])
		medianPosition = midIndex
	} else {
		midIndex1 := len(positions) / 2
		midIndex2 := midIndex1 + 1

		if positions[midIndex1] == positions[midIndex2] {
			targetCost = computeCostPartOne(positions, positions[midIndex1])
			medianPosition = midIndex1
		} else {
			minCost := 0
			var targetIndex int

			for i := midIndex1; i <= midIndex2; i++ {
				cost := computeCostPartOne(positions, positions[midIndex1])

				if i == midIndex1 {
					minCost = cost
					targetIndex = midIndex1
				} else if minCost > cost {
					minCost = cost
					targetIndex = midIndex1
				}
			}

			targetCost = minCost
			medianPosition = targetIndex
		}
	}

	return medianPosition, targetCost
}

func computeCostPartOne(positions []int, targetPosition int) int {
	cost := 0

	for _, position := range positions {
		cost += int(math.Abs(float64(position - targetPosition)))
	}

	return cost
}

func computeCostPartTwo(positions []int, targetPosition int) int {
	cost := 0

	for _, position := range positions {
		numOfSteps := int(math.Abs(float64(position - targetPosition)))

		for i := 1; i <= numOfSteps; i++ {
			cost += i
		}
	}

	return cost
}

func getInput() ([]int, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []int

	for scanner.Scan() {
		line := scanner.Text()

		for _, char := range strings.Split(line, ",") {
			num, _ := strconv.Atoi(char)
			input = append(input, num)
		}
	}

	return input, scanner.Err()
}
