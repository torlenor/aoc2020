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

var re = regexp.MustCompile(`(?m)(\d*)-(\d*) (.): (.*)`)

type passwordAndPolicy struct {
	requiredChar rune
	minAmount    int
	maxAmount    int

	password string
}

func (p *passwordAndPolicy) validOldPolicy() bool {
	foundRequiredCharCnt := 0
	for _, b := range p.password {
		if b == p.requiredChar {
			foundRequiredCharCnt++
		}
	}
	if foundRequiredCharCnt >= p.minAmount && foundRequiredCharCnt <= p.maxAmount {
		return true
	}
	return false
}

func (p *passwordAndPolicy) validNewPolicy() bool {
	result := false
	for i, b := range p.password {
		if i == p.minAmount-1 {
			if b != p.requiredChar {
				result = !result
			}
		}
		if i == p.maxAmount-1 {
			if b != p.requiredChar {
				result = !result
			}
		}
	}
	return result
}

func parseFromString(str string) (passwordAndPolicy, error) {
	matches := re.FindAllStringSubmatch(str, -1)
	if len(matches) != 1 {
		return passwordAndPolicy{}, fmt.Errorf("Something wrong with matches, as it is larger than 1 (len = %d): %v", len(matches), matches)
	}

	match := matches[0]
	if len(match) != 5 {
		return passwordAndPolicy{}, fmt.Errorf("Something wrong with match, as it did not capture 4 groups (+1) (len = %d): %v", len(match), match)
	}

	requiredChar := []rune(match[3])[0]
	minAmount, err := strconv.Atoi(string(match[1]))
	if err != nil {
		return passwordAndPolicy{}, fmt.Errorf("Error converting '%s' to integer: %s", match[1], err)
	}
	maxAmount, err := strconv.Atoi(string(match[2]))
	if err != nil {
		return passwordAndPolicy{}, fmt.Errorf("Error converting '%s' to integer: %s", match[2], err)
	}

	if minAmount > maxAmount {
		return passwordAndPolicy{}, fmt.Errorf("Something wrong with line '%s': minAmount = %d, maxAmount = %d", str, minAmount, maxAmount)
	}

	pwdAndPolicy := passwordAndPolicy{
		requiredChar: requiredChar,
		minAmount:    minAmount,
		maxAmount:    maxAmount,
		password:     match[4],
	}
	return pwdAndPolicy, nil
}

func main() {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file '%s': %s", inputFile, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	passwords := []passwordAndPolicy{}
	for scanner.Scan() {
		input := scanner.Text()
		password, err := parseFromString(input)
		if err != nil {
			log.Fatalf("Error parsing password line '%s': %s", input, err)
		}
		passwords = append(passwords, password)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	validCnt := 0
	for _, password := range passwords {
		if password.validOldPolicy() {
			validCnt++
		}
	}
	fmt.Printf("Answer to part 1: There are %d valid passwords in file %s\n", validCnt, inputFile)

	validCnt = 0
	for _, password := range passwords {
		if password.validNewPolicy() {
			validCnt++
		}
	}
	fmt.Printf("Answer to part 2: There are %d valid passwords in file %s\n", validCnt, inputFile)
}
