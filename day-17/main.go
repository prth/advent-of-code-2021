package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const inputFilePath = "input.txt"

type TargetArea struct {
	xMin int
	xMax int
	yMin int
	yMax int
}

type Position struct {
	xPos int
	yPos int
	xVel int
	yVel int
}

func main() {
	targetArea, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	answer1 := getHighestYPositon(targetArea)
	answer2 := countNumberOfInitialVelocities(targetArea)

	log.Printf("Answer #1 :: %d", answer1)
	log.Printf("Answer #2 :: %d", answer2)
}

func countNumberOfInitialVelocities(targetArea TargetArea) int {
	count := 0

	for xVel := 0; xVel <= targetArea.xMax; xVel++ {
		for yVel := targetArea.yMin; yVel <= int(math.Abs(float64(targetArea.yMin))); yVel++ {
			position := Position{
				xPos: 0,
				yPos: 0,
				xVel: xVel,
				yVel: yVel,
			}

			for true {
				nextPosition := getNextPosition(position)

				if isPositionInTargetArea(nextPosition, targetArea) {
					count++
					break
				}

				if nextPosition.xVel == 0 && (nextPosition.xPos < targetArea.xMin || nextPosition.xPos > targetArea.xMax) {
					break
				}

				if nextPosition.yPos < targetArea.yMin {
					break
				}

				position = nextPosition
			}
		}
	}

	return count
}

func isPositionInTargetArea(position Position, targetArea TargetArea) bool {
	if position.xPos < targetArea.xMin || position.xPos > targetArea.xMax {
		return false
	}

	if position.yPos < targetArea.yMin || position.yPos > targetArea.yMax {
		return false
	}

	return true
}

func getNextPosition(position Position) Position {
	newPosition := Position{}
	newPosition.xPos = position.xPos + position.xVel
	newPosition.yPos = position.yPos + position.yVel

	if position.xVel > 0 {
		newPosition.xVel = position.xVel - 1
	}

	newPosition.yVel = position.yVel - 1

	return newPosition
}

func getHighestYPositon(targetArea TargetArea) int {
	y := math.Abs(float64(targetArea.yMin))

	return int(((y - 1) * y / 2))
}

func getInput() (TargetArea, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return TargetArea{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var targetArea TargetArea

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line[len("target area: "):], ", ")

		xParts := strings.Split(parts[0][len("x="):], "..")
		yParts := strings.Split(parts[1][len("x="):], "..")

		targetArea.xMin, _ = strconv.Atoi(xParts[0])
		targetArea.xMax, _ = strconv.Atoi(xParts[1])

		targetArea.yMin, _ = strconv.Atoi(yParts[0])
		targetArea.yMax, _ = strconv.Atoi(yParts[1])
	}

	return targetArea, scanner.Err()
}
