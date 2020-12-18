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

const inputFile = "input_example2.txt"

var reMatchInstructionAndValue = regexp.MustCompile(`(\w*) ([\-+0-9]*)`)

func main() {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file '%s': %s", inputFile, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	numbers := []int{0}
	for scanner.Scan() {
		input := scanner.Text()
		value, err := strconv.Atoi(input)
		if err != nil {
			log.Fatalf("Could not convert '%s' to number: %s", input, err)
		}
		numbers = append(numbers, value)
	}

	sort.Ints(numbers)
	numbers = append(numbers, numbers[len(numbers)-1]+3)
	for _, n := range numbers {
		fmt.Printf("%d\n", n)
	}

	fmt.Printf("built-in joltage adapter: %d\n", numbers[len(numbers)-1])

	diffCount1 := 0
	diffCount3 := 0
	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]
		if diff == 1 {
			diffCount1++
		} else if diff == 3 {
			diffCount3++
		}
	}

	fmt.Printf("Answer to part 1: 1-jolt differences = %d, 3-jolt differences = %d, multiplied = %d\n", diffCount1, diffCount3, diffCount1*diffCount3)

	diffs := []int{}
	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]
		diffs = append(diffs, diff)
	}

	fmt.Printf("diffs: %v", diffs)

	combinations := 1
	cnt := 0
	for i := 0; i < len(diffs)-1; i++ {
		if diffs[i] == 1 && diffs[i+1] == 3 {
			combinations += cnt
			if cnt == 1 {
				combinations *= 2
			} else if cnt == 2 {
				combinations *= 4
			} else if cnt == 3 {
				combinations *= 7
			}
			cnt = 0
			continue
		} else {
			cnt++
		}
	}

	// combinations := 1
	// for i := 0; i < len(numbers); i++ {
	// 	localCombinations := 0
	// 	for j := i + 1; j < len(numbers); j++ {
	// 		if (numbers[j] - numbers[i]) > 3 {
	// 			fmt.Printf("%d - %d = %d\n", numbers[j], numbers[i], numbers[j]-numbers[i])
	// 			i = j
	// 			break
	// 		}
	// 		fmt.Printf("Allowed to remove %d\n", numbers[j])
	// 		localCombinations++
	// 	}
	// 	if localCombinations > 0 {
	// 		combinations = combinations * localCombinations
	// 	}
	// }

	fmt.Printf("Answer to part 2: %d\n", combinations)
}
