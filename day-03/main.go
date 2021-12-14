package main

import (
	"bufio"
	"log"
	"math"
	"os"
)

const inputFilePath = "input.txt"

func main() {
	report, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	gammaRate, epsilonRate := calculateGammaAndEpsilonRates(report)
	answer1 := gammaRate * epsilonRate

	oxygenGenRating, co2ScrubberRating := calculateOxygenGenAndCO2ScrubberRating(report)
	answer2 := oxygenGenRating * co2ScrubberRating

	log.Printf("Answer #1 :: %d", answer1)
	log.Printf("Answer #2 :: %d", answer2)
}

func calculateGammaAndEpsilonRates(report []string) (int, int) {
	trackBits1 := make([]int, len(report[0]))

	for _, binaryNumStr := range report {
		for index, bit := range binaryNumStr {
			if bit == '1' {
				trackBits1[index]++
			}
		}
	}

	gammaRate := 0
	epsilonRate := 0

	for i := 0; i < len(report[0]); i++ {
		if trackBits1[i] > len(report)/2 {
			gammaRate += mathPowInt(2, len(report[0])-i-1)
		} else {
			epsilonRate += mathPowInt(2, len(report[0])-i-1)
		}
	}

	return gammaRate, epsilonRate
}

// TODO This implementation can be refactored!
func calculateOxygenGenAndCO2ScrubberRating(report []string) (int, int) {
	oxygenSelectedIndex := -1
	co2SelectedIndex := -1

	oxygenAvailableStrings := make([]bool, len(report))
	co2AvailableStrings := make([]bool, len(report))

	for charIndex := 0; charIndex < len(report[0]); charIndex++ {
		trackOxyBit1 := 0
		oxyAvailableCount := 0

		trackCO2Bit1 := 0
		co2AvailableCount := 0

		for reportIndex, binaryNum := range report {
			bit := string(binaryNum[charIndex])

			if charIndex == 0 {
				oxyAvailableCount++
				co2AvailableCount++

				if bit == "1" {
					trackOxyBit1++
					trackCO2Bit1++
				}
			} else {
				if oxygenAvailableStrings[reportIndex] {
					oxyAvailableCount++
					if bit == "1" {
						trackOxyBit1++
					}
				}

				if co2AvailableStrings[reportIndex] {
					co2AvailableCount++
					if bit == "1" {
						trackCO2Bit1++
					}
				}
			}
		}

		var commonOxyBit string
		if float64(trackOxyBit1) >= float64(oxyAvailableCount)/2 {
			commonOxyBit = "1"
		} else {
			commonOxyBit = "0"
		}

		var leastC02Bit string
		if float64(trackCO2Bit1) >= float64(co2AvailableCount)/2 {
			leastC02Bit = "0"
		} else {
			leastC02Bit = "1"
		}

		for reportIndex, binaryNum := range report {
			bit := string(binaryNum[charIndex])

			if charIndex == 0 || oxygenAvailableStrings[reportIndex] {
				if oxyAvailableCount == 1 {
					oxygenSelectedIndex = reportIndex
				} else if oxygenSelectedIndex == -1 {
					oxygenAvailableStrings[reportIndex] = bit == commonOxyBit
				}
			}

			if charIndex == 0 || co2AvailableStrings[reportIndex] {
				if co2AvailableCount == 1 {
					co2SelectedIndex = reportIndex
				} else if co2SelectedIndex == -1 {
					co2AvailableStrings[reportIndex] = bit == leastC02Bit
				}
			}
		}

		if oxygenSelectedIndex != -1 && co2SelectedIndex != -1 {
			break
		}
	}

	if oxygenSelectedIndex == -1 || co2SelectedIndex == -1 {
		for i := 0; i < len(report); i++ {
			if oxygenAvailableStrings[i] {
				oxygenSelectedIndex = i
			}

			if co2AvailableStrings[i] {
				co2SelectedIndex = i
			}
		}
	}

	oxygenGenRating := convertBinaryToInt(report[oxygenSelectedIndex])
	co2ScrubberRating := convertBinaryToInt(report[co2SelectedIndex])

	return oxygenGenRating, co2ScrubberRating
}

func convertBinaryToInt(binaryNum string) int {
	value := 0

	for i, bit := range binaryNum {
		if bit == '1' {
			value += mathPowInt(2, len(binaryNum)-i-1)
		}
	}

	return value
}

func mathPowInt(x int, y int) int {
	return int(math.Pow(float64(x), float64(y)))
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
