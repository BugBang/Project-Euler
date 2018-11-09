package main

//https://projecteuler.net/problem=15
//Starting in the top left corner of a 2×2 grid, and only being able to
//move to the right and down, there are exactly 6 routes to the bottom right corner.

//How many such routes are there through a 20×20 grid?

func main() {
	grid := make([][]uint64, 41)
	for i := 0; i < 41; i++ {
		grid[i] = make([]uint64, 41)
	}
	grid[0][0] = 1
	for i := 1; i <= 40; i++ {
		grid[i][0] = 1
		for j := 1; j <= i; j++ {
			grid[i][j] = grid[i-1][j-1] + grid[i-1][j]
		}
	}
	println(grid[40][20])
}
