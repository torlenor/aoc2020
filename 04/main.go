package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

// const inputFile = "input_example_part2.txt"

var eyeColors = []string{
	"amb", "blu", "brn", "gry", "grn", "hzl", "oth",
}

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string

	// 	cid (Country ID) - ignored, missing or not.
	cid string
}

func passportFromString(str string) (passport, error) {
	re := regexp.MustCompile(`(?m)(\w*):(\W?\w*)`)

	p := passport{}

	matches := re.FindAllStringSubmatch(str, -1)
	for _, match := range matches {
		if len(match) != 3 {
			return passport{}, fmt.Errorf("Error parsing '%s'", str)
		}
		key := match[1]
		value := match[2]
		switch key {
		case "byr":
			p.byr = value
		case "iyr":
			p.iyr = value
		case "eyr":
			p.eyr = value
		case "hgt":
			p.hgt = value
		case "hcl":
			p.hcl = value
		case "ecl":
			p.ecl = value
		case "pid":
			p.pid = value
		case "cid":
			p.cid = value
		}
	}

	return p, nil
}

func validateAndReturnInt(in string, digits int) (int, bool) {
	if len(in) != digits {
		return -1, false
	}

	value, err := strconv.Atoi(in)
	if err != nil {
		return -1, false
	}

	return value, true
}

func (p *passport) validByr() bool {
	// 	byr (Birth Year) - four digits; at least 1920 and at most 2002.
	value, ok := validateAndReturnInt(p.byr, 4)
	if !ok {
		return false
	}

	if value < 1920 || value > 2002 {
		return false
	}

	return true
}

func (p *passport) validIyr() bool {
	// 	iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	value, ok := validateAndReturnInt(p.iyr, 4)
	if !ok {
		return false
	}

	if value < 2010 || value > 2020 {
		return false
	}

	return true
}

func (p *passport) validHgt() bool {
	// 	hgt (Height) - a number followed by either cm or in:
	// 	If cm, the number must be at least 150 and at most 193.
	// 	If in, the number must be at least 59 and at most 76.
	re := regexp.MustCompile(`(?m)(\d*)(cm|in)`)

	matches := re.FindAllStringSubmatch(p.hgt, -1)
	if len(matches) != 1 {
		return false
	}

	match := matches[0]
	if len(match) != 3 {
		return false
	}

	value, err := strconv.Atoi(match[1])
	if err != nil {
		return false
	}
	unit := match[2]

	switch unit {
	case "cm":
		if value < 150 || value > 193 {
			return false
		}
	case "in":
		if value < 59 || value > 76 {
			return false
		}
	default:
		return false
	}

	return true
}

func (p *passport) validHcl() bool {
	// 	hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	re := regexp.MustCompile(`(?m)#([a-f]|[0-9]){6}`)

	matches := re.FindAllStringSubmatch(p.hcl, -1)
	if len(matches) != 1 {
		return false
	}

	return true
}

func (p *passport) validEyr() bool {
	// 	eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	value, ok := validateAndReturnInt(p.eyr, 4)
	if !ok {
		return false
	}

	if value < 2020 || value > 2030 {
		return false
	}

	return true
}

func (p *passport) validPid() bool {
	// 	pid (Passport ID) - a nine-digit number, including leading zeroes.

	if len(p.pid) != 9 {
		return false
	}

	re := regexp.MustCompile(`(?m)([0-9]){9}`)

	matches := re.FindAllStringSubmatch(p.pid, -1)
	if len(matches) != 1 {
		return false
	}

	return true
}

func (p *passport) validEcl() bool {
	// 	ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	if len(p.ecl) != 3 {
		return false
	}
	valid := false
	for _, color := range eyeColors {
		if strings.Contains(p.ecl, color) {
			valid = true
			break
		}
	}

	return valid
}

func (p *passport) valid() bool {
	return p.validByr() &&
		p.validIyr() &&
		p.validEyr() &&
		p.validHgt() &&
		p.validHcl() &&
		p.validEcl() &&
		p.validPid()
}

func main() {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Not able to read file '%s': %s", inputFile, err)
	}

	input := string(data)

	re := regexp.MustCompile(`(.{1,})\n`)
	input = re.ReplaceAllString(input, "${1} ")
	splittedInput := strings.Split(input, "\n")

	requiredField := []string{
		"byr:",
		"iyr:",
		"eyr:",
		"hgt:",
		"hcl:",
		"ecl:",
		"pid:",
		// "cid:",
	}

	validPassports := []passport{}
	for _, passport := range splittedInput {
		valid := true
		for _, field := range requiredField {
			if !strings.Contains(passport, field) {
				valid = false
				break
			}
		}
		if valid {
			p, err := passportFromString(passport)
			if err != nil {
				log.Fatalf("Error parsing passport string '%s': %s", passport, err)
			}
			validPassports = append(validPassports, p)
		}
	}

	fmt.Printf("Answer to part 1: %d valid passports\n", len(validPassports))

	// Part 2
	validPassportCount := 0
	for _, p := range validPassports {
		if p.valid() {
			validPassportCount++
		}
	}

	fmt.Printf("\nAnswer to part 2: %d valid passports\n", validPassportCount)
}
