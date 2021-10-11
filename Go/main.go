package main

import "fmt"

func SolveSnafooz(pcs [6][6][6]int) [6][6][6]int {
	hasSolution, solution := solveRec(pcs[:], [6]int{})
	var solutionPcs [6][6][6]int
	if hasSolution {
		for i, pieceIndex := range solution {
			solutionPcs[i] = pcs[pieceIndex-1]
		}
	} else {
		panic("Not solvable")
	}
	return solutionPcs
}

func solveRec(pcs [][6][6]int, solutionPart [6]int) (bool, [6]int) {
	position := -1
	for i, p := range solutionPart {
		if p == 0 {
			position = i
			break
		}
	}
	if position == -1 {
		return true, solutionPart
	}
	for i := 1; i < 7; i++ {
		solutionPart[position] = i
		if check(pcs, position, solutionPart) {
			hasSolution, solution := solveRec(pcs, solutionPart)
			if hasSolution {
				return true, solution
			}
		}
	}
	return false, [6]int{}
}

func check(pcs [][6][6]int, position int, solutionPart [6]int) bool {
	return true
}

func printPieces(pcs [6][6][6]int) {
	for _, p := range pcs {
		fmt.Println()
		for i := 0; i < 6; i++ {
			fmt.Println()
			for j := 0; j < 6; j++ {
				fmt.Print(p[i][j])
				fmt.Print(" ")
			}
		}
	}
}

func main() {
	pieces := [6][6][6]int{{{0, 0, 1, 1, 0, 0},
		{0, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1, 1},
		{1, 0, 1, 0, 1, 1}},

		{{0, 1, 0, 0, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 1},
			{0, 0, 1, 1, 0, 1}},

		{{0, 0, 1, 1, 0, 1},
			{1, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 1},
			{0, 0, 1, 1, 0, 0}},

		{{0, 0, 1, 1, 0, 0},
			{0, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 0},
			{0, 1, 0, 0, 1, 0}},

		{{0, 0, 1, 1, 0, 1},
			{1, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 0, 0, 1, 1}},

		{{0, 0, 1, 1, 0, 0},
			{0, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 1},
			{0, 1, 0, 0, 1, 0}}}
	printPieces(SolveSnafooz(pieces))
}
