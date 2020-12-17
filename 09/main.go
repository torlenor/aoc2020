package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

// const inputFile = "input_example.txt"
// const preambleLength = 5

const inputFile = "input.txt"
const preambleLength = 25

var reMatchInstructionAndValue = regexp.MustCompile(`(\w*) ([\-+0-9]*)`)

func isSum(value int, numbers []int) bool {
	for i := 0; i < (len(numbers) - 1); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if (numbers[i] + numbers[j]) == value {
				return true
			}
		}
	}
	return false
}

func main() {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file '%s': %s", inputFile, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	numbers := []int{}
	for scanner.Scan() {
		input := scanner.Text()
		value, err := strconv.Atoi(input)
		if err != nil {
			log.Fatalf("Could not convert '%s' to number: %s", input, err)
		}
		numbers = append(numbers, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	answer := 0
	answerEntry := 0
	for i := preambleLength; i < len(numbers); i++ {
		if !isSum(numbers[i], numbers[i-preambleLength:i]) {
			answer = numbers[i]
			answerEntry = i
		}
	}

	fmt.Printf("Answer to part 1: %d\n", answer)

	series := []int{}
	for i := 0; i < len(numbers); i++ {
		if i == answerEntry {
			continue
		}
		sum := 0
		sumNumbers := []int{}
		for j := i; j < len(numbers); j++ {
			sum += numbers[j]
			sumNumbers = append(sumNumbers, numbers[j])
			if sum == answer {
				series = sumNumbers
				break
			}
		}
	}

	sort.Ints(series)

	fmt.Printf("\nAnswer to part 2: Smallest Number: %d, largest Number: %d, added: %d\n", series[0], series[len(series)-1], series[0]+series[len(series)-1])
}
