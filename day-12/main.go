package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"
)

const inputFilePath = "input.txt"

type Graph struct {
	nodeAdjacentsMap map[string][]string
}

func main() {
	graph, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	answer1 := getNumberOfPaths(graph, "start", "end")
	answer2 := answer1 + getNumberOfPathWithTwiceSmallCave(graph, "start", "end")

	log.Printf("Answer #1 :: %d", answer1)
	log.Printf("Answer #2 :: %d", answer2)
}

func getNumberOfPaths(graph Graph, startNode string, endNode string) int {
	visitedCaves := make(map[string]bool)

	var pathCounter *int
	pathCounter = new(int)
	*pathCounter = 0

	countPaths(graph, "start", "end", visitedCaves, pathCounter)

	return *pathCounter
}

func countPaths(graph Graph, sourceNode string, endNode string, visitedCaves map[string]bool, pathCounter *int) {

	if sourceNode == endNode {
		*pathCounter++
		return
	}

	visitedCaves[sourceNode] = true

	for _, nextCave := range graph.nodeAdjacentsMap[sourceNode] {
		if isCaveSmall(nextCave) {
			if isVisited, ok := visitedCaves[nextCave]; !ok || !isVisited {
				countPaths(graph, nextCave, endNode, visitedCaves, pathCounter)
			}
		} else {
			countPaths(graph, nextCave, endNode, visitedCaves, pathCounter)
		}
	}

	visitedCaves[sourceNode] = false
}

func getNumberOfPathWithTwiceSmallCave(graph Graph, startNode string, endNode string) int {
	var pathCounter *int
	pathCounter = new(int)
	*pathCounter = 0

	for vertex := range graph.nodeAdjacentsMap {
		if isCaveSmall(vertex) && vertex != "start" && vertex != "end" {
			visitedCaves := make(map[string]int)

			countPathsWithTwiceSmallCave(graph, "start", "end", visitedCaves, vertex, pathCounter)
		}
	}

	return *pathCounter
}

func countPathsWithTwiceSmallCave(graph Graph, sourceNode string, endNode string, visitedCaves map[string]int, smallCaveMultipleVisit string, pathCounter *int) {

	if sourceNode == endNode {
		for node, value := range visitedCaves {
			if isCaveSmall(node) && value == 2 {
				*pathCounter++
			}
		}

		return
	}

	visitedCaves[sourceNode]++

	for _, nextCave := range graph.nodeAdjacentsMap[sourceNode] {
		if isCaveSmall(nextCave) {
			visitCount, ok := visitedCaves[nextCave]

			if !ok {
				countPathsWithTwiceSmallCave(graph, nextCave, endNode, visitedCaves, smallCaveMultipleVisit, pathCounter)
			} else if smallCaveMultipleVisit == nextCave && visitCount < 2 {
				countPathsWithTwiceSmallCave(graph, nextCave, endNode, visitedCaves, smallCaveMultipleVisit, pathCounter)
			} else if visitCount == 0 {
				countPathsWithTwiceSmallCave(graph, nextCave, endNode, visitedCaves, smallCaveMultipleVisit, pathCounter)
			}
		} else {
			countPathsWithTwiceSmallCave(graph, nextCave, endNode, visitedCaves, smallCaveMultipleVisit, pathCounter)
		}
	}

	visitedCaves[sourceNode]--
}

func isCaveSmall(caveName string) bool {
	firstChar := getFirstCharacter(caveName)

	return unicode.IsLower(firstChar)
}

func getFirstCharacter(str string) (c rune) {
	for _, c = range str {
		return
	}

	return
}

func getInput() (Graph, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return Graph{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var graph Graph
	graph.nodeAdjacentsMap = make(map[string][]string)

	var input []string

	for scanner.Scan() {
		line := scanner.Text()

		elements := strings.Split(line, "-")

		if _, ok := graph.nodeAdjacentsMap[elements[0]]; !ok {
			graph.nodeAdjacentsMap[elements[0]] = make([]string, 0)
		}

		if _, ok := graph.nodeAdjacentsMap[elements[1]]; !ok {
			graph.nodeAdjacentsMap[elements[1]] = make([]string, 0)
		}

		graph.nodeAdjacentsMap[elements[0]] = append(graph.nodeAdjacentsMap[elements[0]], elements[1])
		graph.nodeAdjacentsMap[elements[1]] = append(graph.nodeAdjacentsMap[elements[1]], elements[0])

		input = append(input, line)
	}

	return graph, scanner.Err()
}
