package main

import (
	"advent2024/five"
	"advent2024/six"
	"fmt"
)

func main() {
	fmt.Println("Running Five package...")
	rules, printOrderLines := five.ReadFile()
	five.Solve1(rules, printOrderLines)
	five.Solve2(rules, printOrderLines)

	fmt.Println("Running Six package...")
	puzzleMap := six.ReadFile("./resources/6.txt")
	six.Solve1(puzzleMap)
	six.Solve2(puzzleMap)
}
