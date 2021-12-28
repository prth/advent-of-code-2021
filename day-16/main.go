package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const inputFilePath = "input.txt"

type MathOperation string

const (
	sum         MathOperation = "+"
	product     MathOperation = "*"
	minimum     MathOperation = "min"
	maximum     MathOperation = "max"
	greaterThan MathOperation = ">"
	lessThan    MathOperation = "<"
	equalTo     MathOperation = "="
	exactValue  MathOperation = "exact"
	compute     MathOperation = "compute"
)

type MathExpression struct {
	operation   MathOperation
	value       int
	expressions []*MathExpression
}

type PacketDecoderProcessor struct {
	sum               int
	bitsExpression    *MathExpression
	pointer           int
	expressionPointer *MathExpression
}

func main() {
	input, err := getInput()

	if err != nil {
		log.Fatal(err)
	}

	binaryStr := covertHexadecimalToBinary(input)

	var processor PacketDecoderProcessor
	processor.computeVersionSumAndExpression(binaryStr)

	answer1 := processor.sum
	answer2 := computeExpression(processor.bitsExpression)

	log.Printf("Answer #1 :: %d", answer1)
	log.Printf("Answer #2 :: %d", answer2)
}

func (processor *PacketDecoderProcessor) computeVersionSumAndExpression(str string) {
	if len(str[processor.pointer:]) == 0 {
		return
	}

	if processor.pointer == 0 {
		// initialize the base compute math expression from the start
		processor.bitsExpression = new(MathExpression)
		*processor.bitsExpression = MathExpression{
			operation: compute,
		}
		processor.expressionPointer = new(MathExpression)
		processor.expressionPointer = processor.bitsExpression
	}

	packetVersion := processor.getPacketVersion(str)
	processor.sum += int(packetVersion)

	packetTypeId := processor.getPacketTypeId(str)
	var newExpression *MathExpression
	newExpression = new(MathExpression)

	if packetTypeId == 4 {
		literalValue := processor.getLiteralValue(str)

		*newExpression = getLiteralValueExpression(literalValue)
		processor.appendSubExpression(newExpression)
	} else {
		*newExpression = getMathOperationExpression(packetTypeId)
		processor.appendSubExpression(newExpression)

		lengthTypeId := processor.getLengthTypeId(str)

		if lengthTypeId == 0 {
			// get sub packets by total length way
			lengthOfSubPackets := processor.getLengthOfSubPackets(str)

			end := lengthOfSubPackets + processor.pointer

			for processor.pointer < end {
				processor.switchExpressionPointer(newExpression)
				processor.computeVersionSumAndExpression(str)
			}
		} else {
			// get sub packets by number of sub-packets method
			numOfSubPackets := processor.getCountOfSubPackets(str)
			numOfPacketsProcessed := 0

			for numOfPacketsProcessed < numOfSubPackets {
				processor.switchExpressionPointer(newExpression)
				processor.computeVersionSumAndExpression(str)

				numOfPacketsProcessed += 1
			}
		}
	}
}

func computeExpression(expression *MathExpression) int {
	if expression.operation == compute {
		return computeExpression(expression.expressions[0])
	}

	switch expression.operation {
	case exactValue:
		{
			return expression.value
		}
	case sum:
		{
			sum := 0
			for _, subExp := range expression.expressions {
				sum += computeExpression(subExp)
			}
			return sum
		}
	case product:
		{
			product := 1
			for _, subExp := range expression.expressions {
				product *= computeExpression(subExp)
			}
			return product
		}
	case minimum:
		{
			minValue := -1
			for _, subExp := range expression.expressions {
				value := computeExpression(subExp)
				if minValue == -1 || minValue > value {
					minValue = value
				}
			}
			return minValue
		}
	case maximum:
		{
			maxValue := -1
			for _, subExp := range expression.expressions {
				value := computeExpression(subExp)
				if maxValue == -1 || maxValue < value {
					maxValue = value
				}
			}
			return maxValue
		}
	case greaterThan:
		{
			value1 := computeExpression(expression.expressions[0])
			value2 := computeExpression(expression.expressions[1])

			if value1 > value2 {
				return 1
			}

			return 0
		}
	case lessThan:
		{
			value1 := computeExpression(expression.expressions[0])
			value2 := computeExpression(expression.expressions[1])

			if value1 < value2 {
				return 1
			}

			return 0
		}
	case equalTo:
		{
			value1 := computeExpression(expression.expressions[0])
			value2 := computeExpression(expression.expressions[1])

			if value1 == value2 {
				return 1
			}

			return 0
		}
	default:
		fmt.Println(expression.operation)
		panic("Unknown expression operation")
	}
}

func convertTypeIdToMathOperation(typeId int) MathOperation {
	if typeId == 0 {
		return sum
	} else if typeId == 1 {
		return product
	} else if typeId == 2 {
		return minimum
	} else if typeId == 3 {
		return maximum
	} else if typeId == 5 {
		return greaterThan
	} else if typeId == 6 {
		return lessThan
	} else if typeId == 7 {
		return equalTo
	}

	panic("Unknown type id " + string(typeId))
}

func getLiteralValueExpression(literalValue int) MathExpression {
	return MathExpression{
		operation: exactValue,
		value:     literalValue,
	}
}

func getMathOperationExpression(packetTypeId int) MathExpression {
	return MathExpression{
		operation: convertTypeIdToMathOperation(packetTypeId),
	}
}

func (processor *PacketDecoderProcessor) appendSubExpression(subExpression *MathExpression) {
	processor.expressionPointer.expressions = append(processor.expressionPointer.expressions, subExpression)
}

func (processor *PacketDecoderProcessor) switchExpressionPointer(newExpression *MathExpression) {
	processor.expressionPointer = newExpression
}

func (processor *PacketDecoderProcessor) getLengthTypeId(str string) int {
	lengthTypeIdBin := str[processor.pointer : processor.pointer+1]
	processor.pointer += 1

	return covertBinaryToDecimal(lengthTypeIdBin)
}

func (processor *PacketDecoderProcessor) getLengthOfSubPackets(str string) int {
	lengthTypeIdBin := str[processor.pointer : processor.pointer+15]
	processor.pointer += 15

	return covertBinaryToDecimal(lengthTypeIdBin)
}

func (processor *PacketDecoderProcessor) getCountOfSubPackets(str string) int {
	lengthTypeIdBin := str[processor.pointer : processor.pointer+11]
	processor.pointer += 11

	return covertBinaryToDecimal(lengthTypeIdBin)
}

func (processor *PacketDecoderProcessor) getLiteralValue(str string) int {
	literalBinaryRep := ""

	for processor.pointer < len(str) {
		literalPacket := str[processor.pointer : processor.pointer+5]

		literalBinaryRep += literalPacket[1:]
		processor.pointer += 5

		if literalPacket[0] == '0' {
			break
		}
	}

	return covertBinaryToDecimal(literalBinaryRep)
}

func (processor *PacketDecoderProcessor) getPacketVersion(str string) int {
	packetVersionBin := str[processor.pointer : processor.pointer+3]
	processor.pointer += 3

	return covertBinaryToDecimal(packetVersionBin)
}

func (processor *PacketDecoderProcessor) getPacketTypeId(str string) int {
	packetVersionBin := str[processor.pointer : processor.pointer+3]
	processor.pointer += 3

	return covertBinaryToDecimal(packetVersionBin)
}

func covertBinaryToDecimal(binaryStr string) int {
	i, err := strconv.ParseInt(binaryStr, 2, 64)

	if err != nil {
		panic(err)
	}

	return int(i)
}

func covertHexadecimalToBinary(rawHex string) string {
	binString := ""

	for _, c := range rawHex {
		binString += convertSingleHexaToBinaryString(string(c))
	}

	return binString
}

func convertSingleHexaToBinaryString(hex string) string {
	i, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		fmt.Printf("%s", err)
	}
	return fmt.Sprintf("%04b", i)
}

func getInput() (string, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input string

	for scanner.Scan() {
		line := scanner.Text()
		input = line
	}

	return input, scanner.Err()
}
