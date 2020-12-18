package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const inputFile = "input.txt"

type direction struct {
	vx int
	vy int
}

var directions = []direction{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
	{-1, -1},
	{1, 1},
	{-1, 1},
	{1, -1},
}

const defaultNumberOfToleratedPeople = 4

// Map holds the floor map and the evolution functions
type Map struct {
	floorMap        [][]int8
	toleratedPeople int
}

func mapFromString(str string) (*Map, error) {
	var m Map
	rowStrings := strings.Split(str, "\n")
	for _, row := range rowStrings {
		var rowMap []int8
		for _, c := range row {
			switch c {
			case '.':
				rowMap = append(rowMap, 0)
			case 'L':
				rowMap = append(rowMap, 1)
			case '#':
				rowMap = append(rowMap, 2)
			default:
				return nil, fmt.Errorf("Error parsing line '%s'", row)
			}
		}
		m.floorMap = append(m.floorMap, rowMap)
	}

	m.toleratedPeople = defaultNumberOfToleratedPeople

	return &m, nil
}

func (m *Map) setNumberOfToleratedPeople(n int) {
	m.toleratedPeople = n
}

func cloneFloorMap(inputMap [][]int8) [][]int8 {
	var copyiedMap [][]int8
	for _, row := range inputMap {
		copiedRow := append(row[:0:0], row...)
		copyiedMap = append(copyiedMap, copiedRow)
	}
	return copyiedMap
}

func (m *Map) step() bool {
	changed := false
	copyiedMap := cloneFloorMap(m.floorMap)
	for y := 0; y < len(m.floorMap); y++ {
		for x := 0; x < len(m.floorMap[y]); x++ {
			if m.floorMap[y][x] == 0 {
				continue
			}
			occupiedSeats := 0
			for _, d := range directions {
				xx := x + d.vx
				yy := y + d.vy
				if xx >= len(m.floorMap[y]) {
					continue
				}
				if xx < 0 {
					continue
				}
				if yy >= len(m.floorMap) {
					continue
				}
				if yy < 0 {
					continue
				}
				if m.floorMap[yy][xx] == 2 {
					occupiedSeats++
				}
			}
			if occupiedSeats == 0 && copyiedMap[y][x] != 2 {
				copyiedMap[y][x] = 2
				changed = true
			} else if occupiedSeats >= m.toleratedPeople && copyiedMap[y][x] != 1 {
				copyiedMap[y][x] = 1
				changed = true
			}
		}
	}
	m.floorMap = copyiedMap
	return changed
}

func (m *Map) stepPart2() bool {
	changed := false
	copyiedMap := cloneFloorMap(m.floorMap)
	for y := 0; y < len(m.floorMap); y++ {
		for x := 0; x < len(m.floorMap[y]); x++ {
			if m.floorMap[y][x] == 0 {
				continue
			}
			occupiedSeats := 0
			for _, d := range directions {
				xx := x + d.vx
				yy := y + d.vy
				for xx >= 0 && xx < len(m.floorMap[y]) &&
					yy >= 0 && yy < len(m.floorMap) {
					if m.floorMap[yy][xx] == 2 {
						occupiedSeats++
						break
					}
					if m.floorMap[yy][xx] == 1 {
						{
							break
						}
					}
					xx += d.vx
					yy += d.vy
				}
			}
			if occupiedSeats == 0 && copyiedMap[y][x] != 2 {
				copyiedMap[y][x] = 2
				changed = true
			} else if occupiedSeats >= m.toleratedPeople && copyiedMap[y][x] != 1 {
				copyiedMap[y][x] = 1
				changed = true
			}
		}
	}
	m.floorMap = copyiedMap
	return changed
}

func (m *Map) String() string {
	stringedMap := []string{}
	for y := 0; y < len(m.floorMap); y++ {
		var row string
		for x := 0; x < len(m.floorMap[y]); x++ {
			switch m.floorMap[y][x] {
			case 0:
				row += "."
			case 1:
				row += "L"
			case 2:
				row += "#"
			}
		}
		stringedMap = append(stringedMap, row)
	}
	return strings.Join(stringedMap, "\n")
}

func (m *Map) occupiedSeats() int {
	cnt := 0
	for y := 0; y < len(m.floorMap); y++ {
		for x := 0; x < len(m.floorMap[y]); x++ {
			if m.floorMap[y][x] == 2 {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Not able to read file '%s': %s", inputFile, err)
	}
	input := string(data)
	m, err := mapFromString(input)
	if err != nil {
		log.Fatalf("Error creating map from string: %s", m)
	}

	for m.step() {
	}

	fmt.Printf("Answer to part 1: %d\n", m.occupiedSeats())

	m, err = mapFromString(input)
	if err != nil {
		log.Fatalf("Error creating map from string: %s", m)
	}
	m.setNumberOfToleratedPeople(5)

	for m.stepPart2() {
	}

	fmt.Printf("Answer to part 2: %d\n", m.occupiedSeats())
}
