package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const inputFile = "input.txt"

// findNumberPair finds the two numbers that sum up to 2020
func findNumberPair(expenses []int) (ok bool, firstNumber int, secondNumber int) {
	firstNumberCnt := 0
	for firstNumberCnt < len(expenses)-1 {
		secondNumberCnt := firstNumberCnt + 1
		for secondNumberCnt < len(expenses) {
			if (expenses[firstNumberCnt] + expenses[secondNumberCnt]) == 2020 {
				return true, expenses[firstNumberCnt], expenses[secondNumberCnt]
			}
			secondNumberCnt++
		}
		firstNumberCnt++
	}

	return false, 0, 0
}

// findThreeNumbers finds the three numbers that sum up to 2020
func findThreeNumbers(expenses []int) (ok bool, firstNumber int, secondNumber int, thirdNumber int) {
	firstNumberCnt := 0
	for firstNumberCnt < len(expenses)-2 {
		secondNumberCnt := firstNumberCnt + 1
		for secondNumberCnt < len(expenses)-1 {
			thirdNumberCnt := secondNumberCnt + 1
			for thirdNumberCnt < len(expenses) {
				if (expenses[firstNumberCnt] + expenses[secondNumberCnt] + expenses[thirdNumberCnt]) == 2020 {
					return true, expenses[firstNumberCnt], expenses[secondNumberCnt], expenses[thirdNumberCnt]
				}
				thirdNumberCnt++
			}
			secondNumberCnt++
		}
		firstNumberCnt++
	}

	return false, 0, 0, 0
}

func main() {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file '%s': %s", inputFile, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	expenses := []int{}
	for scanner.Scan() {
		input := scanner.Text()
		val, err := strconv.Atoi(input)
		if err != nil {
			log.Fatalf("Error reading input line '%s': %s", input, err)
		}
		expenses = append(expenses, val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	ok, firstNumber, secondNumber := findNumberPair(expenses)
	if ok {
		fmt.Printf("The solution to the first part of the puzzle is %d * %d = %d\n", firstNumber, secondNumber, firstNumber*secondNumber)
	} else {
		fmt.Println("Could not find two numbers that sum up to 2020 :(")
	}

	ok, firstNumber, secondNumber, thirdNumber := findThreeNumbers(expenses)
	if ok {
		fmt.Printf("The solution to the second part of the puzzle is %d * %d * %d = %d\n", firstNumber, secondNumber, thirdNumber,
			firstNumber*secondNumber*thirdNumber)
	} else {
		fmt.Println("Could not find three numbers that sum up to 2020 :(")
	}

}
