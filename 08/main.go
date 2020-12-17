package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const inputFile = "input.txt"

var reMatchInstructionAndValue = regexp.MustCompile(`(\w*) ([\-+0-9]*)`)

func extractInstructionAndValue(str string) (string, int, error) {
	matches := reMatchInstructionAndValue.FindAllStringSubmatch(str, -1)
	if len(matches) != 1 {
		return "", -1, fmt.Errorf("Something wrong with matches, as it is not equal to 1 (len = %d): %v", len(matches), matches)
	}

	match := matches[0]
	if len(match) != 3 {
		return "", -1, fmt.Errorf("Something wrong with match, as it did not capture 2 groups (+1) (len = %d): %v", len(match), match)
	}

	val, err := strconv.Atoi(match[2])
	if err != nil {
		return "", -1, fmt.Errorf("Could not parse string '%s' to integer: %s", match[1], err)
	}

	return match[1], val, nil
}

func getAccNumberForPart1(instructions []string) int {
	visitedLines := make(map[int]bool)
	accumulator := 0
	for l := 0; l < len(instructions); {
		if visitedLines[l] == true {
			break
		}
		nextLine := 0
		instruction, value, err := extractInstructionAndValue(instructions[l])
		if err != nil {
			log.Fatalf("Could not parse instructions line '%s': %s", instructions[l], err)
		}
		switch instruction {
		case "acc":
			accumulator += value
			nextLine = l + 1
		case "jmp":
			nextLine = l + value
		case "nop":
			nextLine = l + 1
		}
		visitedLines[l] = true
		l = nextLine
	}
	return accumulator
}

func getAccNumberForPart2(instructions []string, replaceOccurrence int) (int, bool) {
	visitedLines := make(map[int]bool)
	accumulator := 0
	occurred := 0
	for l := 0; l < len(instructions); {
		if visitedLines[l] == true {
			return -1, false
		}
		nextLine := 0
		instruction, value, err := extractInstructionAndValue(instructions[l])
		if err != nil {
			log.Fatalf("Could not parse instructions line '%s': %s", instructions[l], err)
		}
		if instruction == "jmp" || instruction == "nop" {
			occurred++
		}
		if occurred == replaceOccurrence {
			if instruction == "jmp" {
				instruction = "nop"
			} else if instruction == "nop" {
				instruction = "jmp"
			}
		}
		switch instruction {
		case "acc":
			accumulator += value
			nextLine = l + 1
		case "jmp":
			nextLine = l + value
		case "nop":
			nextLine = l + 1
		}
		visitedLines[l] = true
		l = nextLine
	}
	return accumulator, true
}

func main() {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file '%s': %s", inputFile, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	instructions := []string{}
	for scanner.Scan() {
		input := scanner.Text()
		instructions = append(instructions, input)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	accumulator := getAccNumberForPart1(instructions)

	fmt.Printf("Answer to part 1: %d\n", accumulator)

	for i := 1; i < 1000; i++ {
		accumulatorPart2, ok := getAccNumberForPart2(instructions, i)
		if ok {
			fmt.Printf("\nAnswer to part 2: %d\n", accumulatorPart2)
		}
	}

}
