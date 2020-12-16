package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

const inputFile = "input.txt"

func main() {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file '%s': %s", inputFile, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	seats := []string{}
	for scanner.Scan() {
		input := scanner.Text()
		seats = append(seats, input)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	seatIDs := []int{}
	highestSeatID := -1
	for _, seat := range seats {
		rowUp := 127
		rowDown := 0
		for cnt := 0; cnt < 6; cnt++ {
			c := seat[cnt]
			if c == 'F' {
				rowUp = rowDown + (rowUp-rowDown)/2
			} else {
				rowDown = rowUp - int((float32(rowUp)-float32(rowDown))/float32(2)+0.5) + 1
			}
		}
		row := 0
		if seat[6] == 'F' {
			row = rowDown
		} else {
			row = rowUp
		}

		colUp := 7
		colDown := 0
		for cnt := 7; cnt < 10; cnt++ {
			c := seat[cnt]
			if c == 'L' {
				colUp = colDown + (colUp-colDown)/2
			} else {
				colDown = colUp - int((float32(colUp)-float32(colDown))/float32(2)+0.5) + 1
			}
		}
		col := 0
		if seat[6] == 'L' {
			col = colDown
		} else {
			col = colUp
		}

		seatID := row*8 + col
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
		seatIDs = append(seatIDs, seatID)
	}

	fmt.Printf("Answer to part 1: The highest seat ID is %d\n", highestSeatID)

	mySeatID := -1

	sort.Ints(seatIDs)
	for i := 1; i < len(seatIDs); i++ {
		if (seatIDs[i] - seatIDs[i-1]) > 1 {
			mySeatID = seatIDs[i-1] + 1
		}

	}

	fmt.Printf("\nAnswer to part 2: My seat ID is %d\n", mySeatID)
}
