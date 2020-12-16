package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

const inputFile = "input.txt"

func main() {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Not able to read file '%s': %s", inputFile, err)
	}

	input := string(data)

	re := regexp.MustCompile(`(.{1,})\n`)
	input = re.ReplaceAllString(input, "${1} ")
	splittedInput := strings.Split(input, "\n")

	totalNumberOfAnswers := 0
	for _, group := range splittedInput {
		uniqueAnswers := make(map[rune]bool)
		trimmed := strings.ReplaceAll(group, " ", "")
		for _, c := range trimmed {
			uniqueAnswers[c] = true
		}
		totalNumberOfAnswers += len(uniqueAnswers)
	}

	fmt.Printf("Answer to part 1: %d\n", totalNumberOfAnswers)

	numberOfCommonAnswers := 0
	for _, group := range splittedInput {
		trimmed := strings.TrimSpace(group)
		peoplesAnswers := strings.Split(trimmed, " ")
		answersCount := make(map[rune]int)
		for _, p := range peoplesAnswers {
			for _, c := range p {
				answersCount[c]++
			}
		}
		commonAnswers := 0
		for _, a := range answersCount {
			if a == len(peoplesAnswers) {
				commonAnswers++
			}
		}
		numberOfCommonAnswers += commonAnswers
	}

	fmt.Printf("\nAnswer to part 2: %d\n", numberOfCommonAnswers)
}
