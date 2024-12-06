package six

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type PuzzleMap struct {
	Lines   []PuzzleLine
	CursorX int
	CursorY int
	Width   int
	Height  int
}

type PuzzleLine struct {
	Positions []PuzzlePosition
}

type PuzzlePosition struct {
	IsObstacle  bool
	IsVisited   bool
	VisitedFrom string
}

const (
	UP    string = "up"
	DOWN  string = "down"
	LEFT  string = "left"
	RIGHT string = "right"
)

func Solve1(puzzleMap PuzzleMap) {
	fmt.Println(move(UP, puzzleMap, 1))
}

func Solve2(puzzleMap PuzzleMap) {
	possibleObstacles := 0
	for y, line := range puzzleMap.Lines {
		for x := range line.Positions {
			modifiedPuzzleMap := addObstruction(ReadFile("./resources/6.txt"), x, y)
			_, didMove := move(UP, modifiedPuzzleMap, 1)
			if !didMove {
				possibleObstacles++
			}
		}
	}
	fmt.Println(possibleObstacles)
}

func getNewDirection(direction string) string {
	newDirection := direction
	switch direction {
	case UP:
		newDirection = RIGHT
	case DOWN:
		newDirection = LEFT
	case LEFT:
		newDirection = UP
	case RIGHT:
		newDirection = DOWN
	}

	return newDirection
}

func wouldMoveBreachEdge(direction string, x int, y int, puzzleMap PuzzleMap) bool {
	if direction == UP && y == 0 {
		return true
	} else if direction == DOWN && y >= puzzleMap.Height-1 {
		return true
	} else if direction == LEFT && x == 0 {
		return true
	} else if direction == RIGHT && x >= puzzleMap.Width-1 {
		return true
	}
	return false
}

func move(direction string, puzzleMap PuzzleMap, positionsMoved int) (int, bool) {
	x := puzzleMap.CursorX
	y := puzzleMap.CursorY
	puzzleMap.Lines[y].Positions[x].IsVisited = true
	puzzleMap.Lines[y].Positions[x].VisitedFrom = direction

	//fmt.Println("current position", x, y, puzzleMap.Lines[y].Positions[x].IsObstacle)

	switch direction {
	case UP:
		y--
	case DOWN:
		y++
	case LEFT:
		x--
	case RIGHT:
		x++
	}

	newPosition := puzzleMap.Lines[y].Positions[x]
	//fmt.Println("next position", x, y, newPosition.IsObstacle, newPosition.IsVisited, newPosition.VisitedFrom)
	if newPosition.IsObstacle {
		if newPosition.IsVisited && newPosition.VisitedFrom == direction {
			//fmt.Println("Hit loop!")
			// We've hit an infinite loop
			return positionsMoved, false
		}
		// Mark the obstacle as visited as well
		puzzleMap.Lines[y].Positions[x].IsVisited = true
		puzzleMap.Lines[y].Positions[x].VisitedFrom = direction
		toMove := getNewDirection(direction)
		//fmt.Println("new dir", toMove)

		return move(toMove, puzzleMap, positionsMoved)
	}

	if !newPosition.IsVisited {
		positionsMoved++
	}

	if wouldMoveBreachEdge(direction, x, y, puzzleMap) {
		return positionsMoved, true
	}

	puzzleMap.CursorX = x
	puzzleMap.CursorY = y

	//fmt.Println("moving ", direction, " to ", x, y, " count ", positionsMoved)

	return move(direction, puzzleMap, positionsMoved)
}

func addObstruction(puzzleMap PuzzleMap, x int, y int) PuzzleMap {
	//fmt.Println("adding obstacle to", x, y)
	puzzleMap.Lines[y].Positions[x].IsObstacle = true

	return puzzleMap
}

func ReadFile(fileName string) PuzzleMap {
	var puzzleMap PuzzleMap
	var cursorX int
	var cursorY int
	var width int
	var height int
	var puzzleLines []PuzzleLine

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var nextLine string
	if scanner.Scan() {
		nextLine = scanner.Text()
	}
	for {
		line := nextLine
		if !scanner.Scan() {
			nextLine = ""
		} else {
			nextLine = scanner.Text()
		}

		var puzzlePositions []PuzzlePosition
		width = len(line)
		height++

		for idx, char := range line {
			isObstacle := false
			if string(char) == "#" {
				isObstacle = true
			}
			if string(char) == "^" {
				cursorX = idx
				cursorY = height - 1
			}
			puzzlePositions = append(puzzlePositions, PuzzlePosition{IsObstacle: isObstacle, IsVisited: false})
		}
		puzzleLines = append(puzzleLines, PuzzleLine{Positions: puzzlePositions})

		if nextLine == "" {
			break
		}
	}

	puzzleMap = PuzzleMap{
		Lines:   puzzleLines,
		CursorX: cursorX,
		CursorY: cursorY,
		Width:   width,
		Height:  height,
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//fmt.Println(puzzleMap)

	return puzzleMap
}
