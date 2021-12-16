package main

import (
	"bufio"
	"log"
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

	overlapMap := make(map[string]int)
	trackOverlappingPoints := 0

	for _, line := range input {
		if line.p1.y == line.p2.y {
			// horizontal line
			for i := line.p1.x; i <= line.p2.x; i++ {
				pointStr := strconv.Itoa(i) + "," + strconv.Itoa(line.p1.y)

				if trackOverlaps(overlapMap, pointStr) {
					trackOverlappingPoints++
				}
			}
		} else {
			// vertical line
			for i := line.p1.y; i <= line.p2.y; i++ {
				pointStr := strconv.Itoa(line.p1.x) + "," + strconv.Itoa(i)

				if trackOverlaps(overlapMap, pointStr) {
					trackOverlappingPoints++
				}
			}
		}
	}

	answer1 := trackOverlappingPoints
	log.Printf("Answer #1 :: %d", answer1)
}

func trackOverlaps(overlapMap map[string]int, pointStr string) bool {
	if _, ok := overlapMap[pointStr]; !ok {
		overlapMap[pointStr] = 0
	}

	overlapMap[pointStr] += 1

	return overlapMap[pointStr] == 2
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

		if p1.y == p2.y {
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

func convertStrToPoint(pointStr string) Point {
	points := strings.Split(pointStr, ",")
	x, _ := strconv.Atoi(points[0])
	y, _ := strconv.Atoi(points[1])
	return Point{
		x: x,
		y: y,
	}
}
