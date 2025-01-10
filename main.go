package main

import (
	"fmt"
	"math/rand"
	"time"
)

type World struct {
	height, width int
	grid          [][]bool
}

func NewWorld(width, height int) World {
	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
	}

	w := World{
		height,
		width,
		grid,
	}
	return w
}

func InitializeWorld(w World) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < w.height; i++ {
		for j := 0; j < w.width; j++ {
			w.grid[i][j] = rand.Intn(2) == 0
		}
	}
}

func CountAliveNeighbors(w World, x, y int) int {
	count := 0
	directions := []struct{ dx, dy int }{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
	for _, d := range directions {
		nx, ny := x+d.dx, y+d.dy
		if nx >= 0 && nx < w.width && ny >= 0 && ny < w.height && w.grid[ny][nx] == true {
			count++
		}
	}
	return count
}

func UpdateWorld(w World) World {
	newWorld := NewWorld(w.width, w.height)
	for y := 0; y < w.height; y++ {
		for x := 0; x < w.width; x++ {
			n := CountAliveNeighbors(w, x, y)
			if w.grid[y][x] {
				newWorld.grid[y][x] = n == 2 || n == 3
			} else {
				newWorld.grid[y][x] = n == 3
			}
		}
	}
	return newWorld
}

func displayWorld(w World) {
	for _, row := range w.grid {
		for _, cell := range row {
			if cell {
				fmt.Print("⬜")
			} else {
				fmt.Print("⬛")
			}
		}
		fmt.Println()
	}
}

func main() {
	maxSteps := 10
	height, width := 3, 3
	// create grid
	world := NewWorld(height, width)
	InitializeWorld(world)

	world = World{
		3, 3,
		[][]bool{
			[]bool{true, false, false},
			[]bool{false, false, false},
			[]bool{false, false, false}},
	}

	for i := 0; i < maxSteps; i++ {
		displayWorld(world)
		fmt.Println()
		world = UpdateWorld(world)
	}
}
