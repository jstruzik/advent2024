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
	var resultCacheMap = make(map[[2]int]int)

	if problem == 1 {
		blinkAmount := 25
		return blink(stoneMap, 1, blinkAmount)
	} else if problem == 2 {
		blinkAmount := 75
		for i := 0; i < len(stoneMap); i++ {
			numStones := blinkPerNumber(stoneMap[i], 1, blinkAmount, resultCacheMap)
			totalStones += numStones
		}
	}

	return totalStones
}

func blink(stoneMap []int, count int, countMax int) int {
	stones := 0
	var newStones []int
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
			leftNum, _ := strconv.Atoi(leftStr)
			rightNum, _ := strconv.Atoi(rightStr)

			stoneMap[i] = leftNum
			newStones = append(newStones, rightNum)
			stones += 2
		} else {
			stoneMap[i] = stoneMap[i] * 2024
			stones++
		}
	}

	stoneMap = append(stoneMap, newStones...)
	//fmt.Println(stoneMap)
	if count == countMax {
		return stones
	}
	count++
	fmt.Println("Stone count", stones)

	return blink(stoneMap, count, countMax)
}

func blinkPerNumber(num int, count int, countMax int, resultCache map[[2]int]int) int {
	if val, exists := resultCache[[2]int{count, num}]; exists {
		return val
	}
	strNum := strconv.Itoa(num)
	if count == countMax {
		if len(strNum)%2 == 0 {
			return 2
		}
		return 1
	}
	if num == 0 {
		return blinkPerNumber(1, count+1, countMax, resultCache)
	} else if len(strNum)%2 == 0 {
		mid := len(strNum) / 2
		leftStr, rightStr := strNum[:mid], strNum[mid:]
		leftNum, _ := strconv.Atoi(leftStr)
		rightNum, _ := strconv.Atoi(rightStr)
		val := blinkPerNumber(leftNum, count+1, countMax, resultCache) + blinkPerNumber(rightNum, count+1, countMax, resultCache)
		resultCache[[2]int{count, num}] = val
		return val
	} else {
		val := blinkPerNumber(num*2024, count+1, countMax, resultCache)
		resultCache[[2]int{count, num}] = val
		return val
	}
}
