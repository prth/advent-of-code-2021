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

func main() {
	targetArea, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	answer1 := getHighestYPositon(targetArea)
	answer2 := 0

	log.Printf("Answer #1 :: %d", answer1)
	log.Printf("Answer #2 :: %d", answer2)
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
