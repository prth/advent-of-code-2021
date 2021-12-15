package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFilePath = "input.txt"

func main() {
	drawNumbers, boards, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	boardsNumMap := make(map[string]map[int][]int)

	for boardIndex, board := range boards {
		for rowIndex, row := range board {
			for columnIndex, column := range row {

				if _, ok := boardsNumMap[column]; !ok {
					boardsNumMap[column] = make(map[int][]int)
				}

				boardsNumMap[column][boardIndex] = []int{rowIndex, columnIndex, 0}
			}
		}
	}

	boardTrackRows := make(map[int]map[int]int)
	boardTrackColumns := make(map[int]map[int]int)

	var firstWinningBoard int
	var lastWinningBoard int

	winningBoardMap := make(map[int]string)

	for _, drawNum := range drawNumbers {
		if numMap, ok := boardsNumMap[drawNum]; ok {

			for boardIndex := 0; boardIndex < len(boards); boardIndex++ {

				if _, ok := winningBoardMap[boardIndex]; ok {
					continue
				}

				if boardPosition, ok := numMap[boardIndex]; ok {
					boardPosition[2] = 1

					if _, ok := boardTrackRows[boardIndex]; !ok {
						boardTrackRows[boardIndex] = make(map[int]int)
						boardTrackColumns[boardIndex] = make(map[int]int)
					}

					if _, ok := boardTrackRows[boardPosition[0]]; !ok {
						boardTrackRows[boardIndex][boardPosition[0]] = 0
					}

					if _, ok := boardTrackColumns[boardPosition[1]]; !ok {
						boardTrackColumns[boardIndex][boardPosition[1]] = 0
					}

					boardTrackRows[boardIndex][boardPosition[0]]++
					boardTrackColumns[boardIndex][boardPosition[1]]++

					if boardTrackRows[boardIndex][boardPosition[0]] == len(boards[0][0]) {
						winningBoardMap[boardIndex] = drawNum

						if len(winningBoardMap) == 1 {
							firstWinningBoard = boardIndex
						} else if len(winningBoardMap) == len(boards) {
							lastWinningBoard = boardIndex
						}
					}

					if boardTrackColumns[boardIndex][boardPosition[1]] == len(boards[0][0]) {
						winningBoardMap[boardIndex] = drawNum

						if len(winningBoardMap) == 0 {
							firstWinningBoard = boardIndex
						} else if len(winningBoardMap) == len(boards) {
							lastWinningBoard = boardIndex
						}
					}
				}
			}
		}
	}

	firstWinSumOfUnmarkedNumbers := getSumOfUnmarkedNumbers(boards, firstWinningBoard, boardsNumMap)
	firstWinLastNumOfSelection, _ := strconv.Atoi(winningBoardMap[firstWinningBoard])
	answer1 := firstWinSumOfUnmarkedNumbers * firstWinLastNumOfSelection

	lastWinSumOfUnmarkedNumbers := getSumOfUnmarkedNumbers(boards, lastWinningBoard, boardsNumMap)
	lastWinLastNumOfSelection, _ := strconv.Atoi(winningBoardMap[lastWinningBoard])
	answer2 := lastWinSumOfUnmarkedNumbers * lastWinLastNumOfSelection

	log.Printf("Answer #1 :: %d", answer1)
	log.Printf("Answer #2 :: %d", answer2)
}

func getSumOfUnmarkedNumbers(boards [][][]string, winningBoardIndex int, boardsNumMap map[string]map[int][]int) int {
	sumOfUnmarkedNumbers := 0

	for _, row := range boards[winningBoardIndex] {
		for _, column := range row {
			if boardsNum, ok := boardsNumMap[column][winningBoardIndex]; !ok || boardsNum[2] == 1 {
				continue
			}

			numValue, _ := strconv.Atoi(column)
			sumOfUnmarkedNumbers += numValue
		}
	}

	return sumOfUnmarkedNumbers
}

func getInput() ([]string, [][][]string, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var drawNumbers []string
	boards := make([][][]string, 0)

	startNewBoard := false
	trackBoardRowIndex := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			startNewBoard = true
			trackBoardRowIndex = 0
			continue
		}

		if drawNumbers == nil {
			drawNumbers = strings.Split(line, ",")
			continue
		}

		boardRow := strings.Fields(line)

		if startNewBoard {
			board := make([][]string, len(boardRow))
			board[0] = boardRow

			boards = append(boards, board)
			startNewBoard = false
		} else {
			trackBoardRowIndex++

			board := boards[len(boards)-1]
			board[trackBoardRowIndex] = boardRow
		}
	}

	return drawNumbers, boards, scanner.Err()
}
