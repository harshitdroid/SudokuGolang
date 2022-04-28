package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
This function changes the
*/
func generateSudoku(base [][]int) {
	for i := 0; i < 100; i++ {
		random := randInt(0, 100)
		if random < 33 {
			swapLine(base)
		} else if random < 66 {
			swapCol(base)
		} else if random < 100 {
			swapNumber(base)
		}
	}
}

/*
makes the elements of the 9x9 matrix from the solved sudoku 0
random number 2 determines the chance of numbers in the grid being 0
*/
func generateUnsolvedSudoku(base [][]int) {
	randomNum1 := randInt(0, 100)
	randomNum2 := randInt(60, 90)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if randomNum1 < randomNum2 {
				base[i][j] = 0
			}
			randomNum1 = randInt(0, 100)
		}
	}
}

/*
Creates generates random integer in a given range
*/
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

/*
 helper function of generateSudoku
 swaps line of the matrix grid
*/
func swapLine(base [][]int) {
	random := [3]int{0, 3, 6}
	r := random[rand.Intn(len(random))]
	l1 := r + randInt(0, 3) // 0-2
	l2 := r + randInt(0, 3) // 0-2
	base[l1], base[l2] = base[l2], base[l1]
}

func swapCol(base [][]int) {
	randInts := [3]int{0, 3, 6}
	r := randInts[rand.Intn(len(randInts))]
	c1 := r + randInt(0, 3) // 0-2
	c2 := r + randInt(0, 3) // 0-2
	for line := 0; line < 9; line++ {
		base[line][c1], base[line][c2] = base[line][c2], base[line][c1]
	}
}

func swapNumber(base [][]int) {
	n1 := randInt(1, 10)
	n2 := randInt(1, 10)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if base[i][j] == n1 {
				base[i][j] = n2
			} else if base[i][j] == n2 {
				base[i][j] = n1
			}
		}
	}
}

func printBase(base [][]int) {
	fmt.Println("+-------+-------+-------+")
	for i := 0; i < 9; i++ {
		fmt.Print("| ")
		for j := 0; j < 9; j++ {
			if j == 3 || j == 6 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", base[i][j])
			if j == 8 {
				fmt.Print("|")
			}
		}
		if i == 2 || i == 5 || i == 8 {
			fmt.Println("\n+-------+-------+-------+")
		} else {
			fmt.Println()
		}

	}

}
func solve(base [][]int) bool {
	for rows := 0; rows < 9; rows++ {
		for cols := 0; cols < 9; cols++ {
			if base[rows][cols] == 0 {
				for testValue := 1; testValue <= 9; testValue++ {
					//if verify(base, rows, cols, testValue) {
					if verifyConcurrent(base, rows, cols, testValue) {
						base[rows][cols] = testValue
						if solve(base) {
							return true
						}
						base[rows][cols] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func verify(b [][]int, r int, c int, testValue int) bool {
	for n := 0; n < 9; n++ {
		if b[n][c] == testValue {
			return false
		}
		if b[r][n] == testValue {
			return false
		}
	}

	for row := (r / 3) * 3; row < (((r / 3) + 1) * 3); row++ {
		for col := (c / 3) * 3; col < (((c / 3) + 1) * 3); col++ {
			if b[row][col] == testValue {
				return false
			}
		}
	}
	return true
}

func verifyConcurrent(b [][]int, r int, c int, testValue int) bool {

	go part1(b, r, c, testValue)
	part2(b, r, c, testValue)
	// this will solve part 1 then part 2 and then goto return

	return true
}

func part1(b [][]int, r int, c int, testValue int) bool {
	for n := 0; n < 9; n++ {
		if b[n][c] == testValue {

			return false
		}
		if b[r][n] == testValue {
			return false
		}

	}
	return true
}

func part2(b [][]int, r int, c int, testValue int) bool {
	for row := (r / 3) * 3; row < ((r/3 + 1) * 3); row++ {
		for col := (c / 3) * 3; col < ((c/3 + 1) * 3); col++ {
			if b[row][col] == testValue {
				return false
			}
		}
	}
	return true
}

//func solve(base [][]int) {
//	start := time.Now()
//	var list []int
//	for i := 0; i < 9; i++ {
//		for j := 0; j < 9; j++ {
//			if base[j][i] == 0 {
//				list = append(list, (j*10)+i)
//			}
//		}
//	}
//
//	iterations := 0
//	for i := 0; i < len(list); i++ {
//		t := 0
//		l := list[i]
//		jc := l % 10
//		ic := (l - l%10) / 10
//
//		for {
//			t = base[ic][jc] + 1
//			base[ic][jc] = t
//			iterations++
//			//if verify(l, base) {
//			//	break
//			//}
//		}
//
//		if t > 9 {
//			if i == len(list) {
//				i = len(list) + 1
//			} else {
//				base[ic][jc] = 0
//				if i > 0 {
//					i = i - 2
//				} else {
//					i = -1
//				}
//			}
//		} else if l == 88 {
//			i = len(list) + 1
//		}
//	}
//	duration := time.Since(start)
//	printBase(base)
//	fmt.Printf("Time taken to solve %dns\n", duration.Nanoseconds())
//	fmt.Printf("Number of iterations: %d\n", iterations)
//}
//
//func verify(base [][]int, r int, c int, d int) bool {
//	go for row := 0; row < 9; row++ {
//		if base[row][c] == d {
//			return false
//		}
//	}
//	go for col := 0; col < 9; col++ {
//		if base[r][col] == d {
//			return false
//		}
//	}
//	for row := (r / 3) * 3; row < ((r/3 + 1) * 3); row++ {
//		for col := (c / 3) * 3; col < ((c/3 + 1) * 3); col++ {
//			if base[row][col] == d {
//				return false
//			}
//		}
//	}
//	return true
//}

//func verify(l int, b [][]int) bool {
//	j := l % 10
//	i := (l - l%10) / 10
//	var ic, jc int
//
//	boolVal := true
//	for i1 := 0; i1 < 9; i1++ {
//		if i1 != i && b[i1][j] == b[i][j] {
//			boolVal = false
//		}
//	}
//	for j1 := 0; j1 < 9; j1++ {
//		if j1 != j && b[i][j1] == b[i][j] {
//			boolVal = false
//		}
//	}
//
//	if i >= 0 && i < 3 {
//		ic = 1
//	} else if i > 2 && i < 6 {
//		ic = 2
//	} else {
//		ic = 3
//	}
//
//	if j >= 0 && j < 3 {
//		jc = 1
//	} else if j > 2 && j < 6 {
//		jc = 2
//	} else {
//		jc = 3
//	}
//
//	for i1 := ic*3 - 3; i1 < ic*3; i1++ {
//		for j1 := jc*3 - 3; j1 < jc*3; j1++ {
//			if (j1 != j || i1 != i) && b[i1][j1] == b[i][j] {
//				boolVal = false
//			}
//		}
//	}
//	return boolVal
//}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	base := make([][]int, 9)
	//base[0] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//base[1] = []int{4, 5, 6, 7, 8, 9, 1, 2, 3}
	//base[2] = []int{7, 8, 9, 1, 2, 3, 4, 5, 6}
	//base[3] = []int{2, 3, 4, 5, 6, 7, 8, 9, 1}
	//base[4] = []int{5, 6, 7, 8, 9, 1, 2, 3, 4}
	//base[5] = []int{8, 9, 1, 2, 3, 4, 5, 6, 7}
	//base[6] = []int{3, 4, 5, 6, 7, 8, 9, 1, 2}
	//base[7] = []int{6, 7, 8, 9, 1, 2, 3, 4, 5}
	//base[8] = []int{9, 1, 2, 3, 4, 5, 6, 7, 8}

	base[0] = []int{8, 2, 7, 1, 5, 4, 3, 9, 6}
	base[1] = []int{9, 6, 5, 3, 2, 7, 1, 4, 8}
	base[2] = []int{3, 4, 1, 6, 8, 9, 7, 5, 2}
	base[3] = []int{5, 9, 3, 4, 6, 8, 2, 7, 1}
	base[4] = []int{4, 7, 2, 5, 1, 3, 6, 8, 9}
	base[5] = []int{6, 1, 8, 9, 7, 2, 4, 3, 5}
	base[6] = []int{7, 8, 6, 2, 3, 5, 9, 1, 4}
	base[7] = []int{1, 5, 4, 7, 9, 6, 8, 2, 3}
	base[8] = []int{2, 3, 9, 8, 4, 1, 5, 6, 7}
	generateSudoku(base)
	//printBase(base)
	generateUnsolvedSudoku(base)
	printBase(base)
	solve(base)
	printBase(base)

}
