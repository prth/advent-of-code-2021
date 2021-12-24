package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFilePath = "input.txt"

const (
	horizontal = "horizontal"
	veritical  = "veritical"
)

type FoldingInstruction struct {
	direction string
	position  int
}

type TransparentPaper struct {
	rowCount    int
	columnCount int
	dotsMap     map[string]struct{}
}

func main() {
	transparentPaper, foldingInstructions, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	answer1 := 0

	for index, instruction := range foldingInstructions {
		transparentPaper.foldPaper(instruction)

		if index == 0 {
			answer1 = len(transparentPaper.dotsMap)
		}
	}

	log.Printf("Answer #1 :: %d", answer1)
	log.Printf("Answer #2")
	transparentPaper.printPaper()
}

func (paper *TransparentPaper) printPaper() {
	for i := 0; i < paper.rowCount; i++ {
		for j := 0; j < paper.columnCount; j++ {
			_, ok := paper.dotsMap[getPositionString(j, i)]

			if ok {
				fmt.Print(" # ")
			} else {
				fmt.Print(" . ")
			}
		}
		fmt.Println()
	}
}

func (paper *TransparentPaper) foldPaper(instruction FoldingInstruction) {
	mergedDotsMap := make(map[string]struct{})

	firstSideIndex := instruction.position - 1
	secondSideIndex := firstSideIndex + 2

	if instruction.direction == horizontal {
		secondSideExtraRows := (paper.rowCount - secondSideIndex - 1) - (firstSideIndex + 1)
		if secondSideExtraRows < 0 {
			secondSideExtraRows = 0
		}

		for firstSideIndex >= 0 && secondSideIndex <= paper.rowCount-1 {
			for i := 0; i < paper.columnCount; i++ {
				mergedFinalMapPosition := getPositionString(i, firstSideIndex+secondSideExtraRows)

				firstSidePosition := getPositionString(i, firstSideIndex)
				_, isDotOnFirstSide := paper.dotsMap[firstSidePosition]

				if isDotOnFirstSide {
					mergedDotsMap[mergedFinalMapPosition] = struct{}{}
					continue
				}

				secondSidePosition := getPositionString(i, secondSideIndex)
				_, isDotOnSecondSide := paper.dotsMap[secondSidePosition]

				if isDotOnSecondSide {
					mergedDotsMap[mergedFinalMapPosition] = struct{}{}
				}
			}

			firstSideIndex--
			secondSideIndex++
		}

		if secondSideIndex < paper.rowCount-1 {
			for rowIndex := secondSideIndex; rowIndex < paper.rowCount; rowIndex++ {
				for columnIndex := 0; columnIndex < paper.columnCount; columnIndex++ {
					position := getPositionString(rowIndex, columnIndex)
					_, isDotOnSecondSide := paper.dotsMap[position]

					if isDotOnSecondSide {
						mergedRowIndex := paper.rowCount - 1 - secondSideIndex
						mergedFinalMapPosition := getPositionString(mergedRowIndex, columnIndex)
						mergedDotsMap[mergedFinalMapPosition] = struct{}{}
					}
				}
			}
		}

		paper.dotsMap = mergedDotsMap
		paper.rowCount = instruction.position + secondSideExtraRows
	} else {
		secondSideExtraColumns := (paper.columnCount - secondSideIndex - 1) - (firstSideIndex + 1)
		if secondSideExtraColumns < 0 {
			secondSideExtraColumns = 0
		}

		for firstSideIndex >= 0 && secondSideIndex <= paper.columnCount-1 {
			for i := 0; i < paper.rowCount; i++ {
				mergedFinalMapPosition := getPositionString(firstSideIndex+secondSideExtraColumns, i)

				firstSidePosition := getPositionString(firstSideIndex, i)
				_, isDotOnFirstSide := paper.dotsMap[firstSidePosition]

				if isDotOnFirstSide {
					mergedDotsMap[mergedFinalMapPosition] = struct{}{}
					continue
				}

				secondSidePosition := getPositionString(secondSideIndex, i)
				_, isDotOnSecondSide := paper.dotsMap[secondSidePosition]

				if isDotOnSecondSide {
					mergedDotsMap[mergedFinalMapPosition] = struct{}{}
				}
			}

			firstSideIndex--
			secondSideIndex++
		}

		if secondSideIndex < paper.columnCount-1 {
			for columnIndex := secondSideIndex; columnIndex < paper.columnCount; columnIndex++ {
				for rowIndex := 0; rowIndex < paper.rowCount; rowIndex++ {
					position := getPositionString(rowIndex, columnIndex)
					_, isDotOnSecondSide := paper.dotsMap[position]

					if isDotOnSecondSide {
						mergedColumnIndex := paper.columnCount - 1 - secondSideIndex
						mergedFinalMapPosition := getPositionString(mergedColumnIndex, rowIndex)
						mergedDotsMap[mergedFinalMapPosition] = struct{}{}
					}
				}
			}
		}

		paper.dotsMap = mergedDotsMap
		paper.columnCount = instruction.position + secondSideExtraColumns

	}
}

func getPositionString(rowPosition int, columnPosition int) string {
	return strconv.Itoa(rowPosition) + "," + strconv.Itoa(columnPosition)
}

func getInput() (TransparentPaper, []FoldingInstruction, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return TransparentPaper{}, []FoldingInstruction{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var transparentPaper TransparentPaper
	transparentPaper.dotsMap = make(map[string]struct{})

	keepSettingDots := true

	foldingInstructions := make([]FoldingInstruction, 0)
	instructionDetailStartIndex := len("fold along ")

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			keepSettingDots = false
		} else if keepSettingDots {
			transparentPaper.dotsMap[line] = struct{}{}

			coordinates := strings.Split(line, ",")
			columnPosition, _ := strconv.Atoi(coordinates[0])
			rowPosition, _ := strconv.Atoi(coordinates[1])

			if transparentPaper.columnCount < columnPosition+1 {
				transparentPaper.columnCount = columnPosition + 1
			}

			if transparentPaper.rowCount < rowPosition+1 {
				transparentPaper.rowCount = rowPosition + 1
			}
		} else {
			instructionDetail := strings.Split(line[instructionDetailStartIndex:], "=")

			foldingInstruction := FoldingInstruction{}
			foldingInstruction.position, _ = strconv.Atoi(instructionDetail[1])

			if instructionDetail[0] == "x" {
				foldingInstruction.direction = veritical
			} else {
				foldingInstruction.direction = horizontal
			}

			foldingInstructions = append(foldingInstructions, foldingInstruction)
		}
	}

	return transparentPaper, foldingInstructions, scanner.Err()
}
