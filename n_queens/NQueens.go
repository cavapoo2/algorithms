package main

import "fmt"

/*board can be arrange as a column like this (example is a 4 by 4 board)
[1,3,0,2]
this means a queen is placed in column 1 for row 0, column 3 for row 1
column 0 for row 2, and column 2 for row 3*/
type board []int

type boards []board

func abs(v int) int {
	if v < 0 {
		return -v
	} else {
		return v
	}
}

//check if no column align (cant have queens in same column)
func noColumns(b board) bool {
	newIndex := len(b) - 1
	for i := 0; i < newIndex; i++ {
		diff := b[i] - b[newIndex]
		if diff == 0 {
			return false
		}
	}
	return true
}

//check no queens are on the same diagonal
//this can be done by  comparing equating the difference in row index
//and their values at the given row. e.g at rows 0 and 3, we have queens at column 0
//and column 3, clearly a diagonal match.
func noDiags(b board) bool {
	newIndex := len(b) - 1
	for i := 0; i < newIndex; i++ {
		diffValue := abs(b[i] - b[newIndex])
		diffIndex := newIndex - i
		if diffValue == diffIndex {
			return false
		}
	}
	return true
}

//compute n queens board for given N
func compute(N int) boards {
	var bs boards
	var b board
	var findBoards func(row int, brd board)
	findBoards = func(row int, brd board) {
		if row == N {
			//done
			r := make(board, N)
			copy(r, brd)
			bs = append(bs, r)
			//fmt.Println("brdf=", brd)
		} else {
			for i := 0; i < N; i++ {

				brd = append(brd, i)
				//				fmt.Println("brd1:", brd)
				if noColumns(brd) && noDiags(brd) {
					//ok to proceed to next row
					findBoards(row+1, brd)
				}
				brd = brd[:len(brd)-1] //remove last value
				//fmt.Println("brd2:", brd)
			}
		}
	}
	findBoards(0, b)
	return bs
}

func main() {
	q2 := compute(2)
	q3 := compute(3)
	q4 := compute(4)
	q8 := compute(8)
	fmt.Println(q2)
	fmt.Println(q3)
	fmt.Println(q4)
	fmt.Println(q8)
	fmt.Println(len(q8))

}
