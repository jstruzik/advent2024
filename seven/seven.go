package seven

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	ADD      string = "add"
	MULTIPLY string = "multiply"
	CONCAT   string = "concat"
)

func genOpPermutations(carry []string, ops []string, length int, nums []int, val int, foundMatch bool) bool {
	if length == 0 {
		// Start with first value
		evalNum := nums[0]
		// Evaluate
		for n := 0; n < len(carry); n++ {
			if carry[n] == MULTIPLY {
				evalNum = evalNum * nums[n+1]
			} else if carry[n] == ADD {
				evalNum = evalNum + nums[n+1]
			} else if carry[n] == CONCAT {
				evalString := strconv.Itoa(evalNum) + strconv.Itoa(nums[n+1])
				convEval, err := strconv.Atoi(evalString)
				if err != nil {
					fmt.Println("could not convert back to number!!!", err)
				}
				evalNum = convEval
			}
		}
		foundMatch = evalNum == val
		// fmt.Println(length, carry, evalNum, val, foundMatch)

		return foundMatch
	}

	for i := 0; i < len(ops); i++ {
		newCarry := carry
		newCarry = append(newCarry, ops[i])
		foundMatch = genOpPermutations(newCarry, ops, length-1, nums, val, foundMatch)
		if foundMatch {
			return true
		}
	}

	return foundMatch
}

func Solve(fileName string, problem int) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sumOfVals := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var nums []int
		split1 := strings.Split(line, ":")
		val, err := strconv.Atoi(split1[0])
		if err != nil {
			fmt.Println("could not convert number", err)
			continue
		}
		split2 := strings.Split(split1[1], " ")
		for _, num := range split2 {
			// fmt.Println(num)
			convNum, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("could not convert number", err)
				continue
			}
			nums = append(nums, convNum)
		}
		fmt.Println(val, nums)
		validOps := []string{ADD, MULTIPLY}
		if problem == 2 {
			validOps = append(validOps, CONCAT)
		}
		var carry []string
		canEval := genOpPermutations(carry, validOps, len(nums)-1, nums, val, false)

		fmt.Println(canEval)

		if canEval {
			sumOfVals += val
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sumOfVals
}
