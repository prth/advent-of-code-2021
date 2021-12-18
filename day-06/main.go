package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFilePath = "input.txt"

const RESET_TIMER_VALUE int = 6
const FIRST_TIMER_VALUE int = 8

func main() {
	timers, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	answer1 := computeNumberOfFishUnoptimised(timers, 80)
	answer2 := computeNumberOfFish(timers, 256)

	log.Printf("Answer #1 :: %d", answer1)
	log.Printf("Answer #2 :: %d", answer2)
}

func computeNumberOfFishUnoptimised(initialStateTimers []int, noOfDays int) int {
	timers := make([]int, len(initialStateTimers))
	copy(timers, initialStateTimers)

	remainingDays := noOfDays
	trackMinimumTimeValue := 0
	reduceTimerBy := 1

	for remainingDays > 0 {
		addNewTimers := 0

		for index, timer := range timers {
			if timer == 0 {
				timers[index] = RESET_TIMER_VALUE
				addNewTimers++
			} else {
				timers[index] -= reduceTimerBy
			}

			if index == 0 {
				trackMinimumTimeValue = timers[index]
			} else if trackMinimumTimeValue > timers[index] {
				trackMinimumTimeValue = timers[index]
			}
		}

		for i := 0; i < addNewTimers; i++ {
			timers = append(timers, FIRST_TIMER_VALUE)
		}

		remainingDays -= reduceTimerBy
	}

	return len(timers)
}

func computeNumberOfFish(initialStateTimers []int, noOfDays int) int {
	remainingDays := noOfDays
	timersCountMap := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
	}

	for _, timer := range initialStateTimers {
		timersCountMap[timer] += 1
	}

	for remainingDays >= 7 {
		addToZero := timersCountMap[7]
		addToOne := timersCountMap[8]

		timersCountMap[7] = timersCountMap[5]
		timersCountMap[8] = timersCountMap[6]

		for i := 6; i >= 0; i-- {
			timersCountMap[i] += timersCountMap[i-2]
		}

		timersCountMap[0] += addToZero
		timersCountMap[1] += addToOne

		remainingDays -= 7
	}

	for i := 0; i < remainingDays; i++ {
		timersCountMap[i] *= 2
	}

	total := 0
	for _, timerCount := range timersCountMap {
		total += timerCount
	}

	return total
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
