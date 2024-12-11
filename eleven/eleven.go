package eleven

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve(fileName string, problem int) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var stoneMap []int

	for scanner.Scan() {
		line := scanner.Text()

		for _, num := range strings.Split(line, " ") {
			val, err := strconv.Atoi(string(num))
			if err != nil {
				fmt.Println("could not convert number", err)
				continue
			}
			stoneMap = append(stoneMap, val)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(stoneMap)
	totalStones := 0

	if problem == 1 {
		totalStones = blink(stoneMap, 1, 75)
	}

	return totalStones
}

func blink(stoneMap []int, count int, countMax int) int {
	stones := 0
	fmt.Println("Blinking", count)
	for i := 0; i < len(stoneMap); i++ {
		num := stoneMap[i]
		strNum := strconv.Itoa(num)
		if num == 0 {
			stoneMap[i] = 1
			stones++
		} else if len(strNum)%2 == 0 {
			mid := len(strNum) / 2
			leftStr, rightStr := strNum[:mid], strNum[mid:]
			leftNum, err := strconv.Atoi(string(leftStr))
			if err != nil {
				fmt.Println("could not convert number", err)
				continue
			}
			rightNum, err := strconv.Atoi(string(rightStr))
			if err != nil {
				fmt.Println("could not convert number", err)
				continue
			}
			stoneMap[i] = leftNum
			stoneMap = append(stoneMap, 0)
			copy(stoneMap[i+1:], stoneMap[i:])
			stoneMap[i+1] = rightNum
			i++
			stones += 2
		} else {
			stoneMap[i] = stoneMap[i] * 2024
			stones++
		}
	}
	//fmt.Println(stoneMap)
	if count == countMax {
		return stones
	}
	count++

	return blink(stoneMap, count, countMax)
}
