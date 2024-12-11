package eight

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	ADD      string = "add"
	MULTIPLY string = "multiply"
	CONCAT   string = "concat"
)

func Solve(fileName string, problem int) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var fileMap []int
	var id int
	blankSpaces := 0

	for scanner.Scan() {
		line := scanner.Text()

		for idx, char := range line {
			val, err := strconv.Atoi(string(char))
			if err != nil {
				fmt.Println("could not convert number", err)
				continue
			}
			if idx%2 == 0 {
				for i := 0; i < val; i++ {
					fileMap = append(fileMap, id)
				}
				id++
			} else {
				for i := 0; i < val; i++ {
					fileMap = append(fileMap, -1)
					blankSpaces++
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(fileMap)

	if problem == 1 {
		return solve1(fileMap)
	} else {
		return solve2(fileMap)
	}
}

func solve1(fileMap []int) int {
forwardLoop:
	for i := 0; i < len(fileMap); i++ {
		if fileMap[i] == -2 {
			break
		}
		if fileMap[i] == -1 {
			for j := len(fileMap) - 1; j >= 0; j-- {
				if fileMap[j] > 0 {
					// Swap! Mark swapped values
					fileMap[i], fileMap[j] = fileMap[j], -2
					continue forwardLoop
				}
			}
		}
	}

	fmt.Println(fileMap)
	checkSum := 0

	for i := 0; i < len(fileMap); i++ {
		if fileMap[i] > 0 {
			checkSum += fileMap[i] * i
		}
	}

	return checkSum
}

func solve2(fileMap []int) int {
	// Move backwards through file
	for j := len(fileMap) - 1; j >= 0; j-- {
		// Grab the start of a block that isn't blank
		if fileMap[j] > 0 {
			// mark the starts and ends
			blockStart := j
			blockEnd := j
			// Now we iterate backwards to find the start of the block
			for n := j; n >= 0; n-- {
				if fileMap[n] != fileMap[j] {
					// Ensure we mark our start as the value to the right
					blockStart = n + 1
					// We've found our start, stop iterating
					//fmt.Println("start", blockStart, "end", blockEnd)
					break
				}
			}
			// Calculate the length of the block
			blockDiff := blockEnd - blockStart
			// Now let's find a block of blanks we can fit into
			blankBlockStart := 0
			blankBlockEnd := 0
		blankLoop:
			// Move forward to find the first blank block that can fit
			for i := 0; i < len(fileMap); i++ {
				if fileMap[i] == -1 {
					blankBlockStart = i
					// Keep moving forward to find the end of the block
					for m := i; m < len(fileMap); m++ {
						if fileMap[m] != fileMap[i] {
							// Ensure we mark the end of the block to be the previous space
							blankBlockEnd = m - 1
							// Calculate the length of the blank block
							blankBlockDiff := blankBlockEnd - blankBlockStart
							//fmt.Println("start", blockStart, "end", blockEnd, "blankstart", blankBlockStart, "blankend", blankBlockEnd)
							// We've found one big enough and it doesn't intersect
							if blankBlockDiff >= blockDiff && blankBlockEnd < blockStart {
								// Swap our blocks
								for t := 0; t <= blockDiff; t++ {
									fileMap[blockStart+t], fileMap[blankBlockStart+t] = fileMap[blankBlockStart+t], fileMap[blockStart+t]
									//fmt.Println(fileMap)
								}
								// Break out since we've swapped. We're done.
								//fmt.Println("Done swapping! moving on")
								break
							} else {
								// Otherwise, continue to find the next blank block
								continue blankLoop
							}
						}
					}
					//fmt.Println("Let's move onto next block...")
					// If we've gotten here, then we've already found a match and swapped blocks or we couldn't find one that fits
					j = blockStart
					break
				}
			}
		}
	}

	fmt.Println(fileMap)
	checkSum := 0

	for i := 0; i < len(fileMap); i++ {
		if fileMap[i] > 0 {
			checkSum += fileMap[i] * i
		}
	}

	return checkSum
}
