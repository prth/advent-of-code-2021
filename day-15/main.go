package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFilePath = "input.txt"

type MinimumDistanceTracker struct {
	minimumDistanceMap  map[int]map[int]int
	processedNodesMap   map[int]map[int]bool
	nodeProcessingQueue [][]int
}

func main() {
	caveGrid, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	source := []int{0, 0}

	destination := []int{len(caveGrid) - 1, len(caveGrid[0]) - 1}
	minimumDistance := calculateMinimumDistance(caveGrid, source, destination)
	answer1 := minimumDistance

	fullMap := getFullMap(caveGrid)
	destinationFullMap := []int{len(fullMap) - 1, len(fullMap[0]) - 1}
	minimumDistanceFullMap := calculateMinimumDistance(fullMap, source, destinationFullMap)
	answer2 := minimumDistanceFullMap

	log.Printf("Answer #1 :: %d", answer1)
	log.Printf("Answer #2 :: %d", answer2)
}

func getFullMap(caveGrid [][]int) [][]int {
	fullMap := make([][]int, len(caveGrid))

	// repeat tiles for first row of 5X5
	for repeatIndex := 0; repeatIndex < 5; repeatIndex++ {
		for rowIndex := 0; rowIndex < len(caveGrid); rowIndex++ {
			if repeatIndex == 0 {
				fullMap[rowIndex] = make([]int, len(caveGrid[0])*5)
			}

			fullMapBaseColumnIndex := repeatIndex * len(caveGrid[0])

			for columnIndex := 0; columnIndex < len(caveGrid[0]); columnIndex++ {
				if repeatIndex > 0 {
					newColumnIndex := fullMapBaseColumnIndex + columnIndex
					prevColumnValueToIncrease := fullMap[rowIndex][newColumnIndex-len(caveGrid[0])]

					if prevColumnValueToIncrease == 9 {
						fullMap[rowIndex][newColumnIndex] = 1
					} else {
						fullMap[rowIndex][newColumnIndex] = prevColumnValueToIncrease + 1
					}
				} else {
					fullMap[rowIndex][columnIndex] = caveGrid[rowIndex][columnIndex]
				}
			}
		}
	}

	// repeat for remaining rows of full map
	for repeatIndex := len(caveGrid); repeatIndex < len(caveGrid)*5; repeatIndex++ {
		fullMap = append(fullMap, make([]int, len(fullMap[0])))

		newRowIndex := len(fullMap) - 1
		prevRowIndex := newRowIndex - len(caveGrid)

		for columnIndex := 0; columnIndex < len(fullMap[0]); columnIndex++ {
			prevColumnValueToIncrease := fullMap[prevRowIndex][columnIndex]

			if prevColumnValueToIncrease == 9 {
				fullMap[newRowIndex][columnIndex] = 1
			} else {
				fullMap[newRowIndex][columnIndex] = prevColumnValueToIncrease + 1
			}
		}
	}

	return fullMap
}

func calculateMinimumDistance(caveGrid [][]int, source []int, destination []int) int {
	var minimumDistanceTracker MinimumDistanceTracker
	minimumDistanceTracker.minimumDistanceMap = make(map[int]map[int]int)
	minimumDistanceTracker.processedNodesMap = map[int]map[int]bool{}
	minimumDistanceTracker.nodeProcessingQueue = make([][]int, 0)

	minimumDistanceTracker.setMinimumDistance(source[0], source[1], 0)
	minimumDistanceTracker.addNodeToProcessNext(source[0], source[1])

	for node, ok := minimumDistanceTracker.getNextNodeToProcess(); ok; node, ok = minimumDistanceTracker.getNextNodeToProcess() {
		if minimumDistanceTracker.isNodeProcessed(node[0], node[1]) {
			continue
		}

		if node[0] == destination[0] && node[1] == destination[1] {
			break
		}

		minimumDistanceTracker.markNodeProcessed(node[0], node[1])
		baseMinDistanceTrackedOfNode, _ := minimumDistanceTracker.getMinimumDistanceTracked(node[0], node[1])

		neighbours := getCaveNeigbours(caveGrid, node)

		for _, neighbourNode := range neighbours {
			if minimumDistanceTracker.isNodeProcessed(neighbourNode[0], neighbourNode[1]) {
				continue
			}

			distanceTrackedSoFar, isTracked := minimumDistanceTracker.getMinimumDistanceTracked(neighbourNode[0], neighbourNode[1])
			checkNewMinDistance := baseMinDistanceTrackedOfNode + caveGrid[neighbourNode[0]][neighbourNode[1]]

			if !isTracked {
				minimumDistanceTracker.setMinimumDistance(neighbourNode[0], neighbourNode[1], checkNewMinDistance)
				minimumDistanceTracker.addNodeToProcessNext(neighbourNode[0], neighbourNode[1])
			} else {
				if distanceTrackedSoFar > checkNewMinDistance {
					minimumDistanceTracker.setMinimumDistance(neighbourNode[0], neighbourNode[1], checkNewMinDistance)
					minimumDistanceTracker.addNodeToProcessNext(neighbourNode[0], neighbourNode[1])
				}
			}
		}
	}

	minDistanceToDestination, _ := minimumDistanceTracker.getMinimumDistanceTracked(destination[0], destination[1])

	return minDistanceToDestination
}

func getCaveNeigbours(caveGrid [][]int, node []int) [][]int {
	neighbours := make([][]int, 0)

	// above
	if node[0] > 0 {
		neighbours = append(neighbours, []int{node[0] - 1, node[1]})
	}

	// below
	if node[0] < len(caveGrid)-1 {
		neighbours = append(neighbours, []int{node[0] + 1, node[1]})
	}

	// left
	if node[1] > 0 {
		neighbours = append(neighbours, []int{node[0], node[1] - 1})
	}

	// right
	if node[1] < len(caveGrid[0])-1 {
		neighbours = append(neighbours, []int{node[0], node[1] + 1})
	}

	return neighbours
}

func (tracker *MinimumDistanceTracker) setMinimumDistance(rowIndex int, columnIndex int, distanceValue int) {
	if _, ok := tracker.minimumDistanceMap[rowIndex]; !ok {
		tracker.minimumDistanceMap[rowIndex] = make(map[int]int)
	}

	tracker.minimumDistanceMap[rowIndex][columnIndex] = distanceValue
}

func (tracker *MinimumDistanceTracker) getMinimumDistanceTracked(rowIndex int, columnIndex int) (int, bool) {
	if _, ok := tracker.minimumDistanceMap[rowIndex]; !ok {
		return 0, false
	}

	if _, ok := tracker.minimumDistanceMap[rowIndex][columnIndex]; !ok {
		return 0, false
	}

	return tracker.minimumDistanceMap[rowIndex][columnIndex], true
}

func (tracker *MinimumDistanceTracker) markNodeProcessed(rowIndex int, columnIndex int) {
	if _, ok := tracker.processedNodesMap[rowIndex]; !ok {
		tracker.processedNodesMap[rowIndex] = make(map[int]bool)
	}

	tracker.processedNodesMap[rowIndex][columnIndex] = true
}

func (tracker *MinimumDistanceTracker) isNodeProcessed(rowIndex int, columnIndex int) bool {
	if _, ok := tracker.processedNodesMap[rowIndex]; !ok {
		return false
	}

	if _, ok := tracker.processedNodesMap[rowIndex][columnIndex]; !ok {
		return false
	}

	return true
}

func (tracker *MinimumDistanceTracker) addNodeToProcessNext(rowIndex int, columnIndex int) {
	tracker.nodeProcessingQueue = append(tracker.nodeProcessingQueue, []int{rowIndex, columnIndex})
}

// TODO speed can be improved here using min heap
func (tracker *MinimumDistanceTracker) getNextNodeToProcess() ([]int, bool) {
	if len(tracker.nodeProcessingQueue) == 0 {
		return []int{}, false
	}

	minNodeIndex := -1
	var minNode []int

	for index, nodeAvailable := range tracker.nodeProcessingQueue {
		if minNodeIndex == -1 {
			minNode = nodeAvailable
			minNodeIndex = index
		} else {
			if tracker.minimumDistanceMap[minNode[0]][minNode[1]] > tracker.minimumDistanceMap[nodeAvailable[0]][nodeAvailable[1]] {
				minNodeIndex = index
				minNode = nodeAvailable
			}
		}
	}

	newProcessingQueue := make([][]int, 0)

	if minNodeIndex > 0 {
		newProcessingQueue = append(newProcessingQueue, tracker.nodeProcessingQueue[0:minNodeIndex]...)
	}

	if minNodeIndex < len(tracker.nodeProcessingQueue)-1 {
		newProcessingQueue = append(newProcessingQueue, tracker.nodeProcessingQueue[minNodeIndex+1:]...)
	}

	tracker.nodeProcessingQueue = newProcessingQueue

	return minNode, true
}

func isDestination(caveGrid [][]int, trackLocation []int) bool {
	return trackLocation[0] == len(caveGrid)-1 && trackLocation[1] == len(caveGrid[0])-1
}

func getInput() ([][]int, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input [][]int

	for scanner.Scan() {
		line := scanner.Text()

		var row []int

		for _, numStr := range strings.Split(line, "") {
			num, _ := strconv.Atoi(numStr)
			row = append(row, num)
		}

		input = append(input, row)
	}

	return input, scanner.Err()
}
