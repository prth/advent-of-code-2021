package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

const inputFilePath = "input.txt"

type Entry struct {
	patterns []string
	outputs  []string
}

func main() {
	entries, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	knownLengthToPatternMap := map[int]int{
		2: 1,
		4: 4,
		3: 7,
		7: 8,
	}

	answer1 := 0
	answer2 := 0

	for _, entry := range entries {
		knownDigitToPatterns := make(map[int]string)

		for _, pattern := range entry.patterns {
			if digit, ok := knownLengthToPatternMap[len(pattern)]; ok {
				knownDigitToPatterns[digit] = pattern
			}
		}

		knownPatternsToDigit := make(map[string]int)

		for _, pattern := range entry.patterns {
			if _, ok := knownLengthToPatternMap[len(pattern)]; ok {
				continue
			}

			pattern = sortString(pattern)

			if len(pattern) == 6 {
				if !containsPattern(pattern, knownDigitToPatterns[1]) {
					knownPatternsToDigit[pattern] = 6
				} else if containsPattern(pattern, knownDigitToPatterns[4]) {
					knownPatternsToDigit[pattern] = 9
				} else {
					knownPatternsToDigit[pattern] = 0
				}
			} else {
				if containsPattern(pattern, knownDigitToPatterns[1]) {
					knownPatternsToDigit[pattern] = 3
				} else if getCountOfMatchedChar(pattern, knownDigitToPatterns[4]) == 3 {
					knownPatternsToDigit[pattern] = 5
				} else {
					knownPatternsToDigit[pattern] = 2
				}
			}
		}

		outputNumber := 0

		for _, output := range entry.outputs {
			if digit, ok := knownLengthToPatternMap[len(output)]; ok {
				answer1++

				outputNumber = (outputNumber * 10) + digit
			} else {
				pattern := sortString(output)
				digit = knownPatternsToDigit[pattern]

				outputNumber = (outputNumber * 10) + digit
			}
		}

		answer2 += outputNumber
	}

	log.Printf("Answer #1 :: %d", answer1)
	log.Printf("Answer #2 :: %d", answer2)
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)

	return strings.Join(s, "")
}

func containsPattern(pattern string, subPattern string) bool {
	if len(pattern) < len(subPattern) {
		return false
	}

	for _, x := range subPattern {
		if !strings.ContainsRune(pattern, x) {
			return false
		}
	}

	return true
}

func getCountOfMatchedChar(pattern string, targetPattern string) int {
	matchCount := 0

	for _, x := range targetPattern {
		if strings.ContainsRune(pattern, x) {
			matchCount++
		}
	}

	return matchCount
}

func getInput() ([]Entry, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []Entry

	for scanner.Scan() {
		line := scanner.Text()

		entryStrParts := strings.Split(line, "|")

		entry := Entry{
			patterns: strings.Fields(entryStrParts[0]),
			outputs:  strings.Fields(entryStrParts[1]),
		}

		input = append(input, entry)
	}

	return input, scanner.Err()
}
