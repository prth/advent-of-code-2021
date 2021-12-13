package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const inputFilePath = "input.txt"

func main() {
	input, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	answer1 := 0
	answer2 := 0

	measurementWindowSize := 3

	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			answer1 += 1
		}

		if i < len(input)-(measurementWindowSize-1) {
			if input[i-1] < input[i+measurementWindowSize-1] {
				answer2 += 1
			}
		}
	}

	log.Printf("Answer #1 :: %d", answer1)
	log.Printf("Answer #2 :: %d", answer2)
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

		num, _ := strconv.Atoi(line)
		input = append(input, num)
	}

	return input, scanner.Err()
}
