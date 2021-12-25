package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const inputFilePath = "input.txt"

func main() {
	polymer, pairInsertionRulesMap, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	noOfSteps1 := 10
	mostCommonElementScore1, leastCommonElementScore1 := calculateScoresApproach1(polymer, pairInsertionRulesMap, noOfSteps1)
	answer1 := mostCommonElementScore1 - leastCommonElementScore1

	noOfSteps2 := 40
	mostCommonElementScore2, leastCommonElementScore2 := calculateScoresApproach2(polymer, pairInsertionRulesMap, noOfSteps2)
	answer2 := mostCommonElementScore2 - leastCommonElementScore2

	log.Printf("Answer #1 :: %d", answer1)
	log.Printf("Answer #2 :: %d", answer2)
}

func calculateScoresApproach1(polymer string, pairInsertionRulesMap map[string]string, noOfSteps int) (int, int) {
	var mostCommonElementScore int
	var leastCommonElementScore int

	for step := 0; step < noOfSteps; step++ {
		tempPolymer := polymer
		charactersAddedInBetween := 0

		elementScoreMap := make(map[string]int)
		var tempMostCommonElement *string = new(string)
		var tempLeastCommonElement *string = new(string)

		for i := 0; i < len(polymer)-1; i++ {
			currentCharToCount := string(polymer[i])
			trackMostAndLeastCharacters(currentCharToCount, 1, elementScoreMap, tempMostCommonElement, tempLeastCommonElement)

			targetSubstring := polymer[i : i+2]

			if ruleChar, ok := pairInsertionRulesMap[targetSubstring]; ok {
				tempPolymer = insertCharacter(tempPolymer, ruleChar, i+charactersAddedInBetween+1)
				charactersAddedInBetween++
				trackMostAndLeastCharacters(ruleChar, 1, elementScoreMap, tempMostCommonElement, tempLeastCommonElement)
			}

			if i == len(polymer)-2 {
				lastCharToCount := string(polymer[i+1])
				trackMostAndLeastCharacters(lastCharToCount, 1, elementScoreMap, tempMostCommonElement, tempLeastCommonElement)
			}
		}

		polymer = tempPolymer
		mostCommonElementScore = elementScoreMap[*tempMostCommonElement]
		leastCommonElementScore = elementScoreMap[*tempLeastCommonElement]
	}

	return mostCommonElementScore, leastCommonElementScore
}

func calculateScoresApproach2(polymer string, pairInsertionRulesMap map[string]string, noOfSteps int) (int, int) {
	countPairsMap := make(map[string]int)
	lastCharacter := rune(polymer[len(polymer)-1])

	for i := 0; i < len(polymer)-1; i++ {
		targetSubstring := polymer[i : i+2]

		if _, ok := countPairsMap[targetSubstring]; !ok {
			countPairsMap[targetSubstring] = 0
		}

		countPairsMap[targetSubstring]++
	}

	for step := 0; step < noOfSteps; step++ {

		tempCountPairsMap := make(map[string]int)

		for pair, count := range countPairsMap {
			ruleChar, ok := pairInsertionRulesMap[pair]

			if !ok {
				if _, ok := tempCountPairsMap[pair]; !ok {
					tempCountPairsMap[pair] = 0
				}
				tempCountPairsMap[pair] += count
			} else {
				newPair1 := string(pair[0]) + ruleChar
				newPair2 := ruleChar + string(pair[1])
				if _, ok := tempCountPairsMap[newPair1]; !ok {
					tempCountPairsMap[newPair1] = 0
				}

				if _, ok := tempCountPairsMap[newPair2]; !ok {
					tempCountPairsMap[newPair2] = 0
				}

				tempCountPairsMap[newPair1] += count
				tempCountPairsMap[newPair2] += count
			}
		}

		countPairsMap = tempCountPairsMap
	}

	var mostCommonElement *string = new(string)
	var leastCommonElement *string = new(string)
	countCharMap := make(map[string]int)

	for pair, count := range countPairsMap {
		for _, char := range pair {
			trackMostAndLeastCharacters(string(char), count, countCharMap, mostCommonElement, leastCommonElement)
			break
		}
	}

	trackMostAndLeastCharacters(string(lastCharacter), 1, countCharMap, mostCommonElement, leastCommonElement)

	return countCharMap[*mostCommonElement], countCharMap[*leastCommonElement]
}

func trackMostAndLeastCharacters(targetChar string, incrementCount int, elementScoreMap map[string]int, mostCommonElement *string, leastCommonElement *string) {
	if _, ok := elementScoreMap[targetChar]; !ok {
		elementScoreMap[targetChar] = 0
	}

	elementScoreMap[targetChar] += incrementCount

	if len(elementScoreMap) == 1 {
		*mostCommonElement = targetChar
		*leastCommonElement = targetChar
	} else {
		if elementScoreMap[*mostCommonElement] < elementScoreMap[targetChar] {
			*mostCommonElement = targetChar
		}

		if elementScoreMap[*leastCommonElement] > elementScoreMap[targetChar] {
			*leastCommonElement = targetChar
		}
	}
}

func insertCharacter(str string, char string, index int) string {
	return str[0:index] + char + str[index:]
}

func getInput() (string, map[string]string, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return "", nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	polymerTemplate := ""
	pairInsertionRulesMap := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()

		if polymerTemplate == "" {
			polymerTemplate = line
		} else if line != "" {
			elements := strings.Fields(line)
			pairInsertionRulesMap[elements[0]] = elements[2]
		}
	}

	return polymerTemplate, pairInsertionRulesMap, scanner.Err()
}
