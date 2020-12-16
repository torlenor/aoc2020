package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const inputFile = "input.txt"

type geology struct {
	treeMap          [][]bool // true: tree, false: no tree, periodic boundary conditions in x direction
	currentPositionX int
	currentPositionY int
}

type slope struct {
	vx int
	vy int
}

func geologyFromFile(filePath string) (*geology, error) {
	var geo geology

	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file '%s': %s", inputFile, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var lineMap []bool
		line := scanner.Text()
		for _, c := range line {
			if c == '.' {
				lineMap = append(lineMap, false)
			} else if c == '#' {
				lineMap = append(lineMap, true)
			} else {
				log.Fatalf("Something was wrong in parsing line '%s'", line)
			}
		}
		geo.treeMap = append(geo.treeMap, lineMap)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &geo, nil
}

// move returns true if it hits a tree and false if it doesn't
// The second return value indicates if the goal was reached
func (g *geology) move(sl slope) (bool, bool) {
	finished := false
	g.currentPositionX += sl.vx
	if g.currentPositionX >= len(g.treeMap[0]) {
		g.currentPositionX -= len(g.treeMap[0])
	}
	g.currentPositionY += sl.vy
	if g.currentPositionY >= len(g.treeMap) {
		finished = true
		return false, finished
	}

	return g.treeMap[g.currentPositionY][g.currentPositionX], finished
}

func (g *geology) reset() {
	g.currentPositionX = 0
	g.currentPositionY = 0
}

func main() {
	geo, err := geologyFromFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading geology from file '%s': %s", inputFile, err)
	}

	sl := slope{
		vx: 3,
		vy: 1,
	}

	treeCount := 0
	finished := false
	for !finished {
		var tree bool
		tree, finished = geo.move(sl)
		if tree && !finished {
			treeCount++
		}
	}
	fmt.Printf("Answer to part 1: Encountered %d trees on the way down... ouch!\n", treeCount)

	fmt.Printf("\n")
	geo.reset()

	// Part 2: Different slopes
	slopes := []slope{
		{vx: 1, vy: 1},
		{vx: 3, vy: 1},
		{vx: 5, vy: 1},
		{vx: 7, vy: 1},
		{vx: 1, vy: 2},
	}

	treesHit := []int{}
	for _, sl := range slopes {
		treeCount := 0
		finished := false
		for !finished {
			var tree bool
			tree, finished = geo.move(sl)
			if tree && !finished {
				treeCount++
			}
		}
		fmt.Printf("For slope (%d, %d): Encountered %d trees on the way down... ouch!\n", sl.vx, sl.vy, treeCount)
		treesHit = append(treesHit, treeCount)
		geo.reset()
	}
	mult := 1
	for _, trees := range treesHit {
		mult *= trees
	}
	fmt.Printf("\nAnswer to part 2: %d\n", mult)
}
