package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFile = "input_example.txt"

func divmod(numerator, denominator int) (quotient, remainder int) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file '%s': %s", inputFile, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	lines := []string{}
	for scanner.Scan() {
		input := scanner.Text()
		lines = append(lines, input)
	}

	earliestDepartureTime, err := strconv.Atoi(lines[0])
	busIDsStringCleaned := strings.ReplaceAll(lines[1], ",x", "")
	busIDsStr := strings.Split(busIDsStringCleaned, ",")
	busIDs := []int{}
	for _, idStr := range busIDsStr {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatalf("Error converting '%s' to int: %s", idStr, err)
		}
		busIDs = append(busIDs, id)
	}

	waitTime := 100000
	bestID := 0
	for _, id := range busIDs {
		for i := 0; ; i++ {
			if (id*i) >= earliestDepartureTime && id*i < (earliestDepartureTime+waitTime) {
				waitTime = id*i - earliestDepartureTime
				bestID = id
			}
			if (id * i) >= earliestDepartureTime+waitTime {
				break
			}
		}
	}

	fmt.Printf("Answer to part 1: Wait time %d, bus ID %d, multiplied = %d\n", waitTime, bestID, waitTime*bestID)

	// TODO
	// Chinese Reminder Theorem?

	fmt.Printf("Answer to part 2: %d\n", -1)
}
