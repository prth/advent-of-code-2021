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

type Point struct {
	x int
	y int
}

type Line struct {
	p1 Point
	p2 Point
}

func main() {
	input, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	overlapMap := make(map[string][]int)
	trackOverlappingPointsAns1 := 0
	trackOverlappingPointsAns2 := 0

	for _, line := range input {
		if line.p1.y == line.p2.y {
			// horizontal line
			for i := line.p1.x; i <= line.p2.x; i++ {
				pointStr := strconv.Itoa(i) + "," + strconv.Itoa(line.p1.y)

				if trackOverlapsForAnswer1(overlapMap, pointStr) {
					trackOverlappingPointsAns1++
				}

				if trackOverlapsForAnswer2(overlapMap, pointStr) {
					trackOverlappingPointsAns2++
				}
			}
		} else if line.p1.x == line.p2.x {
			// vertical line
			for i := line.p1.y; i <= line.p2.y; i++ {
				pointStr := strconv.Itoa(line.p1.x) + "," + strconv.Itoa(i)

				if trackOverlapsForAnswer1(overlapMap, pointStr) {
					trackOverlappingPointsAns1++
				}

				if trackOverlapsForAnswer2(overlapMap, pointStr) {
					trackOverlappingPointsAns2++
				}
			}
		} else {
			// diagonal line
			for i := line.p1.x; i <= line.p2.x; i++ {
				isDiagonalClimbing := line.p1.y < line.p2.y

				pointY := line.p1.y
				if isDiagonalClimbing {
					pointY += (i - line.p1.x)
				} else {
					pointY -= (i - line.p1.x)
				}

				pointStr := strconv.Itoa(i) + "," + strconv.Itoa(pointY)

				if trackOverlapsForAnswer2(overlapMap, pointStr) {
					trackOverlappingPointsAns2++
				}
			}
		}
	}

	answer1 := trackOverlappingPointsAns1
	log.Printf("Answer #1 :: %d", answer1)

	answer2 := trackOverlappingPointsAns2
	log.Printf("Answer #2 :: %d", answer2)
}

func trackOverlapsForAnswer1(overlapMap map[string][]int, pointStr string) bool {
	if _, ok := overlapMap[pointStr]; !ok {
		overlapMap[pointStr] = []int{0, 0}
	}

	overlapMap[pointStr][0] += 1

	return overlapMap[pointStr][0] == 2
}

func trackOverlapsForAnswer2(overlapMap map[string][]int, pointStr string) bool {
	if _, ok := overlapMap[pointStr]; !ok {
		overlapMap[pointStr] = []int{0, 0}
	}

	overlapMap[pointStr][1] += 1

	return overlapMap[pointStr][1] == 2
}

func getInput() ([]Line, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []Line

	for scanner.Scan() {
		text := scanner.Text()

		fields := strings.Fields(text)

		p1 := convertStrToPoint(fields[0])
		p2 := convertStrToPoint(fields[2])

		var line Line

		// TODO can be refactored
		if isLineDiagonal(p1, p2) {
			if p1.x < p2.x {
				line = Line{
					p1: p1,
					p2: p2,
				}
			} else {
				line = Line{
					p1: p2,
					p2: p1,
				}
			}

			input = append(input, line)
		} else if p1.y == p2.y {
			// horizontal line
			if p1.x < p2.x {
				line = Line{
					p1: p1,
					p2: p2,
				}
			} else {
				line = Line{
					p1: p2,
					p2: p1,
				}
			}

			input = append(input, line)
		} else if p1.x == p2.x {
			// vertical line
			if p1.y < p2.y {
				line = Line{
					p1: p1,
					p2: p2,
				}
			} else {
				line = Line{
					p1: p2,
					p2: p1,
				}
			}

			input = append(input, line)
		}

	}

	return input, scanner.Err()
}

func isLineDiagonal(p1 Point, p2 Point) bool {
	return math.Abs(float64(p2.x-p1.x)) == math.Abs(float64(p2.y-p1.y))
}

func convertStrToPoint(pointStr string) Point {
	points := strings.Split(pointStr, ",")
	x, _ := strconv.Atoi(points[0])
	y, _ := strconv.Atoi(points[1])
	return Point{
		x: x,
		y: y,
	}
}
