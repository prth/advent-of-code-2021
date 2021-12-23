package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const inputFilePath string = "input.txt"

type LocationStack []Location

type Location struct {
	rowIndex    int
	columnIndex int
}

func main() {
	energyGrid, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	answer1 := 0
	answer2 := 0

	numOfStepsForAnswer1 := 100
	trackNumOfSteps := 0

	for true {
		trackNumOfSteps++

		var flashLocationStack LocationStack
		trackResetLocations := make(map[string]Location)

		for rowIndex, row := range energyGrid {
			for columnIndex, _ := range row {
				energyGrid[rowIndex][columnIndex]++

				location := Location{
					rowIndex:    rowIndex,
					columnIndex: columnIndex,
				}

				if energyGrid[rowIndex][columnIndex] == 10 {
					flashLocationStack.Push(location)
				}
			}
		}

		for !flashLocationStack.IsEmpty() {
			locationToCheck, _ := flashLocationStack.Pop()

			if energyGrid[locationToCheck.rowIndex][locationToCheck.columnIndex] <= 9 {
				continue
			}

			if _, ok := trackResetLocations[getLocationKey(locationToCheck)]; ok {
				continue
			}

			trackResetLocations[getLocationKey(locationToCheck)] = locationToCheck

			adjacentLocations := getAdjacentLocations(energyGrid, locationToCheck)

			for _, adjacentLocation := range adjacentLocations {
				energyGrid[adjacentLocation.rowIndex][adjacentLocation.columnIndex]++
				flashLocationStack.Push(adjacentLocation)
			}
		}

		if trackNumOfSteps <= numOfStepsForAnswer1 {
			answer1 += len(trackResetLocations)
		}

		if len(trackResetLocations) == getEnergyGridRowCount(energyGrid)*getEnergyGridColumnCount(energyGrid) {
			answer2 = trackNumOfSteps
			break
		}

		for _, location := range trackResetLocations {
			energyGrid[location.rowIndex][location.columnIndex] = 0
		}
	}

	log.Printf("Answer #1 :: %d", answer1)
	log.Printf("Answer #2 :: %d", answer2)
}

func getAdjacentLocations(energyGrid [][]int, centerLocation Location) []Location {
	var adjacentLocations []Location

	if centerLocation.rowIndex > 0 {
		// Vertical adjacent : above
		adjacentLocations = append(adjacentLocations, Location{
			rowIndex:    centerLocation.rowIndex - 1,
			columnIndex: centerLocation.columnIndex,
		})

		// Diagonal adjacent : top left
		if centerLocation.columnIndex > 0 {
			adjacentLocations = append(adjacentLocations, Location{
				rowIndex:    centerLocation.rowIndex - 1,
				columnIndex: centerLocation.columnIndex - 1,
			})
		}

		// Diagonal adjacent : top right
		if centerLocation.columnIndex < getEnergyGridColumnCount(energyGrid)-1 {
			adjacentLocations = append(adjacentLocations, Location{
				rowIndex:    centerLocation.rowIndex - 1,
				columnIndex: centerLocation.columnIndex + 1,
			})
		}
	}

	if centerLocation.rowIndex < getEnergyGridRowCount(energyGrid)-1 {
		// Vertical adjacent : below
		adjacentLocations = append(adjacentLocations, Location{
			rowIndex:    centerLocation.rowIndex + 1,
			columnIndex: centerLocation.columnIndex,
		})

		// Diagonal adjacent : bottom left
		if centerLocation.columnIndex > 0 {
			adjacentLocations = append(adjacentLocations, Location{
				rowIndex:    centerLocation.rowIndex + 1,
				columnIndex: centerLocation.columnIndex - 1,
			})
		}

		// Diagonal adjacent : bottom right
		if centerLocation.columnIndex < getEnergyGridColumnCount(energyGrid)-1 {
			adjacentLocations = append(adjacentLocations, Location{
				rowIndex:    centerLocation.rowIndex + 1,
				columnIndex: centerLocation.columnIndex + 1,
			})
		}
	}

	// Horizontal adjacent : left
	if centerLocation.columnIndex > 0 {
		adjacentLocations = append(adjacentLocations, Location{
			rowIndex:    centerLocation.rowIndex,
			columnIndex: centerLocation.columnIndex - 1,
		})
	}

	// Horizontal adjacent : right
	if centerLocation.columnIndex < getEnergyGridColumnCount(energyGrid)-1 {
		adjacentLocations = append(adjacentLocations, Location{
			rowIndex:    centerLocation.rowIndex,
			columnIndex: centerLocation.columnIndex + 1,
		})
	}

	return adjacentLocations
}

func getEnergyGridRowCount(energyGrid [][]int) int {
	return len(energyGrid)
}

func getEnergyGridColumnCount(energyGrid [][]int) int {
	return len(energyGrid[0])
}

func getLocationKey(location Location) string {
	return string(location.rowIndex) + "," + string(location.columnIndex)
}

func (stack *LocationStack) IsEmpty() bool {
	return len(*stack) == 0
}

func (stack *LocationStack) Push(l Location) {
	*stack = append(*stack, l)
}

func (stack *LocationStack) Pop() (Location, bool) {

	if stack.IsEmpty() {
		return Location{}, false
	}

	indexToPop := len(*stack) - 1
	location := (*stack)[indexToPop]
	*stack = (*stack)[:indexToPop]

	return location, true
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

		for _, numRune := range line {
			num, _ := strconv.Atoi(string(numRune))
			row = append(row, num)
		}

		input = append(input, row)
	}

	return input, scanner.Err()
}
