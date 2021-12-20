package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const inputFilePath = "input.txt"

func main() {
	heightmap, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	riskLevelOfLowPoints := 0
	var basinSizes []int

	for rowIndex, row := range heightmap {
		for columnIndex, height := range row {
			if rowIndex > 0 && height >= heightmap[rowIndex-1][columnIndex] {
				continue
			}

			if rowIndex < len(heightmap)-1 && height >= heightmap[rowIndex+1][columnIndex] {
				continue
			}

			if columnIndex > 0 && height >= heightmap[rowIndex][columnIndex-1] {
				continue
			}

			if columnIndex < len(heightmap[0])-1 && height >= heightmap[rowIndex][columnIndex+1] {
				continue
			}

			riskLevelOfLowPoints += height + 1

			basinSize := findBasinSize(heightmap, []int{rowIndex, columnIndex})
			basinSizes = append(basinSizes, basinSize)
		}
	}

	answer1 := riskLevelOfLowPoints

	answer2 := 1
	sort.Ints(basinSizes)
	for i := 1; i <= 3; i++ {
		answer2 *= basinSizes[len(basinSizes)-i]
	}

	log.Printf("Answer #1 :: %d", answer1)
	log.Printf("Answer #2 :: %d", answer2)
}

func findBasinSize(heightmap [][]int, lowPoint []int) int {
	basinMap := make(map[string]bool)

	trackPoints := [][]int{lowPoint}

	for len(trackPoints) > 0 {
		length := len(trackPoints)
		point := trackPoints[length-1]
		trackPoints = trackPoints[:length-1]

		coordinateStr := getPointCoordinatesStr(point)
		if _, ok := basinMap[coordinateStr]; !ok {
			basinMap[coordinateStr] = true
		}

		pointsToConsider := [][]int{}
		if point[0] > 0 {
			pointsToConsider = append(pointsToConsider, []int{point[0] - 1, point[1]})
		}

		if point[0] < len(heightmap)-1 {
			pointsToConsider = append(pointsToConsider, []int{point[0] + 1, point[1]})
		}

		if point[1] > 0 {
			pointsToConsider = append(pointsToConsider, []int{point[0], point[1] - 1})
		}

		if point[1] < len(heightmap[0])-1 {
			pointsToConsider = append(pointsToConsider, []int{point[0], point[1] + 1})
		}

		for _, point := range pointsToConsider {
			if evaluatePointForBasin(heightmap, basinMap, point) {
				trackPoints = append(trackPoints, point)
			}
		}
	}

	return len(basinMap)
}

func getPointCoordinatesStr(point []int) string {
	return strconv.Itoa(point[0]) + "," + strconv.Itoa(point[1])
}

func evaluatePointForBasin(heightmap [][]int, basinMap map[string]bool, point []int) bool {
	if heightmap[point[0]][point[1]] == 9 {
		return false
	}

	coordinateStr := getPointCoordinatesStr(point)

	_, ok := basinMap[coordinateStr]

	return !ok
}

func getInput() ([][]int, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input [][]int

	for scanner.Scan() {
		line := scanner.Text()

		var row []int

		for _, numStr := range strings.Split(line, "") {
			num, _ := strconv.Atoi(numStr)
			row = append(row, num)
		}

		input = append(input, row)
	}

	return input, scanner.Err()
}
