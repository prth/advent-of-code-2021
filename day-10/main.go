package main

import (
	"bufio"
	"log"
	"os"
	"sort"
)

const inputFilePath string = "input.txt"

type CharacterStack []rune

var bracketsOpenCloseMap = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var illegalCharacterScopeMap = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var incompleteCharacterScopeMap = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func main() {
	input, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	illegalCharacterScoreSum := 0
	var incompleteLineScores []int

	for _, line := range input {
		var expectedClosingCharacterStack CharacterStack
		var illegalCharacterScore int

		for _, character := range line {
			closingCharacter, ok := bracketsOpenCloseMap[character]

			if ok {
				expectedClosingCharacterStack.Push(closingCharacter)
			} else {
				expectedClosingCharacter, isEmpty := expectedClosingCharacterStack.Pop()

				if !isEmpty || expectedClosingCharacter != character {
					illegalCharacterScore, _ = illegalCharacterScopeMap[character]
					break
				}
			}
		}

		if illegalCharacterScore > 0 {
			illegalCharacterScoreSum += illegalCharacterScore
		} else if !expectedClosingCharacterStack.IsEmpty() {
			// so this is that incomplete line!
			incompleteLineScore := 0

			for !expectedClosingCharacterStack.IsEmpty() {
				closingCharacterNeeded, _ := expectedClosingCharacterStack.Pop()
				incompleteCharacterScore, _ := incompleteCharacterScopeMap[closingCharacterNeeded]

				incompleteLineScore = (5 * incompleteLineScore) + incompleteCharacterScore
			}

			incompleteLineScores = append(incompleteLineScores, incompleteLineScore)
		}
	}

	sort.Ints(incompleteLineScores)

	answer1 := illegalCharacterScoreSum
	answer2 := incompleteLineScores[len(incompleteLineScores)/2]

	log.Printf("Answer #1 :: %d", answer1)
	log.Printf("Answer #2 :: %d", answer2)
}

func (stack *CharacterStack) IsEmpty() bool {
	return len(*stack) == 0
}

func (stack *CharacterStack) Push(c rune) {
	*stack = append(*stack, c)
}

func (stack *CharacterStack) Pop() (rune, bool) {

	if stack.IsEmpty() {
		return rune(0), false
	}

	indexToPop := len(*stack) - 1
	character := (*stack)[indexToPop]
	*stack = (*stack)[:indexToPop]

	return character, true
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
