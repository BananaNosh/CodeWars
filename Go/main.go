package main

import (
	"fmt"
)

type Positioning struct {
	positionedPieces [6]int
	rotations        [6]int
	flipped          [6]bool
}

func (p Positioning) getPieces(pcs [6][6][6]int) [6][6][6]int {
	var pieces [6][6][6]int
	for i, pieceIndex := range p.positionedPieces {
		correctPieceIndex := pieceIndex - 1
		pieces[i] = mutatedPiece(pcs[correctPieceIndex], p.rotations[i], p.flipped[i])
	}
	return pieces
}

func SolveSnafooz(pcs [6][6][6]int) [6][6][6]int {
	hasSolution, solution := solveRec(pcs[:], Positioning{}, 0)
	var solutionPcs [6][6][6]int
	if hasSolution {
		solutionPcs = solution.getPieces(pcs)
		printPositioning(solution)
	} else {
		panic("Not solvable")
	}
	return solutionPcs
}

func solveRec(pcs [][6][6]int, positioning Positioning, depth int) (bool, Positioning) {
	currentPosition := -1
	var isSet [6]bool
	for i, p := range positioning.positionedPieces {
		if p == 0 {
			currentPosition = i
			break
		}
		isSet[p-1] = true
	}
	if currentPosition == -1 {
		return true, positioning
	}
	for i := 1; i < 7; i++ {
		if isSet[i-1] {
			continue
		}
		positioning.positionedPieces[currentPosition] = i
		for rotation := 0; rotation < 4; rotation++ {
			positioning.rotations[currentPosition] = rotation
			for _, flipped := range []bool{false, true} {
				positioning.flipped[currentPosition] = flipped
				for i := 0; i < depth; i++ {
					fmt.Print("   ")
				}
				fmt.Printf("Check %v, rot:%v, flipped: %v\n", i, rotation, flipped)
				if check(pcs, currentPosition, positioning) {
					//fmt.Printf("depth check succesful: %v\n", depth)
					//printPositioning(positioning)
					hasSolution, solution := solveRec(pcs, positioning, depth+1)
					if hasSolution {
						return true, solution
					}
				}
				if currentPosition == 0 {
					//fmt.Printf("depth no solution: %v\n", depth)
					return false, Positioning{}
				}
			}
		}
	}
	//fmt.Printf("depth no solution: %v\n", depth)
	return false, Positioning{}
}

func mutatedPiece(piece [6][6]int, rotation int, flipped bool) [6][6]int {
	if flipped {
		flipPiece(piece[:])
	}
	rotatePiece(piece[:], rotation)
	return piece
}

func flipPiece(piece [][6]int) {
	for i, j := 0, len(piece[0])-1; i < j; i, j = i+1, j-1 {
		piece[0][i], piece[0][j] = piece[0][j], piece[0][i]
		piece[5][i], piece[5][j] = piece[5][j], piece[5][i]
	}
	for i := 1; i < 5; i++ {
		piece[i][0], piece[i][5] = piece[i][5], piece[i][0]
	}
}

func rotatePiece(piece [][6]int, rotation int) {
	rotation = rotation % 4
	var borders [4][6]int
	for i := 0; i < 6; i++ {
		borders[0][i] = piece[0][i]
		borders[1][i] = piece[i][5]
		borders[2][i] = piece[5][5-i]
		borders[3][i] = piece[5-i][0]
	}
	rotated := append(borders[4-rotation:], borders[:4-rotation]...)
	for i := 0; i < 6; i++ {
		piece[0][i] = rotated[0][i]
		piece[i][5] = rotated[1][i]
		piece[5][5-i] = rotated[2][i]
		piece[5-i][0] = rotated[3][i]
	}
}

func check(pcs [][6][6]int, newPosition int, positioning Positioning) bool {
	neighbours := [6][4][3]int{
		{{6, 0, 2},
			{4, 3, 0},
			{3, 2, 0},
			{2, 1, 0}},
		{{1, 0, 1},
			{3, 3, 1},
			{5, 2, 1},
			{6, 1, 1}},
		{{1, 0, 2},
			{4, 3, 1},
			{5, 2, 0},
			{2, 1, 3}},
		{{1, 0, 3},
			{6, 3, 3},
			{5, 2, 3},
			{3, 1, 3}},
		{{3, 0, 2},
			{4, 3, 2},
			{6, 2, 0},
			{2, 1, 2}},
		{{5, 0, 2},
			{4, 3, 3},
			{1, 2, 0},
			{2, 1, 1}},
	}
	for k, n := range neighbours[newPosition] {
		pieceAtNeighbourPositionIndex := positioning.positionedPieces[n[0]-1]
		if pieceAtNeighbourPositionIndex == 0 {
			continue
		}
		currentPiece := pcs[positioning.positionedPieces[newPosition]-1]
		if positioning.flipped[newPosition] {
			flipPiece(currentPiece[:])
		}
		rotatePiece(currentPiece[:], n[1]+positioning.rotations[newPosition]) // add rotation of piece to rotation to have the neighbouring site on top
		flipPiece(currentPiece[:])
		// additional flip to counteract the mirror effect of neighbouring pieces
		neighbourPiece := pcs[pieceAtNeighbourPositionIndex-1]
		if positioning.flipped[n[0]-1] {
			flipPiece(neighbourPiece[:])
		}
		rotatePiece(neighbourPiece[:], n[2]+positioning.rotations[n[0]-1])
		for j := 1; j < 5; j++ {
			if currentPiece[0][j]+neighbourPiece[0][j] != 1 {
				return false
			}
		}
		for _, j := range []int{0, 5} {
			sum := currentPiece[0][j] + neighbourPiece[0][j]
			if sum == 2 {
				return false
			}

			var cornerNeighbourDif int
			if j == 0 {
				cornerNeighbourDif = 1
			} else {
				cornerNeighbourDif = 3
			}
			cornerNeighbour := neighbours[newPosition][(k+cornerNeighbourDif)%4]
			cornerNeighbourPositionIndex := positioning.positionedPieces[cornerNeighbour[0]-1]
			if cornerNeighbourPositionIndex == 0 {
				continue
			} else {
				cornerNeighbourPiece := pcs[cornerNeighbourPositionIndex-1]
				if positioning.flipped[cornerNeighbour[0]-1] {
					flipPiece(cornerNeighbourPiece[:])
				}
				rotatePiece(cornerNeighbourPiece[:], cornerNeighbour[2]+positioning.rotations[cornerNeighbour[0]-1])
				corner := cornerNeighbourPiece[0][5-j]
				if corner+sum != 1 {
					return false
				}
			}
		}
	}
	return true
}

func printPieces(pcs [6][6][6]int) {
	for _, p := range pcs {
		printPiece(p)
	}
}

func printPiecesStructured(pcs [6][6][6]int) {
	lineIndices := [72][2]int{}
	for i := 0; i < 6; i++ {
		lineIndices[i*3] = [2]int{-1, -1}
		lineIndices[i*3+1] = [2]int{0, i}
		lineIndices[i*3+2] = [2]int{-1, -1}
	}
	for i := 0; i < 6; i++ {
		offset := 18
		for j := 0; j < 3; j++ {
			lineIndices[offset+i*3+j] = [2]int{j + 1, i}
		}
	}
	for k := 0; k < 2; k++ {
		for i := 0; i < 6; i++ {
			offset := 36 + 18*k
			lineIndices[offset+i*3] = [2]int{-1, -1}
			lineIndices[offset+i*3+1] = [2]int{4 + k, i}
			lineIndices[offset+i*3+2] = [2]int{-1, -1}
		}
	}
	for i, lineIndex := range lineIndices {
		if i > 0 && i%3 == 0 {
			fmt.Println()
		} else if i > 0 {
			fmt.Print("| ")
		}
		if i > 0 && i%18 == 0 {
			for j := 0; j < 20; j++ {
				fmt.Print("--")
			}
			fmt.Println()
		}
		lineStr := ""
		if lineIndex[0] >= 0 {
			line := pcs[lineIndex[0]][lineIndex[1]]
			for j := 0; j < 6; j++ {
				lineStr += fmt.Sprintf("%v ", line[j])
			}
		} else {
			for j := 0; j < 6; j++ {
				lineStr += "  "
			}
		}
		fmt.Printf(lineStr)
	}
	fmt.Println()
}

func printPiece(p [6][6]int) {
	fmt.Println()
	for i := 0; i < 6; i++ {
		fmt.Println()
		for j := 0; j < 6; j++ {
			fmt.Print(p[i][j])
			fmt.Print(" ")
		}
	}
}

func printPositioning(positioning Positioning) {
	fmt.Printf("  %v  \n", positioning.positionedPieces[0])
	fmt.Printf("%v %v %v\n", positioning.positionedPieces[1], positioning.positionedPieces[2], positioning.positionedPieces[3])
	fmt.Printf("  %v  \n", positioning.positionedPieces[4])
	fmt.Printf("  %v  \n", positioning.positionedPieces[5])
	fmt.Printf("  %v  \n", positioning.rotations[0])
	fmt.Printf("%v %v %v\n", positioning.rotations[1], positioning.rotations[2], positioning.rotations[3])
	fmt.Printf("  %v  \n", positioning.rotations[4])
	fmt.Printf("  %v  \n", positioning.rotations[5])
	fmt.Printf("  %v  \n", positioning.flipped[0])
	fmt.Printf("%v %v %v\n", positioning.flipped[1], positioning.flipped[2], positioning.flipped[3])
	fmt.Printf("  %v  \n", positioning.flipped[4])
	fmt.Printf("  %v  \n", positioning.flipped[5])
}

func main() {
	pieces := [6][6][6]int{{
		{0, 0, 1, 1, 0, 0},
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
	pieces = [6][6][6]int{{ // as in example
		{1, 1, 0, 1, 0, 0},
		{1, 1, 1, 1, 1, 1},
		{0, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1},
		{0, 0, 1, 0, 1, 1}},

		{{0, 0, 1, 0, 0, 0},
			{0, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 0},
			{1, 1, 1, 0, 0, 0}},

		{{1, 1, 0, 1, 0, 0},
			{0, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{0, 1, 0, 0, 1, 1}},

		{{0, 0, 0, 0, 0, 0},
			{0, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 0},
			{0, 1, 0, 1, 0, 0}},

		{{1, 0, 1, 1, 0, 0},
			{1, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 1},
			{0, 0, 1, 1, 1, 0}},

		{{0, 1, 0, 0, 0, 1},
			{1, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 1},
			{0, 0, 1, 0, 1, 1}}}
	pieces = [6][6][6]int{{ // already correct
		{1, 1, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1, 1},
		{1, 1, 1, 0, 0, 0}}, // 1

		{{0, 0, 1, 1, 0, 0},
			{1, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 0, 1, 1}}, // 2

		{{0, 0, 0, 1, 1, 1},
			{0, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 0}}, // 3

		{{0, 0, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 0},
			{1, 1, 1, 0, 1, 0}}, // 4

		{{0, 0, 0, 0, 0, 0},
			{0, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 0},
			{0, 1, 0, 0, 1, 0}}, // 5

		{{0, 0, 1, 1, 0, 1},
			{0, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 0},
			{0, 0, 1, 0, 1, 0}}} // 6

	//rotatePiece(pieces[1][:], 2)
	//rotatePiece(pieces[3][:], 3)
	//flipPiece(pieces[2][:])
	//rotatePiece(pieces[4][:], 1)
	//snafoozPieces := SolveSnafooz(pieces)
	SolveSnafooz(pieces)
	//printPieces(snafoozPieces)
	printPiecesStructured(pieces)

	//pieces[1] = mutatedPiece(pieces[0], 1, true)
	//printPieces(pieces)

	pieces2 := [6][6][6]int{
		{{0, 0, 1, 1, 0, 0},
			{1, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 0},
			{0, 0, 1, 1, 0, 0}},

		{{1, 1, 0, 0, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 0, 1, 1, 1}},

		{{0, 0, 1, 0, 0, 0},
			{1, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 0},
			{0, 1, 0, 0, 1, 0}},

		{{0, 0, 1, 1, 0, 0},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 1},
			{0, 0, 1, 1, 0, 0}},

		{{0, 0, 1, 1, 0, 0},
			{0, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 0},
			{1, 1, 0, 0, 1, 1}},

		{{0, 0, 1, 1, 0, 1},
			{0, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 1},
			{0, 1, 0, 0, 1, 1}}}
	out := SolveSnafooz(pieces2)
	printPiecesStructured(out)
	//var shifted [6][6][6]int
	//for i := 0; i < 6; i++ {
	//	shifted[i] = out[(i+3)%6]
	//}
	//fmt.Println()
	//printPiecesStructured(shifted)
}
