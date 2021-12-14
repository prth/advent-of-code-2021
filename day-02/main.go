package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFilePath = "input.txt"
const (
	forward = "forward"
	up      = "up"
	down    = "down"
)

func main() {
	commands, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	horizontalPosition1, depthPosition1 := computePositionsPerManual1(commands)
	answer1 := horizontalPosition1 * depthPosition1

	horizontalPosition2, depthPosition2 := computePositionsPerManual2(commands)
	answer2 := horizontalPosition2 * depthPosition2

	log.Printf("Answer #1 :: %d", answer1)
	log.Printf("Answer #2 :: %d", answer2)
}

func computePositionsPerManual1(commands []string) (int, int) {
	horizontalPosition := 0
	depthPosition := 0

	for _, command := range commands {
		direction, stepCount := decodeCommand(command)

		switch direction {
		case forward:
			horizontalPosition += stepCount
		case up:
			depthPosition -= stepCount
		case down:
			depthPosition += stepCount
		}
	}

	return horizontalPosition, depthPosition
}

func computePositionsPerManual2(commands []string) (int, int) {
	horizontalPosition := 0
	depthPosition := 0
	aim := 0

	for _, command := range commands {
		direction, stepCount := decodeCommand(command)

		switch direction {
		case forward:
			horizontalPosition += stepCount
			depthPosition += aim * stepCount
		case up:
			aim -= stepCount
		case down:
			aim += stepCount
		}
	}

	return horizontalPosition, depthPosition
}

func decodeCommand(command string) (string, int) {
	words := strings.Fields(command)

	stepCount, _ := strconv.Atoi(words[1])

	return words[0], stepCount
}

func getInput() ([]string, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []string

	for scanner.Scan() {
		line := scanner.Text()

		input = append(input, line)
	}

	return input, scanner.Err()
}
