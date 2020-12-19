package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const inputFile = "input.txt"

func trimLeftChar(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return s[:0]
}

func firstRune(str string) (r rune) {
	for _, r = range str {
		return
	}
	return
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type direction int

func (d direction) String() string {
	switch d {
	case 0:
		return "north"
	case 1:
		return "west"
	case 2:
		return "south"
	case 3:
		return "east"
	}
	return "UNKNOWN"
}

// Ship models the ship
type Ship struct {
	x                int
	y                int
	currentDirection direction // 0 north, 1 west, 2 south, 3 east
}

func newShip(initialDirection direction) Ship {
	ship := Ship{
		currentDirection: initialDirection,
	}
	return ship
}

func (s *Ship) move(instruction string) {
	command := firstRune(instruction)
	valueStr := trimLeftChar(instruction)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Fatalf("Error converting '%s' to int: %d", valueStr, err)
	}
	switch command {
	case 'N':
		s.y += value
	case 'S':
		s.y -= value
	case 'E':
		s.x += value
	case 'W':
		s.x -= value
	case 'R':
		s.currentDirection -= direction((value / 90))
		for s.currentDirection < 0 {
			s.currentDirection += 4
		}
	case 'L':
		s.currentDirection += direction((value / 90))
		for s.currentDirection > 3 {
			s.currentDirection -= 4
		}
	case 'F':
		switch s.currentDirection {
		case 0:
			s.y += value
		case 1:
			s.x -= value
		case 2:
			s.y -= value
		case 3:
			s.x += value
		}
	default:
		log.Fatalf("Unknown command: %s", string(command))
	}
}

func (s *Ship) getCurrentPosition() (int, int) {
	return s.x, s.y
}

func (s *Ship) getManhattanDistance() int {
	return absInt(s.x) + absInt(s.y)
}

// ShipPart2 models the ship with Waypoint for part 2
type ShipPart2 struct {
	x int
	y int

	waypointX int
	waypointY int
}

func newShipPart2() ShipPart2 {
	ship := ShipPart2{
		waypointX: 10,
		waypointY: 1,
	}
	return ship
}

func (s *ShipPart2) move(instruction string) {
	command := firstRune(instruction)
	valueStr := trimLeftChar(instruction)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Fatalf("Error converting '%s' to int: %d", valueStr, err)
	}
	switch command {
	case 'N':
		s.waypointY += value
	case 'S':
		s.waypointY -= value
	case 'E':
		s.waypointX += value
	case 'W':
		s.waypointX -= value
	case 'R':
		newDirection := direction((value / 90))
		x1 := 0
		y1 := 0
		switch newDirection {
		case 1: // 90
			x1 = s.waypointY
			y1 = -s.waypointX
		case 2: // 180
			x1 = -s.waypointX
			y1 = -s.waypointY
		case 3: // 270
			x1 = -s.waypointY
			y1 = s.waypointX
		default:
			log.Fatalf("Angle out of bounce: %d", value)
		}
		s.waypointX = x1
		s.waypointY = y1
	case 'L':
		newDirection := direction((value / 90))
		x1 := 0
		y1 := 0
		switch newDirection {
		case 1: // -90
			x1 = -s.waypointY
			y1 = s.waypointX
		case 2: // -180
			x1 = -s.waypointX
			y1 = -s.waypointY
		case 3: // -270
			x1 = s.waypointY
			y1 = -s.waypointX
		default:
			log.Fatalf("Angle out of bounce: %d", value)
		}
		s.waypointX = x1
		s.waypointY = y1
	case 'F':
		s.x += value * s.waypointX
		s.y += value * s.waypointY
	default:
		log.Fatalf("Unknown command: %s", string(command))
	}
}

func (s *ShipPart2) getCurrentPosition() (int, int) {
	return s.x, s.y
}

func (s *ShipPart2) getWaypointPosition() (int, int) {
	return s.waypointX, s.waypointY
}

func (s *ShipPart2) getManhattanDistance() int {
	return absInt(s.x) + absInt(s.y)
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

	ship := newShip(3)

	for _, i := range instructions {
		ship.move(i)
	}

	fmt.Printf("Answer to part 1: %d\n", ship.getManhattanDistance())

	ship2 := newShipPart2()

	for _, i := range instructions {
		ship2.move(i)
	}

	fmt.Printf("Answer to part 2: %d\n", ship2.getManhattanDistance())
}
