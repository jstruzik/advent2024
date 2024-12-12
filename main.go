package main

import (
	"advent2024/eleven"
	"fmt"
)

func main() {
	fmt.Println("Running Five package...")
	//rules, printOrderLines := five.ReadFile()
	//five.Solve1(rules, printOrderLines)
	//five.Solve2(rules, printOrderLines)

	fmt.Println("Running Six package...")
	//puzzleMap := six.ReadFile("./resources/6.txt")
	//six.Solve1(puzzleMap)
	//six.Solve2(puzzleMap)

	fmt.Println("Running Seven package...")
	// totalSum := seven.Solve("./resources/7.txt", 1)
	// fmt.Println(totalSum)
	// totalSum = seven.Solve("./resources/7.txt", 2)
	// fmt.Println(totalSum)

	fmt.Println("Running Nine package...")
	//checkSum := nine.Solve("./resources/9.txt", 2)
	//fmt.Println(checkSum)

	fmt.Println("Running Eleven package...")
	stones := eleven.Solve("./resources/11.txt", 1)
	fmt.Println("Problem 1 stones:", stones)
	stones = eleven.Solve("./resources/11.txt", 2)
	fmt.Println("Problem 2 stones:", stones)
}
