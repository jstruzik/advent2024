package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Rule struct {
	After []int
}

type PrintOrderLine struct {
	Pages  []int
	Middle int
}

func main() {
	rules, printOrderLines := readFile()
	solve1(rules, printOrderLines)
	solve2(rules, printOrderLines)
}

func solve1(rules map[int]Rule, printOrderLines []PrintOrderLine) {
	var middleResults []int
	var middleSum int

lineLoop:
	for _, printOrderLine := range printOrderLines {
		pages := printOrderLine.Pages
		var badPages []int
		for i := len(pages) - 1; i >= 0; i-- {
			curPage := pages[i]
			rule := rules[curPage]
			badPages = append(badPages, rule.After...)
			for _, badPage := range badPages {
				if curPage == badPage {
					fmt.Println("Found a naughty page!!! Skipping to next line")
					continue lineLoop
				}
			}
		}
		fmt.Println("Santa likes this line. Adding to the list!")
		middleResults = append(middleResults, printOrderLine.Middle)
	}

	for _, middleResult := range middleResults {
		middleSum = middleSum + middleResult
	}
	fmt.Println("Total sum of correct middle results", middleSum)
}

func solve2(rules map[int]Rule, printOrderLines []PrintOrderLine) {
	var middleResults []int
	var middleSum int

lineLoop:
	for _, printOrderLine := range printOrderLines {
		pages := printOrderLine.Pages
		var badPages []int
		for i := len(pages) - 1; i >= 0; i-- {
			curPage := pages[i]
			rule := rules[curPage]
			badPages = append(badPages, rule.After...)
			for _, badPage := range badPages {
				if curPage == badPage {
					fmt.Println("Found a naughty page!!! Let's fix...")
					sort.Slice(pages, func(i2, j int) bool {
						for _, validPage := range rules[pages[i2]].After {
							if validPage == pages[j] {
								return true
							}
						}
						return false
					})
					pageMiddle := pages[len(pages)/2]
					middleResults = append(middleResults, pageMiddle)
					continue lineLoop
				}
			}
		}
	}

	for _, middleResult := range middleResults {
		middleSum = middleSum + middleResult
	}
	fmt.Println("Total sum of fixed middle results", middleSum)
}

func readFile() (map[int]Rule, []PrintOrderLine) {
	rules := make(map[int]Rule)
	var printOrderLines []PrintOrderLine

	file, err := os.Open("./resources/5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	r := regexp.MustCompile(`(?P<Lower>\d+)\|(?P<Upper>\d+)`)
	r2 := regexp.MustCompile(`\d+\,\d+`)
	for scanner.Scan() {
		line := scanner.Text()
		matches := r.FindAllStringSubmatch(line, -1)
		match2 := r2.FindStringSubmatch(line)
		for _, match := range matches {
			lower, err1 := strconv.Atoi(match[r.SubexpIndex("Lower")])
			upper, err2 := strconv.Atoi(match[r.SubexpIndex("Upper")])
			if err1 != nil || err2 != nil {
				fmt.Println("is not an integer.")
			}

			existingRule, exists := rules[lower]

			if exists {
				existingRule.After = append(existingRule.After, upper)
				rules[lower] = existingRule
			} else {
				rules[lower] = Rule{After: []int{upper}}
			}
		}
		if match2 != nil {
			parts := strings.Split(line, ",")
			var pages []int
			for _, part := range parts {
				num, err := strconv.Atoi(part)
				if err != nil {
					fmt.Println("could not convert number", err)
					continue
				}
				pages = append(pages, num)
			}
			pageMiddle := pages[len(pages)/2]
			printOrderLine := PrintOrderLine{
				Pages:  pages,
				Middle: int(pageMiddle),
			}
			printOrderLines = append(printOrderLines, printOrderLine)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return rules, printOrderLines
}
