package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

var reMatchColorAndContents = regexp.MustCompile(`(.*) contain ([\w ,]*)`)
var reMatchBagCountAndColor = regexp.MustCompile(`(\d*) (.*)`)

const requiredBagColorString = "shiny gold"

// BagColor is the color of a bag
type BagColor string

// Bag holds other bags
type Bag struct {
	color    BagColor
	contents map[BagColor]int // number of bags of that color
}

func cleanBagString(str string) string {
	r := strings.Replace(str, ".", "", -1)
	r = strings.Replace(r, "bags", "", -1)
	r = strings.TrimSpace(r)
	return strings.Replace(r, "bag", "", -1)
}

func extractCountAndColor(str string) (int, string, error) {
	matches := reMatchBagCountAndColor.FindAllStringSubmatch(str, -1)
	if len(matches) != 1 {
		return 0, "", fmt.Errorf("Something wrong with matches, as it is larger than 1 (len = %d): %v", len(matches), matches)
	}

	match := matches[0]
	if len(match) != 3 {
		return 0, "", fmt.Errorf("Something wrong with match, as it did not capture 4 groups (+1) (len = %d): %v", len(match), match)
	}

	val, err := strconv.Atoi(match[1])
	if err != nil {
		return 0, "", fmt.Errorf("Could not parse string '%s' to intger: %s", match[1], err)
	}

	return val, match[2], nil
}

func checkGoldBag(bagRules map[string]Bag, bag Bag) bool {
	for color := range bag.contents {
		if color == requiredBagColorString {
			return true
		}
		if checkGoldBag(bagRules, bagRules[string(color)]) {
			return true
		}
	}
	return false
}

func countBags(bagRules map[string]Bag, bag Bag) int {
	totalCount := 0
	for bag, count := range bag.contents {
		totalCount += count + count*countBags(bagRules, bagRules[string(bag)])
	}
	return totalCount
}

func main() {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file '%s': %s", inputFile, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	rules := []string{}
	for scanner.Scan() {
		input := scanner.Text()
		rules = append(rules, input)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	bagRules := map[string]Bag{}
	for _, rule := range rules {
		matches := reMatchColorAndContents.FindAllStringSubmatch(rule, -1)
		if len(matches) == 0 {
			log.Fatalf("Something wrong with matches, as it is larger than 1 (len = %d): %v", len(matches), matches)
		}

		match := matches[0]
		if len(match) != 3 {
			log.Fatalf("Something wrong with match, as it did not capture 4 groups (+1) (len = %d): %v", len(match), match)
		}

		bag := match[1]
		bag = cleanBagString(bag)
		rulesForBag := make(map[BagColor]int)
		r := cleanBagString(match[2])
		if r == "no other" {
			continue
		}
		splitted := strings.Split(r, ",")
		for _, r := range splitted {
			trimmed := strings.TrimSpace(r)
			count, color, err := extractCountAndColor(trimmed)
			if err != nil {
				log.Fatalln(err)
			}
			rulesForBag[BagColor(color)] += count
		}
		bagRules[bag] = Bag{color: BagColor(bag), contents: rulesForBag}
	}

	goldCount := 0
	for _, bag := range bagRules {
		if bag.color == requiredBagColorString {
			continue
		}
		if checkGoldBag(bagRules, bag) {
			goldCount++
		}
	}

	fmt.Printf("Answer to part 1: There are %d bags which can carry at least one %s bag\n", goldCount, requiredBagColorString)

	count := countBags(bagRules, bagRules[requiredBagColorString])

	fmt.Printf("\nAnswer to part 2: There are %d individual bags required inside my single %s bag\n", count, requiredBagColorString)
}
